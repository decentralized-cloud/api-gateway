// packae edgecluster implements used edge cluster related types in the GraphQL transport layer
package edgecluster

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/graphql/types/relay"
	edgeclusterGrpcContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
	"github.com/graph-gophers/graphql-go"
)

type EdgeClusterResolverCreatorContract interface {
	// NewEdgeClusterResolver creates new EdgeClusterResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// edgeClusterID: Mandatory. The edge cluster unique identifier
	// edgeClusterDetail: Optional. The edge cluster details, if provided, the value be used instead of contacting  the edge cluster service
	// Returns the EdgeClusterResolverContract or error if something goes wrong
	NewEdgeClusterResolver(
		ctx context.Context,
		edgeClusterID string,
		edgeClusterDetail *EdgeClusterDetail) (EdgeClusterResolverContract, error)

	// NewEdgeClusterTypeConnectionResolver creates new EdgeClusterTypeConnectionResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// edgeClusters: Mandatory. Reference the list of edge clusters
	// hasPreviousPage: Mandatory. Indicates whether more edges exist prior to the set defined by the clients arguments
	// hasNextPage: Mandatory. Indicates whether more edges exist following the set defined by the clients arguments
	// totalCount: Mandatory. The total count of matched edge clusters
	// Returns the EdgeClusterTypeConnectionResolverContract or error if something goes wrong
	NewEdgeClusterTypeConnectionResolver(
		ctx context.Context,
		edgeClusters []*edgeclusterGrpcContract.EdgeClusterWithCursor,
		hasPreviousPage bool,
		hasNextPage bool,
		totalCount int32) (EdgeClusterTypeConnectionResolverContract, error)

	// NewEdgeClusterTypeEdgeResolver creates new EdgeClusterTypeEdgeResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// edgeClusterID: Mandatory. The edge cluster unique identifier
	// cursor: Mandatory. The cursor
	// edgeClusterDetail: Optional. The edge cluster details, if provided, the value be used instead of contacting  the edge cluster service
	// Returns the EdgeClusterTypeEdgeResolverContract or error if something goes wrong
	NewEdgeClusterTypeEdgeResolver(
		ctx context.Context,
		edgeClusterID string,
		cursor string,
		edgeClusterDetail *EdgeClusterDetail) (EdgeClusterTypeEdgeResolverContract, error)

	// NewEdgeClusterProjectResolver creates new EdgeClusterTenatnResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// projectID: Mandatory. The project unique identifier
	// Returns the EdgeClusterTenatnResolverContract or error if something goes wrong
	NewEdgeClusterProjectResolver(
		ctx context.Context,
		projectID string) (EdgeClusterProjectResolverContract, error)

	// NewProvisionDetailsResolver creates new ProvisionDetailsResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// provisionDetails: Optional. The edge cluster provisioning details
	// Returns the ProvisionDetailsResolverContract or error if something goes wrong
	NewProvisionDetailsResolver(
		ctx context.Context,
		provisionDetails *edgeclusterGrpcContract.ProvisionDetail) (ProvisionDetailsResolverContract, error)
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

	// ClusterSecret returns edge cluster secret
	// ctx: Mandatory. Reference to the context
	// Returns the edge cluster secret
	ClusterSecret(ctx context.Context) string

	// ClusterType returns the edge cluster current type
	// ctx: Mandatory. Reference to the context
	// Returns the edge cluster current type or error if something went wrong
	ClusterType(ctx context.Context) (string, error)

	// Project returns edge cluster project
	// ctx: Mandatory. Reference to the context
	// Returns the edge cluster project resolver or error if something goes wrong.
	Project(ctx context.Context) (EdgeClusterProjectResolverContract, error)

	// ProvisionDetails returns edge cluster provisioning detail
	// ctx: Mandatory. Reference to the context
	// Returns the edge cluster provisioning detail resolver or error if something goes wrong.
	ProvisionDetails(ctx context.Context) (ProvisionDetailsResolverContract, error)

	// Nodes returns the resolver that resolves the nodes that are part of the given edge cluster or error if something goes wrong.
	// ctx: Mandatory. Reference to the context
	// Returns the resolver that resolves the nodes that are part of the given edge cluster or error if something goes wrong.
	Nodes(ctx context.Context) ([]NodeResolverContract, error)

	// Pods returns the resolver that resolves the pods that are part of the given edge cluster or error if something goes wrong.
	// ctx: Mandatory. Reference to the context
	// args: Mandatory. Reference to the query argument
	// Returns the resolver that resolves the pods that are part of the given edge cluster or error if something goes wrong.
	Pods(ctx context.Context, args EdgeClusterPodInputArgument) ([]PodResolverContract, error)

	// Services returns the resolver that resolves the services that are part of the given edge cluster or error if something goes wrong.
	// ctx: Mandatory. Reference to the context
	// args: Mandatory. Reference to the query argument
	// Returns the resolver that resolves the services that are part of the given edge cluster or error if something goes wrong.
	Services(ctx context.Context, args EdgeClusterServiceInputArgument) ([]ServiceResolverContract, error)
}

// EdgeClusterTypeConnectionResolverContract declares the resolver that returns edge cluster edge compatible with graphql-relay
type EdgeClusterTypeConnectionResolverContract interface {
	// PageInfo returns the paging information compatible with graphql-relay
	// ctx: Mandatory. Reference to the context
	// Returns the paging information resolver or error if something goes wrong.
	PageInfo(ctx context.Context) (relay.PageInfoResolverContract, error)

	// Edges returns the edge cluster edges compatible with graphql-relay
	// ctx: Mandatory. Reference to the context
	// Returns the edge cluster edges resolver or error if something goes wrong.
	Edges(ctx context.Context) (*[]EdgeClusterTypeEdgeResolverContract, error)

	// TotalCount returns total count of the matched edge clusters
	// ctx: Mandatory. Reference to the context
	// Returns the total count of the matched edge cluster
	TotalCount(ctx context.Context) *int32
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

// EdgeClusterProjectResolverContract declares the resolver that returns edge cluster project
type EdgeClusterProjectResolverContract interface {
	// ID returns project unique identifier
	// ctx: Mandatory. Reference to the context
	// Returns the project  unique identifier
	ID(ctx context.Context) graphql.ID

	// Name returns project name
	// ctx: Mandatory. Reference to the context
	// Returns the project name
	Name(ctx context.Context) string
}

// ProvisionDetailsResolverContract declares the resolver that returns edge cluster provisioning details
type ProvisionDetailsResolverContract interface {
	// LoadBalancer contains the current status of the load-balancer
	// ctx: Mandatory. Reference to the context
	// Returns the load balancer status resolver or error if something goes wrong.
	LoadBalancer(ctx context.Context) (LoadBalancerStatusResolverContract, error)

	// KubeconfigContent returns the edge cluster Kubeconfig content
	// ctx: Mandatory. Reference to the context
	// Returns the edge cluster Kubeconfig content
	KubeconfigContent(ctx context.Context) *string
}
