// package types defines the different interfaces used in the GraphQL implementation
package types

import (
	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types/relay"
	"github.com/graph-gophers/graphql-go"
)

type UserTenantInputArgument struct {
	TenantID graphql.ID
}

type UserTenantsInputArgument struct {
	relay.ConnectionArgument
	TenantIDs  *[]graphql.ID
	SortOption *string
}

type UserEdgeClusterInputArgument struct {
	EdgeClusterID graphql.ID
}

type UserEdgeClustersInputArgument struct {
	relay.ConnectionArgument
	EdgeClusterIDs *[]graphql.ID
	SortOption     *string
}
