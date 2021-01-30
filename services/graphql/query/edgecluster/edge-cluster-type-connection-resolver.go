// Package edgecluster implements different edge cluster GraphQL query resovlers required by the GraphQL transport layer
package edgecluster

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/edgecluster"
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/relay"
	edgeclusterGrpcContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"github.com/thoas/go-funk"
)

type edgeClusterTypeConnectionResolver struct {
	resolverCreator types.ResolverCreatorContract
	edgeClusters    []*edgeclusterGrpcContract.EdgeClusterWithCursor
	hasPreviousPage bool
	hasNextPage     bool
	totalCount      int32
}

// NewEdgeClusterTypeConnectionResolver creates new instance of the edgeClusterTypeConnectionResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can create new instances of resolvers
// edgeClusters: Mandatory. Reference the list of edge clusters
// hasPreviousPage: Mandatory. Indicates whether more edges exist prior to the set defined by the clients arguments
// hasNextPage: Mandatory. Indicates whether more edges exist following the set defined by the clients arguments
// totalCount: Mandatory. The total count of matched edge clusters
// Returns the new instance or error if something goes wrong
func NewEdgeClusterTypeConnectionResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	edgeClusters []*edgeclusterGrpcContract.EdgeClusterWithCursor,
	hasPreviousPage bool,
	hasNextPage bool,
	totalCount int32) (edgecluster.EdgeClusterTypeConnectionResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	return &edgeClusterTypeConnectionResolver{
		resolverCreator: resolverCreator,
		edgeClusters:    edgeClusters,
		hasPreviousPage: hasPreviousPage,
		hasNextPage:     hasNextPage,
		totalCount:      totalCount,
	}, nil
}

// PageInfo returns the paging information compatible with graphql-relay
// ctx: Mandatory. Reference to the context
// Returns the paging information
func (r *edgeClusterTypeConnectionResolver) PageInfo(ctx context.Context) (relay.PageInfoResolverContract, error) {
	var startCursor, endCursor string

	if len(r.edgeClusters) > 0 {
		startCursor = r.edgeClusters[0].Cursor
		endCursor = r.edgeClusters[len(r.edgeClusters)-1].Cursor
	}

	return r.resolverCreator.NewPageInfoResolver(
		ctx,
		&startCursor,
		&endCursor,
		r.hasNextPage,
		r.hasPreviousPage)
}

// Edges returns the edge cluster edges compatible with graphql-relay
// ctx: Mandatory. Reference to the context
// Returns the edge cluster edges
func (r *edgeClusterTypeConnectionResolver) Edges(ctx context.Context) (*[]edgecluster.EdgeClusterTypeEdgeResolverContract, error) {
	edgeClusters := funk.Filter(r.edgeClusters, func(edgeCluster *edgeclusterGrpcContract.EdgeClusterWithCursor) bool {
		return edgeCluster != nil
	}).([]*edgeclusterGrpcContract.EdgeClusterWithCursor)

	edges := []edgecluster.EdgeClusterTypeEdgeResolverContract{}

	for _, item := range edgeClusters {
		edge, err := r.resolverCreator.NewEdgeClusterTypeEdgeResolver(
			ctx,
			item.EdgeClusterID,
			item.Cursor,
			&edgecluster.EdgeClusterDetail{
				EdgeCluster:     item.EdgeCluster,
				ProvisionDetail: item.ProvisionDetail,
			})

		if err != nil {
			return nil, err
		}

		edges = append(edges, edge)
	}

	return &edges, nil
}

// TotalCount returns total count of the matched edge clusters
// ctx: Mandatory. Reference to the context
// Returns the total count of the matched edge cluster
func (r *edgeClusterTypeConnectionResolver) TotalCount(ctx context.Context) *int32 {
	return &r.totalCount
}
