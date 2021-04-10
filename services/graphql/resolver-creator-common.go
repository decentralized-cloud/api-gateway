// Package graphql implements functions to expose api-gateway service endpoint using GraphQL protocol.
package graphql

import (
	"context"

	queryedgecluster "github.com/decentralized-cloud/api-gateway/services/graphql/query/edgecluster"
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/edgecluster"
	edgeclusterGrpcContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
)

// NewObjectMetaResolver creates new instance of the ObjectMetaResolverContract, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// metadata: Mandatory. Contains the object metadata.
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewObjectMetaResolver(
	ctx context.Context,
	metadata *edgeclusterGrpcContract.ObjectMeta) (edgecluster.ObjectMetaResolverContract, error) {
	return queryedgecluster.NewObjectMetaResolver(
		ctx,
		creator.logger,
		metadata)
}

// NewPortStatusResolverContract creates new instance of the PortStatusResolverContract, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// portStatus: Mandatory. Contains the object portStatus.
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewPortStatusResolverContract(
	ctx context.Context,
	portStatus *edgeclusterGrpcContract.PortStatus) (edgecluster.PortStatusResolverContract, error) {
	return queryedgecluster.NewPortStatusResolver(
		ctx,
		creator.logger,
		portStatus)
}

// NewLoadBalancerIngressResolver creates new instance of the LoadBalancerIngressResolverContract, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// loadBalancerIngress: Mandatory. The load balancer ingress details
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewLoadBalancerIngressResolver(
	ctx context.Context,
	ingress *edgeclusterGrpcContract.LoadBalancerIngress) (edgecluster.LoadBalancerIngressResolverContract, error) {
	return queryedgecluster.NewLoadBalancerIngressResolver(
		ctx,
		creator.logger,
		creator,
		ingress)
}

// NewLoadBalancerStatusResolver creates new instance of the LoadBalancerStatusResolverContract, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// loadBalancerIngress: Mandatory. The load balancer status details
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewLoadBalancerStatusResolver(
	ctx context.Context,
	loadBalancerStatus *edgeclusterGrpcContract.LoadBalancerStatus) (edgecluster.LoadBalancerStatusResolverContract, error) {
	return queryedgecluster.NewLoadBalancerStatusResolver(
		ctx,
		creator.logger,
		creator,
		loadBalancerStatus)
}

// NewServiceConditionResolver creates new instance of the ServiceConditionResolverContract, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// serviceCondition: Mandatory. Contains condition information for a service.
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewServiceConditionResolver(
	ctx context.Context,
	serviceCondition *edgeclusterGrpcContract.ServiceCondition) (edgecluster.ServiceConditionResolverContract, error) {
	return queryedgecluster.NewServiceConditionResolver(
		ctx,
		creator.logger,
		serviceCondition)
}
