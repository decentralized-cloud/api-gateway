// package types defines the different interfaces used in the GraphQL implementation
package types

import (
	"context"
)

// CreateTenantContract declares the type to use when creating a new tenant
type CreateTenantContract interface {
	// MutateAndGetPayload creates a new tenant and returns the payload contains the result of creating a new tenant
	// ctx: Mandatory. Reference to the context
	// args: Mandatory. Reference to the input argument contains tenant information to create
	// Returns the new tennt payload or error if something goes wrong
	MutateAndGetPayload(
		ctx context.Context,
		args CreateTenantInputArgument) (CreateTenantPayloadResolverContract, error)
}
