// package tenant implements tenant mutation required by the GraphQL transport layer
package tenant

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/transport/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/transport/graphql/types/tenant"
	"github.com/graph-gophers/graphql-go"
	"github.com/lucsky/cuid"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type createTenant struct {
	logger          *zap.Logger
	resolverCreator types.ResolverCreatorContract
}

type createTenantPayloadResolver struct {
	resolverCreator  types.ResolverCreatorContract
	clientMutationId *string
}

// NewCreateTenant creates new instance of the createTenant, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can create new instances of resolvers
// logger: Mandatory. Reference to the logger service
// Returns the new instance or error if something goes wrong
func NewCreateTenant(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	logger *zap.Logger) (tenant.CreateTenantContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	return &createTenant{
		logger:          logger,
		resolverCreator: resolverCreator,
	}, nil
}

// NewCreateTenantPayloadResolver creates new instance of the createTenantPayloadResolvere, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can create new instances of resolvers
// Returns the new instance or error if something goes wrong
func NewCreateTenantPayloadResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	clientMutationId *string) (tenant.CreateTenantPayloadResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	return &createTenantPayloadResolver{
		resolverCreator:  resolverCreator,
		clientMutationId: clientMutationId,
	}, nil
}

// MutateAndGetPayload creates a new tenant and returns the payload contains the result of creating a new tenant
// ctx: Mandatory. Reference to the context
// args: Mandatory. Reference to the input argument contains tenant information to create
// Returns the new tenant payload or error if something goes wrong
func (m *createTenant) MutateAndGetPayload(
	ctx context.Context,
	args tenant.CreateTenantInputArgument) (tenant.CreateTenantPayloadResolverContract, error) {
	return m.resolverCreator.NewCreateTenantPayloadResolver(ctx, args.Input.ClientMutationId)
}

// Tenant returns the new tenant inforamtion
// ctx: Mandatory. Reference to the context
// Returns the new tenant inforamtion
func (r *createTenantPayloadResolver) Tenant(ctx context.Context) (tenant.TenantTypeEdgeResolverContract, error) {
	resolver, err := r.resolverCreator.NewTenantTypeEdgeResolver(ctx, graphql.ID(cuid.New()), "New tenant cursor")

	return resolver, err
}

// ClientMutationId returns the client mutation ID that was provided as part of the mutation request
// ctx: Mandatory. Reference to the context
// Returns the provided clientMutationId as part of mutation request
func (r *createTenantPayloadResolver) ClientMutationId(ctx context.Context) *string {
	return r.clientMutationId
}
