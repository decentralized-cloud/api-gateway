// packae tenant implements used tenant related types in the GraphQL transport layer
package tenant

import (
	"context"

	"github.com/graph-gophers/graphql-go"
)

type MutationResolverCreatorContract interface {
	// NewCreateTenant creates new instance of the CreateTenantContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// Returns the new instance or error if something goes wrong
	NewCreateTenant(ctx context.Context) (CreateTenantContract, error)

	// NewCreateTenantPayloadResolver creates new instance of the CreateTenantPayloadResolverContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// Returns the new instance or error if something goes wrong
	NewCreateTenantPayloadResolver(
		ctx context.Context,
		clientMutationId *string) (CreateTenantPayloadResolverContract, error)

	// NewUpdateTenant creates new instance of the UpdateTenantContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// Returns the new instance or error if something goes wrong
	NewUpdateTenant(ctx context.Context) (UpdateTenantContract, error)

	// NewUpdateTenantPayloadResolver creates new instance of the UpdateTenantPayloadResolverContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// Returns the new instance or error if something goes wrong
	NewUpdateTenantPayloadResolver(
		ctx context.Context,
		clientMutationId *string) (UpdateTenantPayloadResolverContract, error)

	// NewDeleteTenant creates new instance of the DeleteTenantContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// Returns the new instance or error if something goes wrong
	NewDeleteTenant(ctx context.Context) (DeleteTenantContract, error)

	// NewDeleteTenantPayloadResolver creates new instance of the DeleteTenantPayloadResolverContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// Returns the new instance or error if something goes wrong
	NewDeleteTenantPayloadResolver(
		ctx context.Context,
		clientMutationId *string) (DeleteTenantPayloadResolverContract, error)
}

// RootResolverContract declares the root resolver
type RootResolverContract interface {
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

	// DeleteTenant returns delete tenant mutator
	// ctx: Mandatory. Reference to the context
	// Returns the delete tenant mutator or error if something goes wrong
	DeleteTenant(
		ctx context.Context,
		args DeleteTenantInputArgument) (DeleteTenantPayloadResolverContract, error)
}

// CreateTenantPayloadResolverContract declares the resolver that can return the payload contains the result of creating a new tenant
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

// UpdateTenantPayloadResolverContract declares the resolver that can return the payload contains the result of updating an existing tenant
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

// DeleteTenantPayloadResolverContract declares the resolver that can return the payload contains the result of deleting an existing tenant
type DeleteTenantPayloadResolverContract interface {
	// ClientMutationId returns the client mutation ID that was provided as part of the mutation request
	// ctx: Mandatory. Reference to the context
	// Returns the provided clientMutationId as part of mutation request
	ClientMutationId(ctx context.Context) *string
}

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

type CreateTenantInput struct {
	Name             string
	ClientMutationId *string
}

type CreateTenantInputArgument struct {
	Input CreateTenantInput
}

type UpdateTenantInput struct {
	ID               graphql.ID
	Name             string
	ClientMutationId *string
}

type UpdateTenantInputArgument struct {
	Input UpdateTenantInput
}

type DeleteTenantInput struct {
	ID               graphql.ID
	ClientMutationId *string
}

type DeleteTenantInputArgument struct {
	Input DeleteTenantInput
}
