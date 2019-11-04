// Package tenant implements different tenant GraphQL query resovlers required by the GraphQL transport layer
package tenant

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types/relay"
	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types/tenant"
	tenantGrpcContract "github.com/decentralized-cloud/tenant/contract/grpc/go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"github.com/thoas/go-funk"
)

type tenantTypeConnectionResolver struct {
	resolverCreator types.ResolverCreatorContract
	tenants         []*tenantGrpcContract.TenantWithCursor
	hasPreviousPage bool
	hasNextPage     bool
}

// NewTenantTypeConnectionResolver creates new instance of the tenantTypeConnectionResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can create new instances of resolvers
// tenants: Mandatory. Reference the list of tenants
// hasPreviousPage: Mandatory. Indicates whether more edges exist prior to the set defined by the clients arguments
// hasNextPage: Mandatory. Indicates whether more edges exist following the set defined by the clients arguments
// Returns the new instance or error if something goes wrong
func NewTenantTypeConnectionResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	tenants []*tenantGrpcContract.TenantWithCursor,
	hasPreviousPage bool,
	hasNextPage bool) (tenant.TenantTypeConnectionResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	return &tenantTypeConnectionResolver{
		resolverCreator: resolverCreator,
		tenants:         tenants,
		hasPreviousPage: hasPreviousPage,
		hasNextPage:     hasNextPage,
	}, nil
}

// PageInfo returns the paging information compatible with graphql-relay
// ctx: Mandatory. Reference to the context
// Returns the paging information
func (r *tenantTypeConnectionResolver) PageInfo(ctx context.Context) (relay.PageInfoResolverContract, error) {
	var startCursor, endCursor string

	if len(r.tenants) > 0 {
		startCursor = r.tenants[0].Cursor
		endCursor = r.tenants[len(r.tenants)-1].Cursor
	}

	return r.resolverCreator.NewPageInfoResolver(
		ctx,
		&startCursor,
		&endCursor,
		r.hasNextPage,
		r.hasPreviousPage)
}

// Edges returns the tenant edges compatible with graphql-relay
// ctx: Mandatory. Reference to the context
// Returns the tenant edges
func (r *tenantTypeConnectionResolver) Edges(ctx context.Context) (*[]tenant.TenantTypeEdgeResolverContract, error) {
	tenants := funk.Filter(r.tenants, func(tenant *tenantGrpcContract.TenantWithCursor) bool {
		return tenant != nil
	}).([]*tenantGrpcContract.TenantWithCursor)

	edges := []tenant.TenantTypeEdgeResolverContract{}

	for _, tenant := range tenants {
		edge, err := r.resolverCreator.NewTenantTypeEdgeResolver(
			ctx,
			tenant.TenantID,
			tenant.Cursor,
			tenant.Tenant)

		if err != nil {
			return nil, err
		}

		edges = append(edges, edge)
	}

	return &edges, nil
}
