// Package tenant implements tenant mutation required by the GraphQL transport layer
package tenant

import (
	"context"
	"errors"

	"github.com/decentralized-cloud/api-gateway/services/configuration"
	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types/tenant"
	tenantGrpcContract "github.com/decentralized-cloud/tenant/contract/grpc/go"
	"github.com/graph-gophers/graphql-go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type createTenant struct {
	logger               *zap.Logger
	resolverCreator      types.ResolverCreatorContract
	tenantServiceAddress string
}

type createTenantPayloadResolver struct {
	resolverCreator  types.ResolverCreatorContract
	clientMutationId *string
	tenantID         string
}

// NewCreateTenant creates new instance of the createTenant, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can create new instances of resolvers
// logger: Mandatory. Reference to the logger service
// Returns the new instance or error if something goes wrong
func NewCreateTenant(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	logger *zap.Logger,
	configurationService configuration.ConfigurationContract) (tenant.CreateTenantContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if configurationService == nil {
		return nil, commonErrors.NewArgumentNilError("configurationService", "configurationService is required")
	}

	tenantServiceAddress, err := configurationService.GetTenantServiceAddress()
	if err != nil {
		return nil, err
	}

	return &createTenant{
		logger:               logger,
		resolverCreator:      resolverCreator,
		tenantServiceAddress: tenantServiceAddress,
	}, nil
}

// NewCreateTenantPayloadResolver updates new instance of the createTenantPayloadResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can update new instances of resolvers
// clientMutationId: Optional. Reference to the client mutation ID
// tenantID: Mandatory. The tenant unique identifier
// Returns the new instance or error if something goes wrong
func NewCreateTenantPayloadResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	clientMutationId *string,
	tenantID string) (tenant.CreateTenantPayloadResolverContract, error) {
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
	}, nil
}

// MutateAndGetPayload creates a new tenant and returns the payload contains the result of creating a new tenant
// ctx: Mandatory. Reference to the context
// args: Mandatory. Reference to the input argument contains tenant information to create
// Returns the new tenant payload or error if something goes wrong
func (m *createTenant) MutateAndGetPayload(
	ctx context.Context,
	args tenant.CreateTenantInputArgument) (tenant.CreateTenantPayloadResolverContract, error) {
	connection, err := grpc.Dial(m.tenantServiceAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer connection.Close()

	client := tenantGrpcContract.NewTenantServiceClient(connection)
	response, err := client.CreateTenant(
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
		response.TenantID)
}

// Tenant returns the new tenant inforamtion
// ctx: Mandatory. Reference to the context
// Returns the new tenant inforamtion
func (r *createTenantPayloadResolver) Tenant(ctx context.Context) (tenant.TenantTypeEdgeResolverContract, error) {
	resolver, err := r.resolverCreator.NewTenantTypeEdgeResolver(ctx, graphql.ID(r.tenantID), "Not implemented")

	return resolver, err
}

// ClientMutationId returns the client mutation ID that was provided as part of the mutation request
// ctx: Mandatory. Reference to the context
// Returns the provided clientMutationId as part of mutation request
func (r *createTenantPayloadResolver) ClientMutationId(ctx context.Context) *string {
	return r.clientMutationId
}
