// Package graphql implements functions to expose api-gateway service endpoint using GraphQL protocol.
package graphql

import (
	"context"

	queryedgecluster "github.com/decentralized-cloud/api-gateway/services/graphql/query/edgecluster"
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/edgecluster"
	edgeclusterGrpcContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
)

// NewEdgeClusterResolver creates new EdgeClusterResolverContract and returns it
// ctx: Mandatory. Reference to the context
// edgeClusterID: Mandatory. The edge cluster unique identifier
// edgeClusterDetail: Optional. The edge cluster details, if provided, the value be used instead of contacting  the edge cluster service
// Returns the EdgeClusterResolverContract or error if something goes wrong
func (creator *resolverCreator) NewEdgeClusterResolver(
	ctx context.Context,
	edgeClusterID string,
	edgeClusterDetail *edgecluster.EdgeClusterDetail) (edgecluster.EdgeClusterResolverContract, error) {
	return queryedgecluster.NewEdgeClusterResolver(
		ctx,
		creator,
		creator.logger,
		creator.edgeClusterClientService,
		edgeClusterID,
		edgeClusterDetail)
}

// NewEdgeClusterTypeEdgeResolver creates new EdgeClusterTypeEdgeResolverContract and returns it
// ctx: Mandatory. Reference to the context
// edgeClusterID: Mandatory. The edge cluster unique identifier
// cursor: Mandatory. The cursor
// edgeClusterDetail: Optional. The edge cluster details, if provided, the value be used instead of contacting  the edge cluster service
// Returns the EdgeClusterTypeEdgeResolverContract or error if something goes wrong
func (creator *resolverCreator) NewEdgeClusterTypeEdgeResolver(
	ctx context.Context,
	edgeClusterID string,
	cursor string,
	edgeClusterDetail *edgecluster.EdgeClusterDetail) (edgecluster.EdgeClusterTypeEdgeResolverContract, error) {
	return queryedgecluster.NewEdgeClusterTypeEdgeResolver(
		ctx,
		creator,
		edgeClusterID,
		edgeClusterDetail,
		cursor)
}

// NewEdgeClusterTypeConnectionResolver creates new EdgeClusterTypeConnectionResolverContract and returns it
// ctx: Mandatory. Reference to the context
// edgeClusters: Mandatory. Reference the list of edge clusters
// hasPreviousPage: Mandatory. Indicates whether more edges exist prior to the set defined by the clients arguments
// hasNextPage: Mandatory. Indicates whether more edges exist following the set defined by the clients arguments
// totalCount: Mandatory. The total count of matched edge clusters
// Returns the EdgeClusterTypeConnectionResolverContract or error if something goes wrong
func (creator *resolverCreator) NewEdgeClusterTypeConnectionResolver(
	ctx context.Context,
	edgeclusters []*edgeclusterGrpcContract.EdgeClusterWithCursor,
	hasPreviousPage bool,
	hasNextPage bool,
	totalCount int32) (edgecluster.EdgeClusterTypeConnectionResolverContract, error) {
	return queryedgecluster.NewEdgeClusterTypeConnectionResolver(
		ctx,
		creator,
		edgeclusters,
		hasPreviousPage,
		hasNextPage,
		totalCount)
}

// NewEdgeClusterProjectResolver creates new EdgeClusterTenatnResolverContract and returns it
// ctx: Mandatory. Reference to the context
// projectID: Mandatory. The project unique identifier
// Returns the EdgeClusterTenatnResolverContract or error if something goes wrong
func (creator *resolverCreator) NewEdgeClusterProjectResolver(
	ctx context.Context,
	projectID string) (edgecluster.EdgeClusterProjectResolverContract, error) {
	return queryedgecluster.NewEdgeClusterProjectResolver(
		ctx,
		creator,
		creator.logger,
		projectID)
}

// NewProvisionDetailsResolver creates new ProvisionDetailsResolverContract and returns it
// ctx: Mandatory. Reference to the context
// ProvisionDetails: Mandatory. The edge cluster provisioning details
// Returns the ProvisionDetailsResolverContract or error if something goes wrong
func (creator *resolverCreator) NewProvisionDetailsResolver(
	ctx context.Context,
	provisionDetails *edgeclusterGrpcContract.ProvisionDetail) (edgecluster.ProvisionDetailsResolverContract, error) {
	return queryedgecluster.NewProvisionDetailsResolver(
		ctx,
		creator.logger,
		creator,
		provisionDetails)
}
