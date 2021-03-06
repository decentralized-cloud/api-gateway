// Package types defines the different interfaces used in the GraphQL implementation
package types

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/graphql/types/edgecluster"
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/project"
)

// RootResolverContract declares the root resolver
type RootResolverContract interface {
	// User returns user resolver
	// ctx: Mandatory. Reference to the context
	// Returns the user resolver or error if something goes wrong
	User(ctx context.Context) (UserResolverContract, error)

	project.RootResolverContract
	edgecluster.RootResolverContract
}
