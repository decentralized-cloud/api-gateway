// package types defines the different interfaces used in the GraphQL implementation
package types

import (
	"context"

	"github.com/graph-gophers/graphql-go"
)

// PageInfoResolverContract declares the resolver that returns paging information compatible with graphql-relay specification
type PageInfoResolverContract interface {
	// StartCursor returns start cursor
	// ctx: Mandatory. Reference to the context
	// Returns the start cursor
	StartCursor(ctx context.Context) *string

	// EndCursor returns end cursor
	// ctx: Mandatory. Reference to the context
	// Returns the end cursor
	EndCursor(ctx context.Context) *string

	// HasNextPage indicates whether returned page has next page to be retrieved
	// ctx: Mandatory. Reference to the context
	// Returns the value indicates whether returned page has next page to be retrieved
	HasNextPage(ctx context.Context) bool

	// HasPreviousPage indicates whether returned page has previous page to be retrieved
	// ctx: Mandatory. Reference to the context
	// Returns the value indicates whether returned page has previous page to be retrieved
	HasPreviousPage(ctx context.Context) bool
}

// UserResolverContract declares the resolver that can retrieve user information
type UserResolverContract interface {
	// ID returns user unique identifier
	// ctx: Mandatory. Reference to the context
	// Returns the user unique identifier
	ID(ctx context.Context) graphql.ID

	// Tenant returns tenant resolver
	// ctx: Mandatory. Reference to the context
	// args: Mandatory. The argument list
	// Returns the tenant resolver or error if something goes wrong
	Tenant(
		ctx context.Context,
		args UserTenantInputArgument) (TenantResolverContract, error)

	// Tenants returns tenant conenction compatible with graphql-relay
	// ctx: Mandatory. Reference to the context
	// args: Mandatory. The argument list
	// Returns the tenant resolver or error if something goes wrong
	Tenants(
		ctx context.Context,
		args UserTenantsInputArgument) (TenantTypeConnectionResolverContract, error)

	// EdgeCluster returns tenant resolver
	// ctx: Mandatory. Reference to the context
	// args: Mandatory. The argument list
	// Returns the tenant resolver or error if something goes wrong
	EdgeCluster(
		ctx context.Context,
		args UserEdgeClusterInputArgument) (EdgeClusterResolverContract, error)

	// EdgeClusters returns tenant conenction compatible with graphql-relay
	// ctx: Mandatory. Reference to the context
	// args: Mandatory. The argument list
	// Returns the tenant resolver or error if something goes wrong
	EdgeClusters(
		ctx context.Context,
		args UserEdgeClustersInputArgument) (EdgeClusterTypeConnectionResolverContract, error)
}

// TenantResolverContract declares the resolver that can retrieve tenant information
type TenantResolverContract interface {
	// ID returns tenant unique identifier
	// ctx: Mandatory. Reference to the context
	// Returns the tenant unique identifier
	ID(ctx context.Context) graphql.ID

	// Name returns tenant name
	// ctx: Mandatory. Reference to the context
	// Returns the tenant name
	Name(ctx context.Context) string

	// EdgeCluster returns tenant resolver
	// ctx: Mandatory. Reference to the context
	// args: Mandatory. The argument list
	// Returns the tenant resolver or error if something goes wrong
	EdgeCluster(
		ctx context.Context,
		args TenantClusterEdgeClusterInputArgument) (EdgeClusterResolverContract, error)

	// EdgeClusters returns tenant conenction compatible with graphql-relay
	// ctx: Mandatory. Reference to the context
	// args: Mandatory. The argument list
	// Returns the tenant resolver or error if something goes wrong
	EdgeClusters(
		ctx context.Context,
		args TenantEdgeClustersInputArgument) (EdgeClusterTypeConnectionResolverContract, error)
}

// TenantTypeConnectionResolverContract declares the resolver that returns tenant edge compatible with graphql-relay
type TenantTypeConnectionResolverContract interface {
	// PageInfo returns the paging information compatible with graphql-relay
	// ctx: Mandatory. Reference to the context
	// Returns the paging information
	PageInfo(ctx context.Context) (PageInfoResolverContract, error)

	// Edges returns the tenant edges compatible with graphql-relay
	// ctx: Mandatory. Reference to the context
	// Returns the tenant edges
	Edges(ctx context.Context) (*[]TenantTypeEdgeResolverContract, error)
}

// TenantTypeEdgeResolverContract declares the resolver that returns tenant edge compatible with graphql-relay
type TenantTypeEdgeResolverContract interface {
	// Node returns the tenant resolver
	// ctx: Mandatory. Reference to the context
	// Returns the tenant resolver or error if something goes wrong
	Node(ctx context.Context) (TenantResolverContract, error)

	// Cursor returns the cursor for the tenant edge compatible with graphql-relay
	// ctx: Mandatory. Reference to the context
	// Returns the cursor
	Cursor(ctx context.Context) string
}

// EdgeClusterResolverContract declares the resolver that can retrieve edge-cluster information
type EdgeClusterResolverContract interface {
	// ID returns edge-cluster unique identifier
	// ctx: Mandatory. Reference to the context
	// Returns the edge-cluster unique identifier
	ID(ctx context.Context) graphql.ID

	// Name returns edge-cluster name
	// ctx: Mandatory. Reference to the context
	// Returns the edge-cluster name
	Name(ctx context.Context) string
}

// EdgeClusterTypeConnectionResolverContract declares the resolver that returns edge-cluster edge compatible with graphql-relay
type EdgeClusterTypeConnectionResolverContract interface {
	// PageInfo returns the paging information compatible with graphql-relay
	// ctx: Mandatory. Reference to the context
	// Returns the paging information
	PageInfo(ctx context.Context) (PageInfoResolverContract, error)

	// Edges returns the edge-cluster edges compatible with graphql-relay
	// ctx: Mandatory. Reference to the context
	// Returns the edge-cluster edges
	Edges(ctx context.Context) (*[]EdgeClusterTypeEdgeResolverContract, error)
}

// EdgeClusterTypeEdgeResolverContract declares the resolver that returns edge-cluster edge compatible with graphql-relay
type EdgeClusterTypeEdgeResolverContract interface {
	// Node returns the edge-cluster resolver
	// ctx: Mandatory. Reference to the context
	// Returns the edge-cluster resolver or error if something goes wrong
	Node(ctx context.Context) (EdgeClusterResolverContract, error)

	// Cursor returns the cursor for the edge-cluster edge compatible with graphql-relay
	// ctx: Mandatory. Reference to the context
	// Returns the cursor
	Cursor(ctx context.Context) string
}
