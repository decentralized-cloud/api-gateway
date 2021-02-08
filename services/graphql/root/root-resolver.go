// Package root implements GraphQL root resolvers required by the GraphQL transport layer
package root

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/edgecluster"
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/project"
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
	// TODO: mortezaalizadeh, 16/11/2019 user ID hard coded until we have authentication flow implemented
	return r.resolverCreator.NewUserResolver(ctx, "ck30lptw000060133f1rg1uf9")
}

// CreateProject returns create project mutator
// ctx: Mandatory. Reference to the context
// Returns the create project mutator or error if something goes wrong
func (r *rootResolver) CreateProject(
	ctx context.Context,
	args project.CreateProjectInputArgument) (project.CreateProjectPayloadResolverContract, error) {
	mutation, err := r.resolverCreator.NewCreateProject(ctx)
	if err != nil {
		return nil, err
	}

	return mutation.MutateAndGetPayload(ctx, args)
}

// UpdateProject returns update project mutator
// ctx: Mandatory. Reference to the context
// Returns the update project mutator or error if something goes wrong
func (r *rootResolver) UpdateProject(
	ctx context.Context,
	args project.UpdateProjectInputArgument) (project.UpdateProjectPayloadResolverContract, error) {
	mutation, err := r.resolverCreator.NewUpdateProject(ctx)
	if err != nil {
		return nil, err
	}

	return mutation.MutateAndGetPayload(ctx, args)
}

// DeleteProject returns delete project mutator
// ctx: Mandatory. Reference to the context
// Returns the delete project mutator or error if something goes wrong
func (r *rootResolver) DeleteProject(
	ctx context.Context,
	args project.DeleteProjectInputArgument) (project.DeleteProjectPayloadResolverContract, error) {
	mutation, err := r.resolverCreator.NewDeleteProject(ctx)
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
