// package types defines the different interfaces used in the GraphQL implementation
package types

import (
	"context"
)

// CreateTenantPayloadResolverContract declares the resolver that can returns the payload contains the result of creating a new tenant
type CreateTenantPayloadResolverContract interface {
	// Tenant returns the new tenant inforamtion
	// ctx: Mandatory. Reference to the context
	// Returns the new tenant inforamtion
	Tenant(ctx context.Context) (TenantTypeEdgeResolverContract, error)

	// ClientMutationId returns the client mutation ID that was provided as part of the mutation request
	// ctx: Mandatory. Reference to the context
	// Returns the provided clientMutationId as part of mutation request
	ClientMutationId(ctx context.Context) *string
}

// UpdateTenantPayloadResolverContract declares the resolver that can returns the payload contains the result of updating an existing tenant
type UpdateTenantPayloadResolverContract interface {
	// Tenant returns the updated tenant inforamtion
	// ctx: Mandatory. Reference to the context
	// Returns the updated tenant inforamtion
	Tenant(ctx context.Context) (TenantTypeEdgeResolverContract, error)

	// ClientMutationId returns the client mutation ID that was provided as part of the mutation request
	// ctx: Mandatory. Reference to the context
	// Returns the provided clientMutationId as part of mutation request
	ClientMutationId(ctx context.Context) *string
}
