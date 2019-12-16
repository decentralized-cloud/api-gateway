// Package types defines the different interfaces used in the GraphQL implementation
package types

import (
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/relay"
	"github.com/graph-gophers/graphql-go"
)

type SortingOptionPair struct {
	Name      string
	Direction string
}

type UserTenantInputArgument struct {
	TenantID string
}

type UserTenantsInputArgument struct {
	relay.ConnectionArgument
	SortingOptions *[]SortingOptionPair
	TenantIDs      *[]graphql.ID
}

type UserEdgeClusterInputArgument struct {
	EdgeClusterID graphql.ID
}

type UserEdgeClustersInputArgument struct {
	relay.ConnectionArgument
	SortingOptions *[]SortingOptionPair
	EdgeClusterIDs *[]graphql.ID
	TenantIDs      *[]graphql.ID
}
