// Package edgecluster implements different edge cluster GraphQL query resovlers required by the GraphQL transport layer
package edgecluster

import (
	"context"
	"strings"

	"github.com/decentralized-cloud/api-gateway/services/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/edgecluster"
	commonErrors "github.com/micro-business/go-core/system/errors"
)

type edgeClusterTypeEdgeResolver struct {
	resolverCreator   types.ResolverCreatorContract
	edgeClusterID     string
	edgeClusterDetail *edgecluster.EdgeClusterDetail
	cursor            string
}

// NewEdgeClusterTypeEdgeResolver creates new instance of the edgeClusterTypeEdgeResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can create new instances of resolvers
// edgeClusterID: Mandatory. the edge cluster unique identifier
// edgeClusterDetail: Optional. The edge cluster details, if provided, the value be used instead of contacting  the edge cluster service
// cursor: Mandatory. the cursor
// Returns the new instance or error if something goes wrong
func NewEdgeClusterTypeEdgeResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	edgeClusterID string,
	edgeClusterDetail *edgecluster.EdgeClusterDetail,
	cursor string) (edgecluster.EdgeClusterTypeEdgeResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if strings.Trim(edgeClusterID, " ") == "" {
		return nil, commonErrors.NewArgumentError("edgeClusterID", "edgeClusterID is required")
	}

	if strings.Trim(cursor, " ") == "" {
		return nil, commonErrors.NewArgumentError("cursor", "cursor is required")
	}

	return &edgeClusterTypeEdgeResolver{
		resolverCreator:   resolverCreator,
		edgeClusterID:     edgeClusterID,
		edgeClusterDetail: edgeClusterDetail,
		cursor:            cursor,
	}, nil
}

// Node returns the edge cluster resolver
// ctx: Mandatory. Reference to the context
// Returns the edge cluster resolver or error if something goes wrong
func (r *edgeClusterTypeEdgeResolver) Node(ctx context.Context) (edgecluster.EdgeClusterResolverContract, error) {
	return r.resolverCreator.NewEdgeClusterResolver(
		ctx,
		r.edgeClusterID,
		r.edgeClusterDetail)
}

// Cursor returns the cursor for the edge cluster edge compatible with graphql-relay
// ctx: Mandatory. Reference to the context
// Returns the cursor
func (r *edgeClusterTypeEdgeResolver) Cursor(ctx context.Context) string {
	return r.cursor
}
