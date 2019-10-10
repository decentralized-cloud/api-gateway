// packae relay implements common relay types used in the GraphQL transport layer
package relay

import (
	"context"
)

type PageInfoResolverCreatorContract interface {
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
}

// PageInfoResolverContract declares the resolver that returns paging information compatible with graphql-relay specification
type PageInfoResolverContract interface {
	// StartCursor returns start cursor
	// ctx: Mandatory. Reference to the context
	// Returns the start cursor
	StartCursor(ctx context.Context) *string

	// EndCursor returns end cursor
	// ctx: Mandatory. Reference to the context
	// Returns the end cursor
	EndCursor(ctx context.Context) *string

	// HasNextPage indicates whether returned page has next page to be retrieved
	// ctx: Mandatory. Reference to the context
	// Returns the value indicates whether returned page has next page to be retrieved
	HasNextPage(ctx context.Context) bool

	// HasPreviousPage indicates whether returned page has previous page to be retrieved
	// ctx: Mandatory. Reference to the context
	// Returns the value indicates whether returned page has previous page to be retrieved
	HasPreviousPage(ctx context.Context) bool
}
