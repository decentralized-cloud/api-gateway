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

	// User returns user resolver
	// ctx: Mandatory. Reference to the context
	// Returns the user resolver or error if something goes wrong
	CreateTenant(
		ctx context.Context,
		args CreateTenantInputArgument) (CreateTenantPayloadResolverContract, error)
}
