// package edgecluster implements edge cluster mutation required by the GraphQL transport layer
package edgeclster

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/transport/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/transport/graphql/types/edgecluster"
	"github.com/graph-gophers/graphql-go"
	"github.com/lucsky/cuid"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type updateEdgeCluster struct {
	logger          *zap.Logger
	resolverCreator types.ResolverCreatorContract
}

type updateEdgeClusterPayloadResolver struct {
	resolverCreator  types.ResolverCreatorContract
	clientMutationId *string
}

// NewUpdateEdgeCluster updates new instance of the updateEdgeCluster, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can update new instances of resolvers
// logger: Mandatory. Reference to the logger service
// Returns the new instance or error if something goes wrong
func NewUpdateEdgeCluster(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	logger *zap.Logger) (edgecluster.UpdateEdgeClusterContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	return &updateEdgeCluster{
		logger:          logger,
		resolverCreator: resolverCreator,
	}, nil
}

// EdgeCluster returns the updated edge cluster inforamtion
// ctx: Mandatory. Reference to the context
// Returns the updated edge cluster inforamtion
func NewUpdateEdgeClusterPayloadResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	clientMutationId *string) (edgecluster.UpdateEdgeClusterPayloadResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	return &updateEdgeClusterPayloadResolver{
		resolverCreator:  resolverCreator,
		clientMutationId: clientMutationId,
	}, nil
}

// MutateAndGetPayload update an existing edge cluster and returns the payload contains the result of updating an existing edge cluster
// ctx: Mandatory. Reference to the context
// args: Mandatory. Reference to the input argument contains edge cluster information to update
// Returns the updated edge cluster payload or error if something goes wrong
func (m *updateEdgeCluster) MutateAndGetPayload(
	ctx context.Context,
	args edgecluster.UpdateEdgeClusterInputArgument) (edgecluster.UpdateEdgeClusterPayloadResolverContract, error) {
	return m.resolverCreator.NewUpdateEdgeClusterPayloadResolver(ctx, args.Input.ClientMutationId)
}

// EdgeCluster returns the updated edge cluster inforamtion
// ctx: Mandatory. Reference to the context
// Returns the updated edge cluster inforamtion
func (r *updateEdgeClusterPayloadResolver) EdgeCluster(ctx context.Context) (edgecluster.EdgeClusterTypeEdgeResolverContract, error) {
	resolver, err := r.resolverCreator.NewEdgeClusterTypeEdgeResolver(ctx, graphql.ID(cuid.New()), "New edge cluster cursor")

	return resolver, err
}

// ClientMutationId returns the client mutation ID that was provided as part of the mutation request
// ctx: Mandatory. Reference to the context
// Returns the provided clientMutationId as part of mutation request
func (r *updateEdgeClusterPayloadResolver) ClientMutationId(ctx context.Context) *string {
	return r.clientMutationId
}
