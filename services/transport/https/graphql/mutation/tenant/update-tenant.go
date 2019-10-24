// Package tenant implements tenant mutation required by the GraphQL transport layer
package tenant

import (
	"context"
	"errors"

	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types/tenant"
	tenantGrpcContract "github.com/decentralized-cloud/tenant/contract/grpc/go"
	"github.com/graph-gophers/graphql-go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type updateTenant struct {
	logger              *zap.Logger
	resolverCreator     types.ResolverCreatorContract
	tenantServiceClient tenantGrpcContract.TenantServiceClient
}

type updateTenantPayloadResolver struct {
	resolverCreator  types.ResolverCreatorContract
	clientMutationId *string
	tenantID         string
}

// NewUpdateTenant updates new instance of the updateTenant, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can update new instances of resolvers
// logger: Mandatory. Reference to the logger service
// Returns the new instance or error if something goes wrong
func NewUpdateTenant(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	logger *zap.Logger,
	tenantServiceClient tenantGrpcContract.TenantServiceClient) (tenant.UpdateTenantContract, error) {
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

	return &updateTenant{
		logger:              logger,
		resolverCreator:     resolverCreator,
		tenantServiceClient: tenantServiceClient,
	}, nil
}

// NewUpdateTenantPayloadResolver updates new instance of the updateTenantPayloadResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can update new instances of resolvers
// clientMutationId: Optional. Reference to the client mutation ID
// Returns the new instance or error if something goes wrong
func NewUpdateTenantPayloadResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	clientMutationId *string,
	tenantID string) (tenant.UpdateTenantPayloadResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	return &updateTenantPayloadResolver{
		resolverCreator:  resolverCreator,
		clientMutationId: clientMutationId,
		tenantID:         tenantID,
	}, nil
}

// MutateAndGetPayload update an existing tenant and returns the payload contains the result of updating an existing tenant
// ctx: Mandatory. Reference to the context
// args: Mandatory. Reference to the input argument contains tenant information to update
// Returns the updated tenant payload or error if something goes wrong
func (m *updateTenant) MutateAndGetPayload(
	ctx context.Context,
	args tenant.UpdateTenantInputArgument) (tenant.UpdateTenantPayloadResolverContract, error) {
	tenantID := string(args.Input.TenantID)

	response, err := m.tenantServiceClient.UpdateTenant(
		ctx,
		&tenantGrpcContract.UpdateTenantRequest{
			TenantID: tenantID,
			Tenant: &tenantGrpcContract.Tenant{
				Name: args.Input.Name,
			}})

	if err != nil {
		return nil, err
	}

	if response.Error != tenantGrpcContract.Error_NO_ERROR {
		return nil, errors.New(response.ErrorMessage)
	}

	return m.resolverCreator.NewUpdateTenantPayloadResolver(
		ctx,
		args.Input.ClientMutationId,
		tenantID,
	)
}

// Tenant returns the updated tenant inforamtion
// ctx: Mandatory. Reference to the context
// Returns the updated tenant inforamtion
func (r *updateTenantPayloadResolver) Tenant(ctx context.Context) (tenant.TenantTypeEdgeResolverContract, error) {
	resolver, err := r.resolverCreator.NewTenantTypeEdgeResolver(ctx, graphql.ID(r.tenantID), "Not implemented")

	return resolver, err
}

// ClientMutationId returns the client mutation ID that was provided as part of the mutation request
// ctx: Mandatory. Reference to the context
// Returns the provided clientMutationId as part of mutation request
func (r *updateTenantPayloadResolver) ClientMutationId(ctx context.Context) *string {
	return r.clientMutationId
}
