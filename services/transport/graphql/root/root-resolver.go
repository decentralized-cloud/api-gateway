// package root implements GraphQL root resolvers required by the GraphQL transport layer
package root

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/transport/graphql/types"
	"github.com/graph-gophers/graphql-go"
	"github.com/lucsky/cuid"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type rootResolver struct {
	logger          *zap.Logger
	resolverCreator types.ResolverCreatorContract
}

// NewRootResolver creates new instance of the rootResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can create new instances of resolvers
// logger: Mandatory. Reference to the logger service
// Returns the new instance or error if something goes wrong
func NewRootResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	logger *zap.Logger) (types.RootResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	return &rootResolver{
		logger:          logger,
		resolverCreator: resolverCreator,
	}, nil
}

// User returns user resolver
// ctx: Mandatory. Reference to the context
// Returns the user resolver or error if something goes wrong
func (r *rootResolver) User(ctx context.Context) (types.UserResolverContract, error) {
	return r.resolverCreator.NewUserResolver(ctx, graphql.ID(cuid.New()))
}

// CreateTenant returns create tenant mutator
// ctx: Mandatory. Reference to the context
// Returns the create tenant mutator or error if something goes wrong
func (r *rootResolver) CreateTenant(
	ctx context.Context,
	args types.CreateTenantInputArgument) (types.CreateTenantPayloadResolverContract, error) {
	mutation, err := r.resolverCreator.NewCreateTenant(ctx)
	if err != nil {
		return nil, err
	}

	return mutation.MutateAndGetPayload(ctx, args)
}

// UpdateTenant returns update tenant mutator
// ctx: Mandatory. Reference to the context
// Returns the update tenant mutator or error if something goes wrong
func (r *rootResolver) UpdateTenant(
	ctx context.Context,
	args types.UpdateTenantInputArgument) (types.UpdateTenantPayloadResolverContract, error) {
	mutation, err := r.resolverCreator.NewUpdateTenant(ctx)
	if err != nil {
		return nil, err
	}

	return mutation.MutateAndGetPayload(ctx, args)
}
