// package edgelcuster implements different edge cluster GraphQL query resovlers required by the GraphQL transport layer
package edgeclster

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/transport/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/transport/graphql/types/edgecluster"
	"github.com/decentralized-cloud/api-gateway/services/transport/graphql/types/relay"
	commonErrors "github.com/micro-business/go-core/system/errors"
)

type edgeClusterTypeConnectionResolver struct {
	resolverCreator types.ResolverCreatorContract
}

// NewEdgeClusterTypeConnectionResolver creates new instance of the edgeClusterTypeConnectionResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can create new instances of resolvers
// edgeClusterID: Mandatory. the edge-cluster unique identifier
// cursor: Mandatory. the cursor
// Returns the new instance or error if something goes wrong
func NewEdgeClusterTypeConnectionResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract) (edgecluster.EdgeClusterTypeConnectionResolverContract, error) {
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
func (r *edgeClusterTypeConnectionResolver) PageInfo(ctx context.Context) (relay.PageInfoResolverContract, error) {
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
func (r *edgeClusterTypeConnectionResolver) Edges(ctx context.Context) (*[]edgecluster.EdgeClusterTypeEdgeResolverContract, error) {
	return &[]edgecluster.EdgeClusterTypeEdgeResolverContract{}, nil
}
