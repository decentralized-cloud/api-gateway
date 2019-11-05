// Package tenant implements different tenant GraphQL query resovlers required by the GraphQL transport layer
package tenant

import (
	"context"
	"strings"

	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types/tenant"
	tenantGrpcContract "github.com/decentralized-cloud/tenant/contract/grpc/go"
	commonErrors "github.com/micro-business/go-core/system/errors"
)

type tenantTypeEdgeResolver struct {
	resolverCreator types.ResolverCreatorContract
	tenantID        string
	tenant          *tenantGrpcContract.Tenant
	cursor          string
}

// NewTenantTypeEdgeResolver creates new instance of the tenantTypeEdgeResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can create new instances of resolvers
// tenantID: Mandatory. the tenant unique identifier
// tenant: Optional. The tenant details
// cursor: Mandatory. the cursor
// Returns the new instance or error if something goes wrong
func NewTenantTypeEdgeResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	tenantID string,
	tenant *tenantGrpcContract.Tenant,
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
		tenant:          tenant,
		cursor:          cursor,
	}, nil
}

// Node returns the tenant resolver
// ctx: Mandatory. Reference to the context
// Returns the tenant resolver or error if something goes wrong
func (r *tenantTypeEdgeResolver) Node(ctx context.Context) (tenant.TenantResolverContract, error) {
	return r.resolverCreator.NewTenantResolver(ctx, r.tenantID, r.tenant)
}

// Cursor returns the cursor for the tenant edge compatible with graphql-relay
// ctx: Mandatory. Reference to the context
// Returns the cursor
func (r *tenantTypeEdgeResolver) Cursor(ctx context.Context) string {
	return r.cursor
}
