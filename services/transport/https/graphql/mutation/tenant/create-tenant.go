// Package tenant implements tenant mutation required by the GraphQL transport layer
package tenant

import (
	"context"
	"errors"

	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types/tenant"
	tenantGrpcContract "github.com/decentralized-cloud/tenant/contract/grpc/go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type createTenant struct {
	logger              *zap.Logger
	resolverCreator     types.ResolverCreatorContract
	tenantServiceClient tenantGrpcContract.TenantServiceClient
}

type createTenantPayloadResolver struct {
	resolverCreator  types.ResolverCreatorContract
	clientMutationId *string
	tenantID         string
	tenant           *tenantGrpcContract.Tenant
}

// NewCreateTenant creates new instance of the createTenant, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can create new instances of resolvers
// logger: Mandatory. Reference to the logger service
// tenantServiceClient: Mandatory. Reference to the tenant service gRPC client that will be used to contact the tenant service
// Returns the new instance or error if something goes wrong
func NewCreateTenant(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	logger *zap.Logger,
	tenantServiceClient tenantGrpcContract.TenantServiceClient) (tenant.CreateTenantContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if tenantServiceClient == nil {
		return nil, commonErrors.NewArgumentNilError("tenantServiceClient", "tenantServiceClient is required")
	}

	return &createTenant{
		logger:              logger,
		resolverCreator:     resolverCreator,
		tenantServiceClient: tenantServiceClient,
	}, nil
}

// NewCreateTenantPayloadResolver updates new instance of the createTenantPayloadResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can update new instances of resolvers
// clientMutationId: Optional. Reference to the client mutation ID
// tenantID: Mandatory. The tenant unique identifier
// tenant: Optional. The tenant details
// Returns the new instance or error if something goes wrong
func NewCreateTenantPayloadResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	clientMutationId *string,
	tenantID string,
	tenant *tenantGrpcContract.Tenant) (tenant.CreateTenantPayloadResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	return &createTenantPayloadResolver{
		resolverCreator:  resolverCreator,
		clientMutationId: clientMutationId,
		tenantID:         tenantID,
		tenant:           tenant,
	}, nil
}

// MutateAndGetPayload creates a new tenant and returns the payload contains the result of creating a new tenant
// ctx: Mandatory. Reference to the context
// args: Mandatory. Reference to the input argument contains tenant information to create
// Returns the new tenant payload or error if something goes wrong
func (m *createTenant) MutateAndGetPayload(
	ctx context.Context,
	args tenant.CreateTenantInputArgument) (tenant.CreateTenantPayloadResolverContract, error) {
	response, err := m.tenantServiceClient.CreateTenant(
		ctx,
		&tenantGrpcContract.CreateTenantRequest{
			Tenant: &tenantGrpcContract.Tenant{
				Name: args.Input.Name,
			}})
	if err != nil {
		return nil, err
	}

	if response.Error != tenantGrpcContract.Error_NO_ERROR {
		return nil, errors.New(response.ErrorMessage)
	}

	return m.resolverCreator.NewCreateTenantPayloadResolver(
		ctx,
		args.Input.ClientMutationId,
		response.TenantID,
		response.Tenant)
}

// Tenant returns the new tenant inforamtion
// ctx: Mandatory. Reference to the context
// Returns the new tenant inforamtion
func (r *createTenantPayloadResolver) Tenant(ctx context.Context) (tenant.TenantTypeEdgeResolverContract, error) {
	resolver, err := r.resolverCreator.NewTenantTypeEdgeResolver(ctx, r.tenantID, "Not implemented", r.tenant)

	return resolver, err
}

// ClientMutationId returns the client mutation ID that was provided as part of the mutation request
// ctx: Mandatory. Reference to the context
// Returns the provided clientMutationId as part of mutation request
func (r *createTenantPayloadResolver) ClientMutationId(ctx context.Context) *string {
	return r.clientMutationId
}
