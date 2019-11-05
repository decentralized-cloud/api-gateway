// Package edgelcuster implements different edge cluster GraphQL query resovlers required by the GraphQL transport layer
package edgeclster

import (
	"context"
	"strings"

	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types/edgecluster"
	"github.com/graph-gophers/graphql-go"
	"github.com/lucsky/cuid"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type edgeClusterTenantResolver struct {
	logger          *zap.Logger
	resolverCreator types.ResolverCreatorContract
	tenantID        string
	name            string
}

// NewEdgeClusterTenantResolver creates new instance of the edgeClusterTenantResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can create new instances of resolvers
// tenantID: Mandatory. the tenant unique identifier
// Returns the new instance or error if something goes wrong
func NewEdgeClusterTenantResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	logger *zap.Logger,
	tenantID string) (edgecluster.EdgeClusterTenantResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if strings.Trim(tenantID, " ") == "" {
		return nil, commonErrors.NewArgumentError("tenantID", "tenantID is required")
	}

	return &edgeClusterTenantResolver{
		logger:          logger,
		resolverCreator: resolverCreator,
		tenantID:        tenantID,
		name:            cuid.New(),
	}, nil
}

// ID returns tenant unique identifier
// ctx: Mandatory. Reference to the context
// Returns the tenant unique identifier
func (r *edgeClusterTenantResolver) ID(ctx context.Context) graphql.ID {
	return graphql.ID(r.tenantID)
}

// Name returns tenant name
// ctx: Mandatory. Reference to the context
// Returns the tenant name or error
func (r *edgeClusterTenantResolver) Name(ctx context.Context) string {
	return r.name
}
