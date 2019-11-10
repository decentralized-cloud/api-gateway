// packae edgecluster implements used edge cluster related types in the GraphQL transport layer
package edgecluster

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types/relay"
	edgeclusterGrpcContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
	"github.com/graph-gophers/graphql-go"
)

type QueryResolverCreatorContract interface {
	// NewEdgeClusterResolver creates new EdgeClusterResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// edgeClusterID: Mandatory. The edge cluster unique identifier
	// edgeCluster: Optional. The edge clusterdetails
	// Returns the EdgeClusterResolverContract or error if something goes wrong
	NewEdgeClusterResolver(
		ctx context.Context,
		edgeClusterID string,
		edgeCluster *edgeclusterGrpcContract.EdgeCluster) (EdgeClusterResolverContract, error)

	// NewEdgeClusterTypeEdgeResolver creates new EdgeClusterTypeEdgeResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// edgeClusterID: Mandatory. The edge cluster unique identifier
	// edgeCluster: Optional. The edge cluster details
	// cursor: Mandatory. The cursor
	// Returns the EdgeClusterTypeEdgeResolverContract or error if something goes wrong
	NewEdgeClusterTypeEdgeResolver(
		ctx context.Context,
		edgeClusterID string,
		cursor string,
		edgeCluster *edgeclusterGrpcContract.EdgeCluster) (EdgeClusterTypeEdgeResolverContract, error)

	// NewEdgeClusterTypeConnectionResolver creates new EdgeClusterTypeConnectionResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// edgeClusters: Mandatory. Reference the list of edge clusters
	// hasPreviousPage: Mandatory. Indicates whether more edges exist prior to the set defined by the clients arguments
	// hasNextPage: Mandatory. Indicates whether more edges exist following the set defined by the clients arguments
	// Returns the EdgeClusterTypeConnectionResolverContract or error if something goes wrong
	NewEdgeClusterTypeConnectionResolver(
		ctx context.Context,
		edgeClusters []*edgeclusterGrpcContract.EdgeClusterWithCursor,
		hasPreviousPage bool,
		hasNextPage bool) (EdgeClusterTypeConnectionResolverContract, error)

	// NewEdgeClusterTenantResolver creates new EdgeClusterTenatnResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// tenantID: Mandatory. The tenant unique identifier
	// Returns the EdgeClusterTenatnResolverContract or error if something goes wrong
	NewEdgeClusterTenantResolver(
		ctx context.Context,
		tenantID string) (EdgeClusterTenantResolverContract, error)
}

// EdgeClusterResolverContract declares the resolver that can retrieve edge cluster information
type EdgeClusterResolverContract interface {
	// ID returns edge cluster unique identifier
	// ctx: Mandatory. Reference to the context
	// Returns the edge cluster unique identifier
	ID(ctx context.Context) graphql.ID

	// Name returns edge cluster name
	// ctx: Mandatory. Reference to the context
	// Returns the edge cluster name
	Name(ctx context.Context) string

	// ClusterPublicIPAddress returns edge cluster public IP address
	// ctx: Mandatory. Reference to the context
	// Returns the edge cluster public IP address
	ClusterPublicIPAddress(ctx context.Context) string

	// Tenant returns edge cluster tenant
	// ctx: Mandatory. Reference to the context
	// Returns the edge cluster tenant
	Tenant(ctx context.Context) (EdgeClusterTenantResolverContract, error)
}

// EdgeClusterTypeConnectionResolverContract declares the resolver that returns edge cluster edge compatible with graphql-relay
type EdgeClusterTypeConnectionResolverContract interface {
	// PageInfo returns the paging information compatible with graphql-relay
	// ctx: Mandatory. Reference to the context
	// Returns the paging information
	PageInfo(ctx context.Context) (relay.PageInfoResolverContract, error)

	// Edges returns the edge cluster edges compatible with graphql-relay
	// ctx: Mandatory. Reference to the context
	// Returns the edge cluster edges
	Edges(ctx context.Context) (*[]EdgeClusterTypeEdgeResolverContract, error)
}

// EdgeClusterTypeEdgeResolverContract declares the resolver that returns edge cluster edge compatible with graphql-relay
type EdgeClusterTypeEdgeResolverContract interface {
	// Node returns the edge cluster resolver
	// ctx: Mandatory. Reference to the context
	// Returns the edge cluster resolver or error if something goes wrong
	Node(ctx context.Context) (EdgeClusterResolverContract, error)

	// Cursor returns the cursor for the edge cluster edge compatible with graphql-relay
	// ctx: Mandatory. Reference to the context
	// Returns the cursor
	Cursor(ctx context.Context) string
}

// EdgeClusterTenantResolverContract declares the resolver that returns edge cluster tenant
type EdgeClusterTenantResolverContract interface {
	// ID returns tenant unique identifier
	// ctx: Mandatory. Reference to the context
	// Returns the tenant  unique identifier
	ID(ctx context.Context) graphql.ID

	// Name returns tenant name
	// ctx: Mandatory. Reference to the context
	// Returns the tenant name
	Name(ctx context.Context) string
}

type EdgeClusterClusterEdgeClusterInputArgument struct {
	EdgeClusterID graphql.ID
}

type EdgeClusterEdgeClustersInputArgument struct {
	relay.ConnectionArgument
	EdgeClusterIDs *[]graphql.ID
	SortOption     *string
}
