// Package graphql implements functions to expose api-gateway service endpoint using GraphQL protocol.
package graphql

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/transport/graphql/mutation"
	"github.com/decentralized-cloud/api-gateway/services/transport/graphql/query"
	"github.com/decentralized-cloud/api-gateway/services/transport/graphql/root"
	"github.com/decentralized-cloud/api-gateway/services/transport/graphql/types"
	"github.com/graph-gophers/graphql-go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type resolverCreator struct {
	logger *zap.Logger
}

// NewResolverCreator creates new instance of the resolverCreator, setting up all dependencies and returns the instance
// logger: Mandatory. Reference to the logger service
// Returns the new instance or error if something goes wrong
func NewResolverCreator(logger *zap.Logger) (types.ResolverCreatorContract, error) {
	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	return &resolverCreator{
		logger: logger,
	}, nil
}

// NewPageInfoResolver creates new PageInfoResolverContract and returns it
// ctx: Mandatory. Reference to the context
// startCursor: Mandatory. Reference to the start cursor
// endCursor: Mandatory. Reference to the end cursor
// hasNextPage: Mandatory. Reference to the value indicates whether returned page has next page to be retrieved
// hasPreviousPage: Mandatory. Reference to the value indicates whether returned page has previous page to be retrieved
// Returns the PageInfoResolverContract or error if something goes wrong
func (creator *resolverCreator) NewPageInfoResolver(
	ctx context.Context,
	startCursor *string,
	endCursor *string,
	hasNextPage bool,
	hasPreviousPage bool) (types.PageInfoResolverContract, error) {
	return query.NewPageInfoResolver(
		startCursor,
		endCursor,
		hasNextPage,
		hasPreviousPage)
}

// NewRootResolver creates new RootResolverContract and returns it
// ctx: Mandatory. Reference to the context
// Returns the RootResolverContract or error if something goes wrong
func (creator *resolverCreator) NewRootResolver(ctx context.Context) (types.RootResolverContract, error) {
	return root.NewRootResolver(
		ctx,
		creator,
		creator.logger)
}

// NewUserResolver creates new UserResolverContract and returns it
// ctx: Mandatory. Reference to the context
// userID: Mandatory. The user unique identifier
// Returns the UserResolverContract or error if something goes wrong
func (creator *resolverCreator) NewUserResolver(
	ctx context.Context,
	userID graphql.ID) (types.UserResolverContract, error) {
	return query.NewUserResolver(
		ctx,
		creator,
		creator.logger,
		userID)
}

// NewTenantResolver creates new TenantResolverContract and returns it
// ctx: Mandatory. Reference to the context
// tenantID: Mandatory. The tenant unique identifier
// Returns the TenantResolverContract or error if something goes wrong
func (creator *resolverCreator) NewTenantResolver(
	ctx context.Context,
	tenantID graphql.ID) (types.TenantResolverContract, error) {
	return query.NewTenantResolver(
		ctx,
		creator,
		creator.logger,
		tenantID)
}

// NewTenantTypeEdgeResolver creates new TenantTypeEdgeResolverContract and returns it
// ctx: Mandatory. Reference to the context
// tenantID: Mandatory. The tenant unique identifier
// cursor: Mandatory. The cursor
// Returns the TenantTypeEdgeResolverContract or error if something goes wrong
func (creator *resolverCreator) NewTenantTypeEdgeResolver(
	ctx context.Context,
	tenantID graphql.ID,
	cursor string) (types.TenantTypeEdgeResolverContract, error) {
	return query.NewTenantTypeEdgeResolver(
		ctx,
		creator,
		tenantID,
		cursor)
}

// NewTenantTypeConnectionResolver creates new TenantTypeConnectionResolverContract and returns it
// ctx: Mandatory. Reference to the context
// Returns the TenantTypeConnectionResolverContract or error if something goes wrong
func (creator *resolverCreator) NewTenantTypeConnectionResolver(ctx context.Context) (types.TenantTypeConnectionResolverContract, error) {
	return query.NewTenantTypeConnectionResolver(ctx, creator)
}

// NewEdgeClusterResolver creates new EdgeClusterResolverContract and returns it
// ctx: Mandatory. Reference to the context
// tenantID: Mandatory. The tenant unique identifier
// Returns the EdgeClusterResolverContract or error if something goes wrong
func (creator *resolverCreator) NewEdgeClusterResolver(
	ctx context.Context,
	tenantID graphql.ID) (types.EdgeClusterResolverContract, error) {
	return query.NewEdgeClusterResolver(
		ctx,
		creator,
		creator.logger,
		tenantID)
}

// NewEdgeClusterTypeEdgeResolver creates new EdgeClusterTypeEdgeResolverContract and returns it
// ctx: Mandatory. Reference to the context
// tenantID: Mandatory. The tenant unique identifier
// cursor: Mandatory. The cursor
// Returns the EdgeClusterTypeEdgeResolverContract or error if something goes wrong
func (creator *resolverCreator) NewEdgeClusterTypeEdgeResolver(
	ctx context.Context,
	tenantID graphql.ID,
	cursor string) (types.EdgeClusterTypeEdgeResolverContract, error) {
	return query.NewEdgeClusterTypeEdgeResolver(
		ctx,
		creator,
		tenantID,
		cursor)
}

// NewEdgeClusterTypeConnectionResolver creates new EdgeClusterTypeConnectionResolverContract and returns it
// ctx: Mandatory. Reference to the context
// Returns the EdgeClusterTypeConnectionResolverContract or error if something goes wrong
func (creator *resolverCreator) NewEdgeClusterTypeConnectionResolver(ctx context.Context) (types.EdgeClusterTypeConnectionResolverContract, error) {
	return query.NewEdgeClusterTypeConnectionResolver(ctx, creator)
}

// NewCreateTenant creates new instance of the createTenant, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewCreateTenant(ctx context.Context) (types.CreateTenantContract, error) {
	return mutation.NewCreateTenant(
		ctx,
		creator,
		creator.logger)
}

// NewCreateTenantPayloadResolver creates new instance of the createTenantPayloadResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewCreateTenantPayloadResolver(
	ctx context.Context,
	clientMutationId *string) (types.CreateTenantPayloadResolverContract, error) {
	return mutation.NewCreateTenantPayloadResolver(
		ctx,
		creator,
		clientMutationId)
}

// NewUpdateTenant creates new instance of the updateTenant, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewUpdateTenant(ctx context.Context) (types.UpdateTenantContract, error) {
	return mutation.NewUpdateTenant(
		ctx,
		creator,
		creator.logger)
}

// NewUpdateTenantPayloadResolver creates new instance of the updateTenantPayloadResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewUpdateTenantPayloadResolver(
	ctx context.Context,
	clientMutationId *string) (types.UpdateTenantPayloadResolverContract, error) {
	return mutation.NewUpdateTenantPayloadResolver(
		ctx,
		creator,
		clientMutationId)
}
