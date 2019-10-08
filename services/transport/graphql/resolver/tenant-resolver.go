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

type tenantClusterEdgeClusterInputArgument struct {
	EdgeClusterID graphql.ID
}

type tenantEdgeClustersInputArgument struct {
	connectionArgument
	EdgeClusterIDs *[]graphql.ID
	SortOption     *string
}

// TenantResolverContract declares the resolver that can retrieve tenant information
type TenantResolverContract interface {
	// ID returns tenant unique identifier
	// ctx: Mandatory. Reference to the context
	// Returns the tenant unique identifier
	ID(ctx context.Context) graphql.ID

	// Name returns tenant name
	// ctx: Mandatory. Reference to the context
	// Returns the tenant name
	Name(ctx context.Context) string

	// EdgeCluster returns tenant resolver
	// ctx: Mandatory. Reference to the context
	// args: Mandatory. The argument list
	// Returns the tenant resolver or error if something goes wrong
	EdgeCluster(
		ctx context.Context,
		args tenantClusterEdgeClusterInputArgument) (EdgeClusterResolverContract, error)

	// EdgeClusters returns tenant conenction compatible with graphql-relay
	// ctx: Mandatory. Reference to the context
	// args: Mandatory. The argument list
	// Returns the tenant resolver or error if something goes wrong
	EdgeClusters(
		ctx context.Context,
		args tenantEdgeClustersInputArgument) (EdgeClusterTypeConnectionResolverContract, error)
}

type tenantResolver struct {
	logger          *zap.Logger
	resolverCreator ResolverCreatorContract
	id              graphql.ID
	name            string
}

// NewTenantResolver creates new instance of the tenantResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can create new instances of resolvers
// logger: Mandatory. Reference to the logger service
// tenantID: Mandatory. the tenant unique identifier
// Returns the new instance or error if something goes wrong
func NewTenantResolver(
	ctx context.Context,
	resolverCreator ResolverCreatorContract,
	logger *zap.Logger,
	tenantID graphql.ID) (TenantResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if strings.Trim(string(tenantID), " ") == "" {
		return nil, commonErrors.NewArgumentError("tenantID", "tenantID is required")
	}

	return &tenantResolver{
		logger:          logger,
		resolverCreator: resolverCreator,
		id:              tenantID,
		name:            cuid.New(),
	}, nil
}

// ID returns tenant unique identifier
// ctx: Mandatory. Reference to the context
// Returns the tenant unique identifier
func (r *tenantResolver) ID(ctx context.Context) graphql.ID {
	return r.id
}

// Name returns tenant name
// ctx: Mandatory. Reference to the context
// Returns the tenant name or error
func (r *tenantResolver) Name(ctx context.Context) string {
	return r.name
}

// EdgeCluster returns tenant resolver
// ctx: Mandatory. Reference to the context
// args: Mandatory. The argument list
// Returns the tenant resolver or error if something goes wrong
func (r *tenantResolver) EdgeCluster(
	ctx context.Context,
	args tenantClusterEdgeClusterInputArgument) (EdgeClusterResolverContract, error) {
	return r.resolverCreator.NewEdgeClusterResolver(
		ctx,
		args.EdgeClusterID)
}

// EdgeClusters returns tenant conenction compatible with graphql-relay
// ctx: Mandatory. Reference to the context
// args: Mandatory. The argument list
// Returns the tenant resolver or error if something goes wrong
func (r *tenantResolver) EdgeClusters(
	ctx context.Context,
	args tenantEdgeClustersInputArgument) (EdgeClusterTypeConnectionResolverContract, error) {
	return r.resolverCreator.NewEdgeClusterTypeConnectionResolver(ctx)
}
