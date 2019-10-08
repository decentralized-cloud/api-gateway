// package resolver implements different GraphQL resolvers required by the GraphQL transport layer
package resolver

import (
	"context"

	commonErrors "github.com/micro-business/go-core/system/errors"
)

// EdgeClusterTypeConnectionResolverContract declares the resolver that returns edge-cluster edge compatible with graphql-relay
type EdgeClusterTypeConnectionResolverContract interface {
	// PageInfo returns the paging information compatible with graphql-relay
	// ctx: Mandatory. Reference to the context
	// Returns the paging information
	PageInfo(ctx context.Context) (PageInfoResolverContract, error)

	// Edges returns the edge-cluster edges compatible with graphql-relay
	// ctx: Mandatory. Reference to the context
	// Returns the edge-cluster edges
	Edges(ctx context.Context) (*[]EdgeClusterTypeEdgeResolverContract, error)
}

type edgeClusterTypeConnectionResolver struct {
	resolverCreator ResolverCreatorContract
}

// NewEdgeClusterTypeConnectionResolver creates new instance of the edgeClusterTypeConnectionResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can create new instances of resolvers
// edgeClusterID: Mandatory. the edge-cluster unique identifier
// cursor: Mandatory. the cursor
// Returns the new instance or error if something goes wrong
func NewEdgeClusterTypeConnectionResolver(
	ctx context.Context,
	resolverCreator ResolverCreatorContract) (EdgeClusterTypeConnectionResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	return &edgeClusterTypeConnectionResolver{
		resolverCreator: resolverCreator,
	}, nil
}

// PageInfo returns the paging information compatible with graphql-relay
// ctx: Mandatory. Reference to the context
// Returns the paging information
func (r *edgeClusterTypeConnectionResolver) PageInfo(ctx context.Context) (PageInfoResolverContract, error) {
	startCursor := "start cursor"
	endCurstor := "End cursor"

	return r.resolverCreator.NewPageInfoResolver(
		ctx,
		&startCursor,
		&endCurstor,
		true,
		false)
}

// Edges returns the edge-cluster edges compatible with graphql-relay
// ctx: Mandatory. Reference to the context
// Returns the edge-cluster edges
func (r *edgeClusterTypeConnectionResolver) Edges(ctx context.Context) (*[]EdgeClusterTypeEdgeResolverContract, error) {
	return &[]EdgeClusterTypeEdgeResolverContract{}, nil
}
