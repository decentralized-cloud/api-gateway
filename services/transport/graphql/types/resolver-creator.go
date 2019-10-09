// package types defines the different interfaces used in the GraphQL implementation
package types

import (
	"context"

	"github.com/graph-gophers/graphql-go"
)

// ResolverCreatorContract declares the service that can create different resolvers
type ResolverCreatorContract interface {
	// NewRootResolver creates new RootResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// Returns the RootResolverContract or error if something goes wrong
	NewRootResolver(ctx context.Context) (RootResolverContract, error)

	// NewPageInfoResolver creates new PageInfoResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// startCursor: Mandatory. Reference to the start cursor
	// endCursor: Mandatory. Reference to the end cursor
	// hasNextPage: Mandatory. Reference to the value indicates whether returned page has next page to be retrieved
	// hasPreviousPage: Mandatory. Reference to the value indicates whether returned page has previous page to be retrieved
	// Returns the PageInfoResolverContract or error if something goes wrong
	NewPageInfoResolver(
		ctx context.Context,
		startCursor *string,
		endCursor *string,
		hasNextPage bool,
		hasPreviousPage bool) (PageInfoResolverContract, error)

	// NewUserResolver creates new UserResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// userID: Mandatory. The user unique identifier
	// Returns the UserResolverContract or error if something goes wrong
	NewUserResolver(
		ctx context.Context,
		userID graphql.ID) (UserResolverContract, error)

	// NewTenantResolver creates new TenantResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// tenantID: Mandatory. The tenant unique identifier
	// Returns the TenantResolverContract or error if something goes wrong
	NewTenantResolver(
		ctx context.Context,
		tenantID graphql.ID) (TenantResolverContract, error)

	// NewTenantTypeEdgeResolver creates new TenantTypeEdgeResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// tenantID: Mandatory. The tenant unique identifier
	// cursor: Mandatory. The cursor
	// Returns the TenantTypeEdgeResolverContract or error if something goes wrong
	NewTenantTypeEdgeResolver(
		ctx context.Context,
		tenantID graphql.ID,
		cursor string) (TenantTypeEdgeResolverContract, error)

	// NewTenantTypeConnectionResolver creates new TenantTypeConnectionResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// Returns the TenantTypeConnectionResolverContract or error if something goes wrong
	NewTenantTypeConnectionResolver(ctx context.Context) (TenantTypeConnectionResolverContract, error)

	// NewEdgeClusterResolver creates new EdgeClusterResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// tenantID: Mandatory. The tenant unique identifier
	// Returns the EdgeClusterResolverContract or error if something goes wrong
	NewEdgeClusterResolver(
		ctx context.Context,
		tenantID graphql.ID) (EdgeClusterResolverContract, error)

	// NewEdgeClusterTypeEdgeResolver creates new EdgeClusterTypeEdgeResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// tenantID: Mandatory. The tenant unique identifier
	// cursor: Mandatory. The cursor
	// Returns the EdgeClusterTypeEdgeResolverContract or error if something goes wrong
	NewEdgeClusterTypeEdgeResolver(
		ctx context.Context,
		tenantID graphql.ID,
		cursor string) (EdgeClusterTypeEdgeResolverContract, error)

	// NewEdgeClusterTypeConnectionResolver creates new EdgeClusterTypeConnectionResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// Returns the EdgeClusterTypeConnectionResolverContract or error if something goes wrong
	NewEdgeClusterTypeConnectionResolver(ctx context.Context) (EdgeClusterTypeConnectionResolverContract, error)

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
}
