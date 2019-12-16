// packae tenant implements used tenant related types in the GraphQL transport layer
package tenant

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/graphql/types/edgecluster"
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/relay"
	tenantGrpcContract "github.com/decentralized-cloud/tenant/contract/grpc/go"
	"github.com/graph-gophers/graphql-go"
)

type QueryResolverCreatorContract interface {
	// NewTenantResolver creates new TenantResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// tenantID: Mandatory. The tenant unique identifier
	// tenantDetail: Optional. The tennat details, if provided, the value be used instead of contacting  the edge cluster service
	// Returns the TenantResolverContract or error if something goes wrong
	NewTenantResolver(
		ctx context.Context,
		tenantID string,
		tenantDetail *TenantDetail) (TenantResolverContract, error)

	// NewTenantTypeEdgeResolver creates new TenantTypeEdgeResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// tenantID: Mandatory. The tenant unique identifier
	// tenantDetail: Optional. The tennat details, if provided, the value be used instead of contacting  the edge cluster service
	// cursor: Mandatory. The cursor
	// Returns the TenantTypeEdgeResolverContract or error if something goes wrong
	NewTenantTypeEdgeResolver(
		ctx context.Context,
		tenantID string,
		cursor string,
		tenantDetail *TenantDetail) (TenantTypeEdgeResolverContract, error)

	// NewTenantTypeConnectionResolver creates new TenantTypeConnectionResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// tenants: Mandatory. Reference the list of tenants
	// hasPreviousPage: Mandatory. Indicates whether more edges exist prior to the set defined by the clients arguments
	// hasNextPage: Mandatory. Indicates whether more edges exist following the set defined by the clients arguments
	// totalCount: Mandatory. The total count of matched tenants
	// Returns the TenantTypeConnectionResolverContract or error if something goes wrong
	NewTenantTypeConnectionResolver(
		ctx context.Context,
		tenants []*tenantGrpcContract.TenantWithCursor,
		hasPreviousPage bool,
		hasNextPage bool,
		totalCount int32) (TenantTypeConnectionResolverContract, error)
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
		args TenantClusterEdgeClusterInputArgument) (edgecluster.EdgeClusterResolverContract, error)

	// EdgeClusters returns tenant connection compatible with graphql-relay
	// ctx: Mandatory. Reference to the context
	// args: Mandatory. The argument list
	// Returns the tenant resolver or error if something goes wrong
	EdgeClusters(
		ctx context.Context,
		args TenantEdgeClustersInputArgument) (edgecluster.EdgeClusterTypeConnectionResolverContract, error)
}

// TenantTypeConnectionResolverContract declares the resolver that returns tenant edge compatible with graphql-relay
type TenantTypeConnectionResolverContract interface {
	// PageInfo returns the paging information compatible with graphql-relay
	// ctx: Mandatory. Reference to the context
	// Returns the paging information
	PageInfo(ctx context.Context) (relay.PageInfoResolverContract, error)

	// Edges returns the tenant edges compatible with graphql-relay
	// ctx: Mandatory. Reference to the context
	// Returns the tenant edges
	Edges(ctx context.Context) (*[]TenantTypeEdgeResolverContract, error)

	// TotalCount returns total count of the matched tenants
	// ctx: Mandatory. Reference to the context
	// Returns the total count of the matched tenants
	TotalCount(ctx context.Context) *int32
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

type TenantDetail struct {
	Tenant *tenantGrpcContract.Tenant
}

type TenantClusterEdgeClusterInputArgument struct {
	EdgeClusterID graphql.ID
}

type TenantEdgeClustersInputArgument struct {
	relay.ConnectionArgument
	EdgeClusterIDs *[]graphql.ID
	SortOption     *string
}
