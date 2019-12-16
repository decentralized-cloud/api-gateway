// Package tenant implements different tenant GraphQL query resovlers required by the GraphQL transport layer
package tenant

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/relay"
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/tenant"
	tenantGrpcContract "github.com/decentralized-cloud/tenant/contract/grpc/go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"github.com/thoas/go-funk"
)

type tenantTypeConnectionResolver struct {
	resolverCreator types.ResolverCreatorContract
	tenants         []*tenantGrpcContract.TenantWithCursor
	hasPreviousPage bool
	hasNextPage     bool
	totalCount      int32
}

// NewTenantTypeConnectionResolver creates new instance of the tenantTypeConnectionResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can create new instances of resolvers
// tenants: Mandatory. Reference the list of tenants
// hasPreviousPage: Mandatory. Indicates whether more edges exist prior to the set defined by the clients arguments
// hasNextPage: Mandatory. Indicates whether more edges exist following the set defined by the clients arguments
// totalCount: Mandatory. The total count of matched tenants
// Returns the new instance or error if something goes wrong
func NewTenantTypeConnectionResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	tenants []*tenantGrpcContract.TenantWithCursor,
	hasPreviousPage bool,
	hasNextPage bool,
	totalCount int32) (tenant.TenantTypeConnectionResolverContract, error) {
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
		totalCount:      totalCount,
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

	for _, item := range tenants {
		edge, err := r.resolverCreator.NewTenantTypeEdgeResolver(
			ctx,
			item.TenantID,
			item.Cursor,
			&tenant.TenantDetail{
				Tenant: item.Tenant,
			})

		if err != nil {
			return nil, err
		}

		edges = append(edges, edge)
	}

	return &edges, nil
}

// TotalCount returns total count of the matched tenants
// ctx: Mandatory. Reference to the context
// Returns the total count of the matched tenants
func (r *tenantTypeConnectionResolver) TotalCount(ctx context.Context) *int32 {
	return &r.totalCount
}
