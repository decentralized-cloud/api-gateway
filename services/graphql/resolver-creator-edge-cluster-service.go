// Package graphql implements functions to expose api-gateway service endpoint using GraphQL protocol.
package graphql

import (
	"context"

	queryedgecluster "github.com/decentralized-cloud/api-gateway/services/graphql/query/edgecluster"
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/edgecluster"
	edgeclusterGrpcContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
)

// NewEdgeClusterServiceResolver creates new instance of the ServiceResolverContract, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// service: Mandatory. Contains information about the edge cluster service
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewEdgeClusterServiceResolver(
	ctx context.Context,
	service *edgeclusterGrpcContract.EdgeClusterService) (edgecluster.ServiceResolverContract, error) {
	return queryedgecluster.NewEdgeClusterServiceResolver(
		ctx,
		creator.logger,
		creator,
		service)
}

// NewServiceStatusResolver creates new instance of the ServiceStatusResolverContract, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// serviceStatus: Mandatory. Contains information about the service status
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewServiceStatusResolver(
	ctx context.Context,
	serviceStatus *edgeclusterGrpcContract.ServiceStatus) (edgecluster.ServiceStatusResolverContract, error) {
	return queryedgecluster.NewServiceStatusResolver(
		ctx,
		creator.logger,
		creator,
		serviceStatus)
}

// NewServicePortResolver creates new instance of the ServicePortResolverContract, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// servicePort: Mandatory. Contains information about the service port
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewServicePortResolver(
	ctx context.Context,
	servicePort *edgeclusterGrpcContract.ServicePort) (edgecluster.ServicePortResolverContract, error) {
	return queryedgecluster.NewServicePortResolver(
		ctx,
		creator.logger,
		servicePort)
}

// NewServiceSpecResolver creates new instance of the ServiceSpecResolverContract, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// serviceSpec: Mandatory. Contains spec information for a service.
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewServiceSpecResolver(
	ctx context.Context,
	serviceSpec *edgeclusterGrpcContract.ServiceSpec) (edgecluster.ServiceSpecResolverContract, error) {
	return queryedgecluster.NewServiceSpecResolver(
		ctx,
		creator.logger,
		creator,
		serviceSpec)
}
