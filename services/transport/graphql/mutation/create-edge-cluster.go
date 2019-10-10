// package resolver implements different GraphQL resolvers required by the GraphQL transport layer
package mutation

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/transport/graphql/types"
	"github.com/graph-gophers/graphql-go"
	"github.com/lucsky/cuid"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type createEdgeCluster struct {
	logger          *zap.Logger
	resolverCreator types.ResolverCreatorContract
}

type createEdgeClusterPayloadResolver struct {
	resolverCreator  types.ResolverCreatorContract
	clientMutationId *string
}

// NewCreateEdgeCluster creates new instance of the createEdgeCluster, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can create new instances of resolvers
// logger: Mandatory. Reference to the logger service
// Returns the new instance or error if something goes wrong
func NewCreateEdgeCluster(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	logger *zap.Logger) (types.CreateEdgeClusterContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	return &createEdgeCluster{
		logger:          logger,
		resolverCreator: resolverCreator,
	}, nil
}

// NewCreateEdgeClusterPayloadResolver creates new instance of the createEdgeClusterPayloadResolvere, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can create new instances of resolvers
// Returns the new instance or error if something goes wrong
func NewCreateEdgeClusterPayloadResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	clientMutationId *string) (types.CreateEdgeClusterPayloadResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	return &createEdgeClusterPayloadResolver{
		resolverCreator:  resolverCreator,
		clientMutationId: clientMutationId,
	}, nil
}

// MutateAndGetPayload creates a new edge cluster and returns the payload contains the result of creating a new edge cluster
// ctx: Mandatory. Reference to the context
// args: Mandatory. Reference to the input argument contains edge cluster information to create
// Returns the new edge cluster payload or error if something goes wrong
func (m *createEdgeCluster) MutateAndGetPayload(
	ctx context.Context,
	args types.CreateEdgeClusterInputArgument) (types.CreateEdgeClusterPayloadResolverContract, error) {
	return m.resolverCreator.NewCreateEdgeClusterPayloadResolver(ctx, args.Input.ClientMutationId)
}

// EdgeCluster returns the new edge cluster inforamtion
// ctx: Mandatory. Reference to the context
// Returns the new edge cluster inforamtion
func (r *createEdgeClusterPayloadResolver) EdgeCluster(ctx context.Context) (types.EdgeClusterTypeEdgeResolverContract, error) {
	resolver, err := r.resolverCreator.NewEdgeClusterTypeEdgeResolver(ctx, graphql.ID(cuid.New()), "New edge cluster cursor")

	return resolver, err
}

// ClientMutationId returns the client mutation ID that was provided as part of the mutation request
// ctx: Mandatory. Reference to the context
// Returns the provided clientMutationId as part of mutation request
func (r *createEdgeClusterPayloadResolver) ClientMutationId(ctx context.Context) *string {
	return r.clientMutationId
}
