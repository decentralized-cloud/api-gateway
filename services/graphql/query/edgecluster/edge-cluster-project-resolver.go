// Package edgecluster implements different edge cluster GraphQL query resovlers required by the GraphQL transport layer
package edgecluster

import (
	"context"
	"strings"

	"github.com/decentralized-cloud/api-gateway/services/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/edgecluster"
	"github.com/graph-gophers/graphql-go"
	"github.com/lucsky/cuid"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type edgeClusterProjectResolver struct {
	logger          *zap.Logger
	resolverCreator types.ResolverCreatorContract
	projectID       string
	name            string
}

// NewEdgeClusterProjectResolver creates new instance of the edgeClusterProjectResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can create new instances of resolvers
// projectID: Mandatory. the project unique identifier
// Returns the new instance or error if something goes wrong
func NewEdgeClusterProjectResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	logger *zap.Logger,
	projectID string) (edgecluster.EdgeClusterProjectResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if strings.Trim(projectID, " ") == "" {
		return nil, commonErrors.NewArgumentError("projectID", "projectID is required")
	}

	return &edgeClusterProjectResolver{
		logger:          logger,
		resolverCreator: resolverCreator,
		projectID:       projectID,
		name:            cuid.New(),
	}, nil
}

// ID returns project unique identifier
// ctx: Mandatory. Reference to the context
// Returns the project unique identifier
func (r *edgeClusterProjectResolver) ID(ctx context.Context) graphql.ID {
	return graphql.ID(r.projectID)
}

// Name returns project name
// ctx: Mandatory. Reference to the context
// Returns the project name or error
func (r *edgeClusterProjectResolver) Name(ctx context.Context) string {
	return r.name
}
