// package query implements different GraphQL query resovlers required by the GraphQL transport layer
package query

import (
	"context"
	"strings"

	"github.com/decentralized-cloud/api-gateway/services/transport/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/transport/graphql/types/edgecluster"
	"github.com/decentralized-cloud/api-gateway/services/transport/graphql/types/tenant"
	"github.com/graph-gophers/graphql-go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type userResolver struct {
	logger          *zap.Logger
	resolverCreator types.ResolverCreatorContract
	userID          graphql.ID
}

// NewUserResolver creates new instance of the userResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can create new instances of resolvers
// logger: Mandatory. Reference to the logger service
// userID: Mandatory. the tenant unique identifier
// Returns the new instance or error if something goes wrong
func NewUserResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	logger *zap.Logger,
	userID graphql.ID) (types.UserResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if strings.Trim(string(userID), " ") == "" {
		return nil, commonErrors.NewArgumentError("userID", "userID is required")
	}

	return &userResolver{
		logger:          logger,
		resolverCreator: resolverCreator,
		userID:          userID,
	}, nil
}

// ID returns user unique identifier
// ctx: Mandatory. Reference to the context
// Returns the user unique identifier
func (r *userResolver) ID(ctx context.Context) graphql.ID {
	return r.userID
}

// Tenant returns tenant resolver
// ctx: Mandatory. Reference to the context
// args: Mandatory. The argument list
// Returns the tenant resolver or error if something goes wrong
func (r *userResolver) Tenant(
	ctx context.Context,
	args types.UserTenantInputArgument) (tenant.TenantResolverContract, error) {
	return r.resolverCreator.NewTenantResolver(
		ctx,
		args.TenantID)
}

// Tenants returns tenant conenction compatible with graphql-relay
// ctx: Mandatory. Reference to the context
// args: Mandatory. The argument list
// Returns the tenant resolver or error if something goes wrong
func (r *userResolver) Tenants(
	ctx context.Context,
	args types.UserTenantsInputArgument) (tenant.TenantTypeConnectionResolverContract, error) {
	return r.resolverCreator.NewTenantTypeConnectionResolver(ctx)
}

// EdgeCluster returns tenant resolver
// ctx: Mandatory. Reference to the context
// args: Mandatory. The argument list
// Returns the tenant resolver or error if something goes wrong
func (r *userResolver) EdgeCluster(
	ctx context.Context,
	args types.UserEdgeClusterInputArgument) (edgecluster.EdgeClusterResolverContract, error) {
	return r.resolverCreator.NewEdgeClusterResolver(
		ctx,
		args.EdgeClusterID)
}

// EdgeClusters returns tenant conenction compatible with graphql-relay
// ctx: Mandatory. Reference to the context
// args: Mandatory. The argument list
// Returns the tenant resolver or error if something goes wrong
func (r *userResolver) EdgeClusters(
	ctx context.Context,
	args types.UserEdgeClustersInputArgument) (edgecluster.EdgeClusterTypeConnectionResolverContract, error) {
	return r.resolverCreator.NewEdgeClusterTypeConnectionResolver(ctx)
}
