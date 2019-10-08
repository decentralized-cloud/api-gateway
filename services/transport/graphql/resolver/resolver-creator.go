// package resolver implements different GraphQL resolvers required by the GraphQL transport layer
package resolver

import (
	"context"

	"github.com/graph-gophers/graphql-go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

// ResolverCreatorContract declares the service that can create different resolvers
type ResolverCreatorContract interface {
	// NewQueryResolver creates new QueryResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// Returns the QueryResolverContract or error if something goes wrong
	NewQueryResolver(ctx context.Context) (QueryResolverContract, error)

	// NewUserResolver creates new UserResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// userID: Mandatory. The user unique identifier
	// Returns the UserResolverContract or error if something goes wrong
	NewUserResolver(
		ctx context.Context,
		userID graphql.ID) (UserResolverContract, error)

	// NewTenantResolver creates new TenantResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// tenantID: Mandatory. The tenant unique identifier
	// Returns the TenantResolverContract or error if something goes wrong
	NewTenantResolver(
		ctx context.Context,
		tenantID graphql.ID) (TenantResolverContract, error)

	// NewTenantTypeEdgeResolver creates new TenantTypeEdgeResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// tenantID: Mandatory. The tenant unique identifier
	// cursor: Mandatory. The cursor
	// Returns the TenantTypeEdgeResolverContract or error if something goes wrong
	NewTenantTypeEdgeResolver(
		ctx context.Context,
		tenantID graphql.ID,
		cursor string) (TenantTypeEdgeResolverContract, error)

	// NewTenantTypeConnectionResolver creates new TenantTypeConnectionResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// Returns the TenantTypeConnectionResolverContract or error if something goes wrong
	NewTenantTypeConnectionResolver(ctx context.Context) (TenantTypeConnectionResolverContract, error)

	// NewEdgeClusterResolver creates new EdgeClusterResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// tenantID: Mandatory. The tenant unique identifier
	// Returns the EdgeClusterResolverContract or error if something goes wrong
	NewEdgeClusterResolver(
		ctx context.Context,
		tenantID graphql.ID) (EdgeClusterResolverContract, error)

	// NewEdgeClusterTypeEdgeResolver creates new EdgeClusterTypeEdgeResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// tenantID: Mandatory. The tenant unique identifier
	// cursor: Mandatory. The cursor
	// Returns the EdgeClusterTypeEdgeResolverContract or error if something goes wrong
	NewEdgeClusterTypeEdgeResolver(
		ctx context.Context,
		tenantID graphql.ID,
		cursor string) (EdgeClusterTypeEdgeResolverContract, error)

	// NewEdgeClusterTypeConnectionResolver creates new EdgeClusterTypeConnectionResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// Returns the EdgeClusterTypeConnectionResolverContract or error if something goes wrong
	NewEdgeClusterTypeConnectionResolver(ctx context.Context) (EdgeClusterTypeConnectionResolverContract, error)

	// NewPageInfoResolver creates new PageInfoResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// startCursor: Mandatory. Reference to the start cursor
	// endCursor: Mandatory. Reference to the end cursor
	// hasNextPage: Mandatory. Reference to the value indicates whether returned page has next page to be retrieved
	// hasPreviousPage: Mandatory. Reference to the value indicates whether returned page has previous page to be retrieved
	// Returns the PageInfoResolverContract or error if something goes wrong
	NewPageInfoResolver(
		ctx context.Context,
		startCursor *string,
		endCursor *string,
		hasNextPage bool,
		hasPreviousPage bool) (PageInfoResolverContract, error)
}

type resolverCreator struct {
	logger *zap.Logger
}

// NewResolverCreator creates new instance of the resolverCreator, setting up all dependencies and returns the instance
// logger: Mandatory. Reference to the logger service
// Returns the new instance or error if something goes wrong
func NewResolverCreator(logger *zap.Logger) (ResolverCreatorContract, error) {
	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	return &resolverCreator{
		logger: logger,
	}, nil
}

// NewQueryResolver creates new QueryResolverContract and returns it
// ctx: Mandatory. Reference to the context
// Returns the QueryResolverContract or error if something goes wrong
func (creator *resolverCreator) NewQueryResolver(ctx context.Context) (QueryResolverContract, error) {
	return NewQueryResolver(
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
	userID graphql.ID) (UserResolverContract, error) {
	return NewUserResolver(
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
	tenantID graphql.ID) (TenantResolverContract, error) {
	return NewTenantResolver(
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
	cursor string) (TenantTypeEdgeResolverContract, error) {
	return NewTenantTypeEdgeResolver(
		ctx,
		creator,
		tenantID,
		cursor)
}

// NewTenantTypeConnectionResolver creates new TenantTypeConnectionResolverContract and returns it
// ctx: Mandatory. Reference to the context
// Returns the TenantTypeConnectionResolverContract or error if something goes wrong
func (creator *resolverCreator) NewTenantTypeConnectionResolver(
	ctx context.Context) (TenantTypeConnectionResolverContract, error) {
	return NewTenantTypeConnectionResolver(ctx, creator)
}

// NewEdgeClusterResolver creates new EdgeClusterResolverContract and returns it
// ctx: Mandatory. Reference to the context
// tenantID: Mandatory. The tenant unique identifier
// Returns the EdgeClusterResolverContract or error if something goes wrong
func (creator *resolverCreator) NewEdgeClusterResolver(
	ctx context.Context,
	tenantID graphql.ID) (EdgeClusterResolverContract, error) {
	return NewEdgeClusterResolver(
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
	cursor string) (EdgeClusterTypeEdgeResolverContract, error) {
	return NewEdgeClusterTypeEdgeResolver(
		ctx,
		creator,
		tenantID,
		cursor)
}

// NewEdgeClusterTypeConnectionResolver creates new EdgeClusterTypeConnectionResolverContract and returns it
// ctx: Mandatory. Reference to the context
// Returns the EdgeClusterTypeConnectionResolverContract or error if something goes wrong
func (creator *resolverCreator) NewEdgeClusterTypeConnectionResolver(
	ctx context.Context) (EdgeClusterTypeConnectionResolverContract, error) {
	return NewEdgeClusterTypeConnectionResolver(ctx, creator)
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
	hasPreviousPage bool) (PageInfoResolverContract, error) {
	return NewPageInfoResolver(
		startCursor,
		endCursor,
		hasNextPage,
		hasPreviousPage)
}
