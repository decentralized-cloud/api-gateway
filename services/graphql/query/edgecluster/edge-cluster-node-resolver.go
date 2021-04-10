// Package edgecluster implements different edge cluster GraphQL query resovlers required by the GraphQL transport layer
package edgecluster

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/edgecluster"
	edgeclusterGrpcContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type edgeClusterNodeResolver struct {
	logger          *zap.Logger
	resolverCreator types.ResolverCreatorContract
	node            *edgeclusterGrpcContract.EdgeClusterNode
}

// NewEdgeClusterNodeResolver creates new instance of the edgeClusterNodeResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// logger: Mandatory. Reference to the logger service
// node: Mandatory. Contains information about the edge cluster node.
// Returns the new instance or error if something goes wrong
func NewEdgeClusterNodeResolver(
	ctx context.Context,
	logger *zap.Logger,
	resolverCreator types.ResolverCreatorContract,
	node *edgeclusterGrpcContract.EdgeClusterNode) (edgecluster.NodeResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if node == nil {
		return nil, commonErrors.NewArgumentNilError("node", "node is required")
	}

	return &edgeClusterNodeResolver{
		logger:          logger,
		resolverCreator: resolverCreator,
		node:            node,
	}, nil
}

// Metadata contains the node metadata
// ctx: Mandatory. Reference to the context
// Returns the node metadata resolver or error if something goes wrong.
func (r *edgeClusterNodeResolver) Metadata(ctx context.Context) (edgecluster.ObjectMetaResolverContract, error) {
	return r.resolverCreator.NewObjectMetaResolver(ctx, r.node.Metadata)
}

// Status contains the most recently observed status of the node
// ctx: Mandatory. Reference to the context
// Returns the most recently observed status of the node resolver or error if something goes wrong.
func (r *edgeClusterNodeResolver) Status(ctx context.Context) (edgecluster.NodeStatusResolverContract, error) {
	return r.resolverCreator.NewNodeStatusResolver(ctx, r.node.Status)
}
