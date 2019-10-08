// package resolver implements different GraphQL resolvers required by the GraphQL transport layer
package resolver

import (
	"context"

	"github.com/graph-gophers/graphql-go"
	"github.com/lucsky/cuid"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

// QueryResolverContract declares the root query resolver
type QueryResolverContract interface {
	// User returns user resolver
	// ctx: Mandatory. Reference to the context
	// Returns the user resolver or error if something goes wrong
	User(ctx context.Context) (UserResolverContract, error)
}

type queryResolver struct {
	logger          *zap.Logger
	resolverCreator ResolverCreatorContract
}

// NewQueryResolver creates new instance of the queryResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can create new instances of resolvers
// logger: Mandatory. Reference to the logger service
// Returns the new instance or error if something goes wrong
func NewQueryResolver(
	ctx context.Context,
	resolverCreator ResolverCreatorContract,
	logger *zap.Logger) (QueryResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	return &queryResolver{
		logger:          logger,
		resolverCreator: resolverCreator,
	}, nil
}

// User returns user resolver
// ctx: Mandatory. Reference to the context
// Returns the user resolver or error if something goes wrong
func (r *queryResolver) User(ctx context.Context) (UserResolverContract, error) {
	return r.resolverCreator.NewUserResolver(ctx, graphql.ID(cuid.New()))
}
