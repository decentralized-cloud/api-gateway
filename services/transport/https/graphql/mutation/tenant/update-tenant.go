// package tenant implements tenant mutation required by the GraphQL transport layer
package tenant

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types/tenant"
	"github.com/graph-gophers/graphql-go"
	"github.com/lucsky/cuid"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type updateTenant struct {
	logger          *zap.Logger
	resolverCreator types.ResolverCreatorContract
}

type updateTenantPayloadResolver struct {
	resolverCreator  types.ResolverCreatorContract
	clientMutationId *string
}

// NewUpdateTenant updates new instance of the updateTenant, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can update new instances of resolvers
// logger: Mandatory. Reference to the logger service
// Returns the new instance or error if something goes wrong
func NewUpdateTenant(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	logger *zap.Logger) (tenant.UpdateTenantContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	return &updateTenant{
		logger:          logger,
		resolverCreator: resolverCreator,
	}, nil
}

// Tenant returns the updated tenant inforamtion
// ctx: Mandatory. Reference to the context
// Returns the updated tenant inforamtion
func NewUpdateTenantPayloadResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	clientMutationId *string) (tenant.UpdateTenantPayloadResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	return &updateTenantPayloadResolver{
		resolverCreator:  resolverCreator,
		clientMutationId: clientMutationId,
	}, nil
}

// MutateAndGetPayload update an existing tenant and returns the payload contains the result of updating an existing tenant
// ctx: Mandatory. Reference to the context
// args: Mandatory. Reference to the input argument contains tenant information to update
// Returns the updated tenant payload or error if something goes wrong
func (m *updateTenant) MutateAndGetPayload(
	ctx context.Context,
	args tenant.UpdateTenantInputArgument) (tenant.UpdateTenantPayloadResolverContract, error) {
	return m.resolverCreator.NewUpdateTenantPayloadResolver(ctx, args.Input.ClientMutationId)
}

// Tenant returns the updated tenant inforamtion
// ctx: Mandatory. Reference to the context
// Returns the updated tenant inforamtion
func (r *updateTenantPayloadResolver) Tenant(ctx context.Context) (tenant.TenantTypeEdgeResolverContract, error) {
	resolver, err := r.resolverCreator.NewTenantTypeEdgeResolver(ctx, graphql.ID(cuid.New()), "Not implemented")

	return resolver, err
}

// ClientMutationId returns the client mutation ID that was provided as part of the mutation request
// ctx: Mandatory. Reference to the context
// Returns the provided clientMutationId as part of mutation request
func (r *updateTenantPayloadResolver) ClientMutationId(ctx context.Context) *string {
	return r.clientMutationId
}
