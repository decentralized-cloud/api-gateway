// Package graphql implements functions to expose api-gateway service endpoint using GraphQL protocol.
package graphql

import (
	"context"

	queryedgecluster "github.com/decentralized-cloud/api-gateway/services/graphql/query/edgecluster"
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/edgecluster"
	edgeclusterGrpcContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
)

// NewEdgeClusterNodeResolver creates new instance of the NodeResolverContract, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// node: Mandatory. Contains information about the edge cluster node.
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewEdgeClusterNodeResolver(
	ctx context.Context,
	node *edgeclusterGrpcContract.EdgeClusterNode) (edgecluster.NodeResolverContract, error) {
	return queryedgecluster.NewEdgeClusterNodeResolver(
		ctx,
		creator.logger,
		creator,
		node)
}

// NewNodeStatusResolver creates new instance of the NodeStatusResolverContract, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// status: Mandatory. Contains information about the edge cluster node status.
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewNodeStatusResolver(
	ctx context.Context,
	status *edgeclusterGrpcContract.NodeStatus) (edgecluster.NodeStatusResolverContract, error) {
	return queryedgecluster.NewNodeStatusResolver(
		ctx,
		creator.logger,
		creator,
		status)
}

// NewNodeConditionResolver creates new instance of the NodeConditionResolverContract, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// condition: Mandatory. Contains condition information for a node.
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewNodeConditionResolver(
	ctx context.Context,
	condition *edgeclusterGrpcContract.NodeCondition) (edgecluster.NodeConditionResolverContract, error) {
	return queryedgecluster.NewNodeConditionResolver(
		ctx,
		creator.logger,
		condition)
}

// NewNodeAddressResolver creates new instance of the NodeAddressResolverContract, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// nodeAddresss: Mandatory. Contains information for the edge cluster node's address.
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewNodeAddressResolver(
	ctx context.Context,
	nodeAddresss *edgeclusterGrpcContract.EdgeClusterNodeAddress) (edgecluster.NodeAddressResolverContract, error) {
	return queryedgecluster.NewNodeAddressResolver(
		ctx,
		creator.logger,
		nodeAddresss)
}

// NewNodeSystemInfoResolver creates new instance of the NodeSystemInfoResolverContract, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// nodeInfo: Mandatory. Contains a set of ids/uuids to uniquely identify the node.
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewNodeSystemInfoResolver(
	ctx context.Context,
	nodeInfo *edgeclusterGrpcContract.NodeSystemInfo) (edgecluster.NodeSystemInfoResolverContract, error) {
	return queryedgecluster.NewNodeSystemInfoResolver(
		ctx,
		creator.logger,
		nodeInfo)
}
