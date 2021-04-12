// Package graphql implements functions to expose api-gateway service endpoint using GraphQL protocol.
package graphql

import (
	"context"

	queryedgecluster "github.com/decentralized-cloud/api-gateway/services/graphql/query/edgecluster"
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/edgecluster"
	edgeclusterGrpcContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
)

// NewEdgeClusterPodResolver creates new instance of the PodResolverContract, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// pod: Mandatory. Contains information about the edge cluster pod
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewEdgeClusterPodResolver(
	ctx context.Context,
	pod *edgeclusterGrpcContract.EdgeClusterPod) (edgecluster.PodResolverContract, error) {
	return queryedgecluster.NewEdgeClusterPodResolver(
		ctx,
		creator.logger,
		creator,
		pod)
}

// NewPodStatusResolver creates new instance of the PodStatusResolverContract, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// status: Mandatory. Contains information about the edge cluster pod status
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewPodStatusResolver(
	ctx context.Context,
	status *edgeclusterGrpcContract.PodStatus) (edgecluster.PodStatusResolverContract, error) {
	return queryedgecluster.NewPodStatusResolver(
		ctx,
		creator.logger,
		creator,
		status)
}

// NewPodSpecResolver creates new instance of the PodSpecResolverContract, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// spec: Mandatory. Contains information about the edge cluster pod specification
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewPodSpecResolver(
	ctx context.Context,
	spec *edgeclusterGrpcContract.PodSpec) (edgecluster.PodSpecResolverContract, error) {
	return queryedgecluster.NewPodSpecResolver(
		ctx,
		creator.logger,
		spec)
}

// NewPodConditionResolver creates new instance of the PodConditionResolverContract, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// podCondition: Mandatory. Contains condition information for a pod.
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewPodConditionResolver(
	ctx context.Context,
	podCondition *edgeclusterGrpcContract.PodCondition) (edgecluster.PodConditionResolverContract, error) {
	return queryedgecluster.NewPodConditionResolver(
		ctx,
		creator.logger,
		podCondition)
}
