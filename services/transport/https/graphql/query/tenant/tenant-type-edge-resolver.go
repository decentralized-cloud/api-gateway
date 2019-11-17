// Package tenant implements different tenant GraphQL query resovlers required by the GraphQL transport layer
package tenant

import (
	"context"
	"strings"

	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types/tenant"
	commonErrors "github.com/micro-business/go-core/system/errors"
)

type tenantTypeEdgeResolver struct {
	resolverCreator types.ResolverCreatorContract
	tenantID        string
	tenantDetail    *tenant.TenantDetail
	cursor          string
}

// NewTenantTypeEdgeResolver creates new instance of the tenantTypeEdgeResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can create new instances of resolvers
// tenantID: Mandatory. the tenant unique identifier
// tenantDetail: Optional. The tenant details, if provided, the value be used instead of contacting  the edge cluster service
// cursor: Mandatory. the cursor
// Returns the new instance or error if something goes wrong
func NewTenantTypeEdgeResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	tenantID string,
	tenantDetail *tenant.TenantDetail,
	cursor string) (tenant.TenantTypeEdgeResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if strings.Trim(tenantID, " ") == "" {
		return nil, commonErrors.NewArgumentError("tenantID", "tenantID is required")
	}

	if strings.Trim(cursor, " ") == "" {
		return nil, commonErrors.NewArgumentError("cursor", "cursor is required")
	}

	return &tenantTypeEdgeResolver{
		resolverCreator: resolverCreator,
		tenantID:        tenantID,
		tenantDetail:    tenantDetail,
		cursor:          cursor,
	}, nil
}

// Node returns the tenant resolver
// ctx: Mandatory. Reference to the context
// Returns the tenant resolver or error if something goes wrong
func (r *tenantTypeEdgeResolver) Node(ctx context.Context) (tenant.TenantResolverContract, error) {
	return r.resolverCreator.NewTenantResolver(
		ctx,
		r.tenantID,
		r.tenantDetail)
}

// Cursor returns the cursor for the tenant edge compatible with graphql-relay
// ctx: Mandatory. Reference to the context
// Returns the cursor
func (r *tenantTypeEdgeResolver) Cursor(ctx context.Context) string {
	return r.cursor
}
