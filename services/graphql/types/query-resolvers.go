// Package types defines the different interfaces used in the GraphQL implementation
package types

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/graphql/types/edgecluster"
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/tenant"
	"github.com/graph-gophers/graphql-go"
)

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
		args UserTenantInputArgument) (tenant.TenantResolverContract, error)

	// Tenants returns tenant connection compatible with graphql-relay
	// ctx: Mandatory. Reference to the context
	// args: Mandatory. The argument list
	// Returns the tenant resolver or error if something goes wrong
	Tenants(
		ctx context.Context,
		args UserTenantsInputArgument) (tenant.TenantTypeConnectionResolverContract, error)

	// EdgeCluster returns tenant resolver
	// ctx: Mandatory. Reference to the context
	// args: Mandatory. The argument list
	// Returns the tenant resolver or error if something goes wrong
	EdgeCluster(
		ctx context.Context,
		args UserEdgeClusterInputArgument) (edgecluster.EdgeClusterResolverContract, error)

	// EdgeClusters returns tenant connection compatible with graphql-relay
	// ctx: Mandatory. Reference to the context
	// args: Mandatory. The argument list
	// Returns the tenant resolver or error if something goes wrong
	EdgeClusters(
		ctx context.Context,
		args UserEdgeClustersInputArgument) (edgecluster.EdgeClusterTypeConnectionResolverContract, error)
}
