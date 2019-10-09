// package query implements different GraphQL query resovlers required by the GraphQL transport layer
package query

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/transport/graphql/types"
)

type pageInfoResolver struct {
	startCursor     *string
	endCursor       *string
	hasNextPage     bool
	hasPreviousPage bool
}

// NewPageInfoResolver creates new instance of the pageInfoResolver, setting up all dependencies and returns the instance
// startCursor: Mandatory. Reference to the start cursor
// endCursor: Mandatory. Reference to the end cursor
// hasNextPage: Mandatory. Reference to the value indicates whether returned page has next page to be retrieved
// hasPreviousPage: Mandatory. Reference to the value indicates whether returned page has previous page to be retrieved
// Returns the new instance or error if something goes wrong
func NewPageInfoResolver(
	startCursor *string,
	endCursor *string,
	hasNextPage bool,
	hasPreviousPage bool) (types.PageInfoResolverContract, error) {
	return &pageInfoResolver{
		startCursor:     startCursor,
		endCursor:       endCursor,
		hasNextPage:     hasNextPage,
		hasPreviousPage: hasPreviousPage,
	}, nil
}

// StartCursor returns start cursor
// ctx: Mandatory. Reference to the context
// Returns the start cursor
func (r *pageInfoResolver) StartCursor(ctx context.Context) *string {
	return r.startCursor
}

// EndCursor returns end cursor
// ctx: Mandatory. Reference to the context
// Returns the end cursor
func (r *pageInfoResolver) EndCursor(ctx context.Context) *string {
	return r.endCursor
}

// HasNextPage indicates whether returned page has next page to be retrieved
// ctx: Mandatory. Reference to the context
// Returns the value indicates whether returned page has next page to be retrieved
func (r *pageInfoResolver) HasNextPage(ctx context.Context) bool {
	return r.hasNextPage
}

// HasPreviousPage indicates whether returned page has previous page to be retrieved
// ctx: Mandatory. Reference to the context
// Returns the value indicates whether returned page has previous page to be retrieved
func (r *pageInfoResolver) HasPreviousPage(ctx context.Context) bool {
	return r.hasPreviousPage
}
