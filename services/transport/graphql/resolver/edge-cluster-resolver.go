// package resolver implements different GraphQL resolvers required by the GraphQL transport layer
package resolver

import (
	"context"
	"strings"

	"github.com/graph-gophers/graphql-go"
	"github.com/lucsky/cuid"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

// EdgeClusterResolverContract declares the resolver that can retrieve edge-cluster information
type EdgeClusterResolverContract interface {
	// ID returns edge-cluster unique identifier
	// ctx: Mandatory. Reference to the context
	// Returns the edge-cluster unique identifier
	ID(ctx context.Context) graphql.ID

	// Name returns edge-cluster name
	// ctx: Mandatory. Reference to the context
	// Returns the edge-cluster name
	Name(ctx context.Context) string
}

type edgeClusterResolver struct {
	logger          *zap.Logger
	resolverCreator ResolverCreatorContract
	id              graphql.ID
	name            string
}

// NewEdgeClusterResolver creates new instance of the edgeClusterResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can create new instances of resolvers
// logger: Mandatory. Reference to the logger service
// edgeClusterID: Mandatory. the edge-cluster unique identifier
// Returns the new instance or error if something goes wrong
func NewEdgeClusterResolver(
	ctx context.Context,
	resolverCreator ResolverCreatorContract,
	logger *zap.Logger,
	edgeClusterID graphql.ID) (EdgeClusterResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if strings.Trim(string(edgeClusterID), " ") == "" {
		return nil, commonErrors.NewArgumentError("edgeClusterID", "edgeClusterID is required")
	}

	return &edgeClusterResolver{
		logger:          logger,
		resolverCreator: resolverCreator,
		id:              edgeClusterID,
		name:            cuid.New(),
	}, nil
}

// ID returns edge-cluster unique identifier
// ctx: Mandatory. Reference to the context
// Returns the edge-cluster unique identifier
func (r *edgeClusterResolver) ID(ctx context.Context) graphql.ID {
	return r.id
}

// Name returns edge-cluster name
// ctx: Mandatory. Reference to the context
// Returns the edge-cluster name or error
func (r *edgeClusterResolver) Name(ctx context.Context) string {
	return r.name
}