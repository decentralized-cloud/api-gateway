// package types defines the different interfaces used in the GraphQL implementation
package types

import (
	"context"
)

// RootResolverContract declares the root resolver
type RootResolverContract interface {
	// User returns user resolver
	// ctx: Mandatory. Reference to the context
	// Returns the user resolver or error if something goes wrong
	User(ctx context.Context) (UserResolverContract, error)

	// CreateTenant returns create tenant mutator
	// ctx: Mandatory. Reference to the context
	// Returns the create tenant mutator or error if something goes wrong
	CreateTenant(
		ctx context.Context,
		args CreateTenantInputArgument) (CreateTenantPayloadResolverContract, error)

	// UpdateTenant returns update tenant mutator
	// ctx: Mandatory. Reference to the context
	// Returns the update tenant mutator or error if something goes wrong
	UpdateTenant(
		ctx context.Context,
		args UpdateTenantInputArgument) (UpdateTenantPayloadResolverContract, error)
}
