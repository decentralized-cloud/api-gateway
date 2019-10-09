// package types defines the different interfaces used in the GraphQL implementation
package types

import "github.com/graph-gophers/graphql-go"

type UserTenantInputArgument struct {
	TenantID graphql.ID
}

type UserTenantsInputArgument struct {
	ConnectionArgument
	TenantIDs  *[]graphql.ID
	SortOption *string
}

type UserEdgeClusterInputArgument struct {
	EdgeClusterID graphql.ID
}

type UserEdgeClustersInputArgument struct {
	ConnectionArgument
	EdgeClusterIDs *[]graphql.ID
	SortOption     *string
}

type TenantClusterEdgeClusterInputArgument struct {
	EdgeClusterID graphql.ID
}

type TenantEdgeClustersInputArgument struct {
	ConnectionArgument
	EdgeClusterIDs *[]graphql.ID
	SortOption     *string
}
