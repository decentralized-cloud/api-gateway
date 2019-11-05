// Package types defines the different interfaces used in the GraphQL implementation
package types

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types/edgecluster"
	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types/relay"
	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types/tenant"
)

// ResolverCreatorContract declares the service that can create different resolvers
type ResolverCreatorContract interface {
	// NewRootResolver creates new RootResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// Returns the RootResolverContract or error if something goes wrong
	NewRootResolver(ctx context.Context) (RootResolverContract, error)

	// NewUserResolver creates new UserResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// userID: Mandatory. The user unique identifier
	// Returns the UserResolverContract or error if something goes wrong
	NewUserResolver(
		ctx context.Context,
		userID string) (UserResolverContract, error)

	relay.PageInfoResolverCreatorContract
	tenant.QueryResolverCreatorContract
	tenant.MutationResolverCreatorContract
	edgecluster.QueryResolverCreatorContract
	edgecluster.MutationResolverCreatorContract
}
