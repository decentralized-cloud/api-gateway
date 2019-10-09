// package query implements different GraphQL query resovlers required by the GraphQL transport layer
package query

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/transport/graphql/types"
	commonErrors "github.com/micro-business/go-core/system/errors"
)

type tenantTypeConnectionResolver struct {
	resolverCreator types.ResolverCreatorContract
}

// NewTenantTypeConnectionResolver creates new instance of the tenantTypeConnectionResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can create new instances of resolvers
// tenantID: Mandatory. the tenant unique identifier
// cursor: Mandatory. the cursor
// Returns the new instance or error if something goes wrong
func NewTenantTypeConnectionResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract) (types.TenantTypeConnectionResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	return &tenantTypeConnectionResolver{
		resolverCreator: resolverCreator,
	}, nil
}

// PageInfo returns the paging information compatible with graphql-relay
// ctx: Mandatory. Reference to the context
// Returns the paging information
func (r *tenantTypeConnectionResolver) PageInfo(ctx context.Context) (types.PageInfoResolverContract, error) {
	startCursor := "start cursor"
	endCurstor := "End cursor"

	return r.resolverCreator.NewPageInfoResolver(
		ctx,
		&startCursor,
		&endCurstor,
		true,
		false)
}

// Edges returns the tenant edges compatible with graphql-relay
// ctx: Mandatory. Reference to the context
// Returns the tenant edges
func (r *tenantTypeConnectionResolver) Edges(ctx context.Context) (*[]types.TenantTypeEdgeResolverContract, error) {
	return &[]types.TenantTypeEdgeResolverContract{}, nil
}
