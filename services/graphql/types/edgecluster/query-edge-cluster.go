// packae edgecluster implements used edge cluster related types in the GraphQL transport layer
package edgecluster

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/graphql/types/relay"
	"github.com/graph-gophers/graphql-go"
)

// EdgeClusterObjectMetadataResolverContract declares the standard edge cluster object's metadata.
type EdgeClusterObjectMetadataResolverContract interface {
	// ID returns the object unique identitfier
	// ctx: Mandatory. Reference to the context
	// Returns the object unique identitfier
	ID(ctx context.Context) graphql.ID

	// Name returns the object name
	// ctx: Mandatory. Reference to the context
	// Returns the object name
	Name(ctx context.Context) string

	// Namespace returns the object namespace
	// ctx: Mandatory. Reference to the context
	// Returns the object namespace
	Namespace(ctx context.Context) *string
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

	// ProvisionDetail returns edge cluster provisioning detail
	// ctx: Mandatory. Reference to the context
	// Returns the edge cluster provisioning detail resolver or error if something goes wrong.
	ProvisionDetail(ctx context.Context) (EdgeClusterProvisionDetailResolverContract, error)

	// Nodes returns the resolver that resolves the nodes that are part of the given edge cluster or error if something goes wrong.
	// ctx: Mandatory. Reference to the context
	// Returns the resolver that resolves the nodes that are part of the given edge cluster or error if something goes wrong.
	Nodes(ctx context.Context) ([]EdgeClusterNodeResolverContract, error)

	// Pods returns the resolver that resolves the pods that are part of the given edge cluster or error if something goes wrong.
	// ctx: Mandatory. Reference to the context
	// args: Mandatory. Reference to the query argument
	// Returns the resolver that resolves the pods that are part of the given edge cluster or error if something goes wrong.
	Pods(ctx context.Context, args EdgeClusterPodInputArgument) ([]EdgeClusterPodResolverContract, error)
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

// IngressResolverContract declares the resolver that returns Ingress
type IngressResolverContract interface {
	// IP is set for load-balancer ingress points that are IP based
	// (typically GCE or OpenStack load-balancers)
	// ctx: Mandatory. Reference to the context
	// Returns the IP is set for load-balancer ingress points that are IP based
	IP(ctx context.Context) *string

	// Hostname is set for load-balancer ingress points that are DNS based
	// (typically AWS load-balancers)
	// ctx: Mandatory. Reference to the context
	// Returns the Hostname is set for load-balancer ingress points that are DNS based
	Hostname(ctx context.Context) *string
}

// PortResolverContract declares the resolver that returns Port
type PortResolverContract interface {
	// Port returns the port number of the edge-cluster master port of which status is recorded here
	// ctx: Mandatory. Reference to the context
	// Returns the port number of the edge-cluster master port of which status is recorded here
	Port(ctx context.Context) int32

	// Protocol returns the protocol of the edge-cluster master port of which status is recorded here
	// ctx: Mandatory. Reference to the context
	// Returns the protocol of the edge-cluster master port of which status is recorded here
	Protocol(ctx context.Context) string
}

// EdgeClusterProvisionDetailResolverContract declares the resolver that returns edge cluster provisioning details
type EdgeClusterProvisionDetailResolverContract interface {
	// Ingress returns the ingress details of the edge cluster master node
	// ctx: Mandatory. Reference to the context
	// Returns the ingress details of the edge cluster master node resolvers or error if something goes wrong.
	Ingress(ctx context.Context) ([]IngressResolverContract, error)

	// Ports is a list of records of edge-cluster master ports
	// ctx: Mandatory. Reference to the context
	// Returns the Ports is a list of records of edge-cluster master ports resolvers or error if something goes wrong.
	Ports(ctx context.Context) ([]PortResolverContract, error)

	// KubeconfigContent returns the edge cluster Kubeconfig content
	// ctx: Mandatory. Reference to the context
	// Returns the edge cluster Kubeconfig content
	KubeconfigContent(ctx context.Context) *string
}
