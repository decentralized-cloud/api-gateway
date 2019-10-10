// package root implements GraphQL root resolvers required by the GraphQL transport layer
package root

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/transport/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/transport/graphql/types/edgecluster"
	"github.com/decentralized-cloud/api-gateway/services/transport/graphql/types/tenant"
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
	args tenant.CreateTenantInputArgument) (tenant.CreateTenantPayloadResolverContract, error) {
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
	args tenant.UpdateTenantInputArgument) (tenant.UpdateTenantPayloadResolverContract, error) {
	mutation, err := r.resolverCreator.NewUpdateTenant(ctx)
	if err != nil {
		return nil, err
	}

	return mutation.MutateAndGetPayload(ctx, args)
}

// DeleteTenant returns delete tenant mutator
// ctx: Mandatory. Reference to the context
// Returns the delete tenant mutator or error if something goes wrong
func (r *rootResolver) DeleteTenant(
	ctx context.Context,
	args tenant.DeleteTenantInputArgument) (tenant.DeleteTenantPayloadResolverContract, error) {
	mutation, err := r.resolverCreator.NewDeleteTenant(ctx)
	if err != nil {
		return nil, err
	}

	return mutation.MutateAndGetPayload(ctx, args)
}

// CreateEdgeCluster returns create edge cluster mutator
// ctx: Mandatory. Reference to the context
// Returns the create edge cluster mutator or error if something goes wrong
func (r *rootResolver) CreateEdgeCluster(
	ctx context.Context,
	args edgecluster.CreateEdgeClusterInputArgument) (edgecluster.CreateEdgeClusterPayloadResolverContract, error) {
	mutation, err := r.resolverCreator.NewCreateEdgeCluster(ctx)
	if err != nil {
		return nil, err
	}

	return mutation.MutateAndGetPayload(ctx, args)
}

// UpdateEdgeCluster returns update edge cluster mutator
// ctx: Mandatory. Reference to the context
// Returns the update edge cluster mutator or error if something goes wrong
func (r *rootResolver) UpdateEdgeCluster(
	ctx context.Context,
	args edgecluster.UpdateEdgeClusterInputArgument) (edgecluster.UpdateEdgeClusterPayloadResolverContract, error) {
	mutation, err := r.resolverCreator.NewUpdateEdgeCluster(ctx)
	if err != nil {
		return nil, err
	}

	return mutation.MutateAndGetPayload(ctx, args)
}

// DeleteEdgeCluster returns delete edge cluster mutator
// ctx: Mandatory. Reference to the context
// Returns the delete edge cluster mutator or error if something goes wrong
func (r *rootResolver) DeleteEdgeCluster(
	ctx context.Context,
	args edgecluster.DeleteEdgeClusterInputArgument) (edgecluster.DeleteEdgeClusterPayloadResolverContract, error) {
	mutation, err := r.resolverCreator.NewDeleteEdgeCluster(ctx)
	if err != nil {
		return nil, err
	}

	return mutation.MutateAndGetPayload(ctx, args)
}
