// package resolver implements different GraphQL resolvers required by the GraphQL transport layer
package mutation

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/transport/graphql/types"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type deleteEdgeCluster struct {
	logger          *zap.Logger
	resolverCreator types.ResolverCreatorContract
}

type deleteEdgeClusterPayloadResolver struct {
	resolverCreator  types.ResolverCreatorContract
	clientMutationId *string
}

// NewDeleteEdgeCluster deletes new instance of the deleteEdgeCluster, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can delete new instances of resolvers
// logger: Mandatory. Reference to the logger service
// Returns the new instance or error if something goes wrong
func NewDeleteEdgeCluster(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	logger *zap.Logger) (types.DeleteEdgeClusterContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	return &deleteEdgeCluster{
		logger:          logger,
		resolverCreator: resolverCreator,
	}, nil
}

// EdgeCluster returns the deleted edge cluster inforamtion
// ctx: Mandatory. Reference to the context
// Returns the deleted edge cluster inforamtion
func NewDeleteEdgeClusterPayloadResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	clientMutationId *string) (types.DeleteEdgeClusterPayloadResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	return &deleteEdgeClusterPayloadResolver{
		resolverCreator:  resolverCreator,
		clientMutationId: clientMutationId,
	}, nil
}

// MutateAndGetPayload delete an existing edge cluster and returns the payload contains the result of deleting an existing edge cluster
// ctx: Mandatory. Reference to the context
// args: Mandatory. Reference to the input argument contains edge cluster information to delete
// Returns the deleted edge cluster payload or error if something goes wrong
func (m *deleteEdgeCluster) MutateAndGetPayload(
	ctx context.Context,
	args types.DeleteEdgeClusterInputArgument) (types.DeleteEdgeClusterPayloadResolverContract, error) {
	return m.resolverCreator.NewDeleteEdgeClusterPayloadResolver(ctx, args.Input.ClientMutationId)
}

// ClientMutationId returns the client mutation ID that was provided as part of the mutation request
// ctx: Mandatory. Reference to the context
// Returns the provided clientMutationId as part of mutation request
func (r *deleteEdgeClusterPayloadResolver) ClientMutationId(ctx context.Context) *string {
	return r.clientMutationId
}
