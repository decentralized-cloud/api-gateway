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
	// Returns the new tenant payload or error if something goes wrong
	MutateAndGetPayload(
		ctx context.Context,
		args CreateTenantInputArgument) (CreateTenantPayloadResolverContract, error)
}

// UpdateTenantContract declares the type to use when updating an existing tenant
type UpdateTenantContract interface {
	// MutateAndGetPayload update an existing tenant and returns the payload contains the result of updating an existing tenant
	// ctx: Mandatory. Reference to the context
	// args: Mandatory. Reference to the input argument contains tenant information to update
	// Returns the updated tenant payload or error if something goes wrong
	MutateAndGetPayload(
		ctx context.Context,
		args UpdateTenantInputArgument) (UpdateTenantPayloadResolverContract, error)
}

// DeleteTenantContract declares the type to use when updating an existing tenant
type DeleteTenantContract interface {
	// MutateAndGetPayload update an existing tenant and returns the payload contains the result of deleting an existing tenant
	// ctx: Mandatory. Reference to the context
	// args: Mandatory. Reference to the input argument contains tenant information to update
	// Returns the deleted tenant payload or error if something goes wrong
	MutateAndGetPayload(
		ctx context.Context,
		args DeleteTenantInputArgument) (DeleteTenantPayloadResolverContract, error)
}

// CreateEdgeClusterContract declares the type to use when creating a new edge cluster
type CreateEdgeClusterContract interface {
	// MutateAndGetPayload creates a new edge cluster and returns the payload contains the result of creating a new edge cluster
	// ctx: Mandatory. Reference to the context
	// args: Mandatory. Reference to the input argument contains edge cluster information to create
	// Returns the new edge cluster payload or error if something goes wrong
	MutateAndGetPayload(
		ctx context.Context,
		args CreateEdgeClusterInputArgument) (CreateEdgeClusterPayloadResolverContract, error)
}

// UpdateEdgeClusterContract declares the type to use when updating an existing edge cluster
type UpdateEdgeClusterContract interface {
	// MutateAndGetPayload update an existing edge cluster and returns the payload contains the result of updating an existing edge cluster
	// ctx: Mandatory. Reference to the context
	// args: Mandatory. Reference to the input argument contains edge cluster information to update
	// Returns the updated edge cluster payload or error if something goes wrong
	MutateAndGetPayload(
		ctx context.Context,
		args UpdateEdgeClusterInputArgument) (UpdateEdgeClusterPayloadResolverContract, error)
}

// DeleteEdgeClusterContract declares the type to use when updating an existing edge cluster
type DeleteEdgeClusterContract interface {
	// MutateAndGetPayload update an existing edge cluster and returns the payload contains the result of deleting an existing edge cluster
	// ctx: Mandatory. Reference to the context
	// args: Mandatory. Reference to the input argument contains edge cluster information to update
	// Returns the deleted edge cluster payload or error if something goes wrong
	MutateAndGetPayload(
		ctx context.Context,
		args DeleteEdgeClusterInputArgument) (DeleteEdgeClusterPayloadResolverContract, error)
}
