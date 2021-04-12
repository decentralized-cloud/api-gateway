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

type nodeStatusResolver struct {
	logger          *zap.Logger
	resolverCreator types.ResolverCreatorContract
	status          *edgeclusterGrpcContract.NodeStatus
}

// NewNodeStatusResolver creates new instance of the nodeStatusResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// logger: Mandatory. Reference to the logger service
// node: Mandatory. Contains information about the edge cluster node status.
// Returns the new instance or error if something goes wrong
func NewNodeStatusResolver(
	ctx context.Context,
	logger *zap.Logger,
	resolverCreator types.ResolverCreatorContract,
	status *edgeclusterGrpcContract.NodeStatus) (edgecluster.NodeStatusResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if status == nil {
		return nil, commonErrors.NewArgumentNilError("status", "status is required")
	}

	return &nodeStatusResolver{
		logger:          logger,
		resolverCreator: resolverCreator,
		status:          status,
	}, nil
}

// Conditions is an array of current observed node conditions.
// ctx: Mandatory. Reference to the context
// Returns an array of current observed node conditions resolver or error if something goes wrong.
func (r *nodeStatusResolver) Conditions(ctx context.Context) ([]edgecluster.NodeConditionResolverContract, error) {
	response := []edgecluster.NodeConditionResolverContract{}
	for _, condition := range r.status.Conditions {
		if resolver, err := r.resolverCreator.NewNodeConditionResolver(ctx, condition); err != nil {
			return nil, err
		} else {
			response = append(response, resolver)
		}
	}

	return response, nil
}

// Addresses is the list of addresses reachable to the node.
// ctx: Mandatory. Reference to the context
// Returns the list of addresses reachable to the node resolver or error if something goes wrong.
func (r *nodeStatusResolver) Addresses(ctx context.Context) ([]edgecluster.NodeAddressResolverContract, error) {
	response := []edgecluster.NodeAddressResolverContract{}
	for _, address := range r.status.Addresses {
		if resolver, err := r.resolverCreator.NewNodeAddressResolver(ctx, address); err != nil {
			return nil, err
		} else {
			response = append(response, resolver)
		}
	}

	return response, nil
}

// NodeInfo is the set of ids/uuids to uniquely identify the node.
// ctx: Mandatory. Reference to the context
// Returns the set of ids/uuids to uniquely identify the node resolver or error if something goes wrong.
func (r *nodeStatusResolver) NodeInfo(ctx context.Context) (edgecluster.NodeSystemInfoResolverContract, error) {
	return r.resolverCreator.NewNodeSystemInfoResolver(ctx, r.status.NodeInfo)
}
