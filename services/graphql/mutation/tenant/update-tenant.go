// Package tenant implements tenant mutation required by the GraphQL transport layer
package tenant

import (
	"context"
	"errors"
	"strings"

	"github.com/decentralized-cloud/api-gateway/services/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/tenant"
	tenantGrpcContract "github.com/decentralized-cloud/tenant/contract/grpc/go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type updateTenant struct {
	logger              *zap.Logger
	resolverCreator     types.ResolverCreatorContract
	tenantClientService tenant.TenantClientContract
}

type updateTenantPayloadResolver struct {
	resolverCreator  types.ResolverCreatorContract
	clientMutationId *string
	tenantID         string
	tenantDetail     *tenant.TenantDetail
	cursor           string
}

// NewUpdateTenant updates new instance of the updateTenant, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can update new instances of resolvers
// logger: Mandatory. Reference to the logger service
// tenantClientService: Mandatory. the tenant client service that creates gRPC connection and client to the tenant
// Returns the new instance or error if something goes wrong
func NewUpdateTenant(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	logger *zap.Logger,
	tenantClientService tenant.TenantClientContract) (tenant.UpdateTenantContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if tenantClientService == nil {
		return nil, commonErrors.NewArgumentNilError("tenantClientService", "tenantClientService is required")
	}

	return &updateTenant{
		logger:              logger,
		resolverCreator:     resolverCreator,
		tenantClientService: tenantClientService,
	}, nil
}

// NewUpdateTenantPayloadResolver updates new instance of the updateTenantPayloadResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can update new instances of resolvers
// clientMutationId: Optional. Reference to the client mutation ID
// tenantID: Mandatory. The tenant unique identifier
// tenantDetail: Mandatory. The tenant details
// cursor: Mandatory. The edge cluster cursor
// Returns the new instance or error if something goes wrong
func NewUpdateTenantPayloadResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	clientMutationId *string,
	tenantID string,
	tenantDetail *tenant.TenantDetail,
	cursor string) (tenant.UpdateTenantPayloadResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if strings.Trim(tenantID, " ") == "" {
		return nil, commonErrors.NewArgumentError("tenantID", "tenantID is required")
	}

	if tenantDetail == nil {
		return nil, commonErrors.NewArgumentNilError("tenantDetail", "tenantDetail is required")
	}

	if strings.Trim(cursor, " ") == "" {
		return nil, commonErrors.NewArgumentError("cursor", "cursor is required")
	}

	return &updateTenantPayloadResolver{
		resolverCreator:  resolverCreator,
		clientMutationId: clientMutationId,
		tenantID:         tenantID,
		tenantDetail:     tenantDetail,
		cursor:           cursor,
	}, nil
}

// MutateAndGetPayload update an existing tenant and returns the payload contains the result of updating an existing tenant
// ctx: Mandatory. Reference to the context
// args: Mandatory. Reference to the input argument contains tenant information to update
// Returns the updated tenant payload or error if something goes wrong
func (m *updateTenant) MutateAndGetPayload(
	ctx context.Context,
	args tenant.UpdateTenantInputArgument) (tenant.UpdateTenantPayloadResolverContract, error) {
	tenantID := string(args.Input.ID)
	connection, tenantServiceClient, err := m.tenantClientService.CreateClient()
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = connection.Close()
	}()

	response, err := tenantServiceClient.UpdateTenant(
		ctx,
		&tenantGrpcContract.UpdateTenantRequest{
			TenantID: tenantID,
			Tenant: &tenantGrpcContract.Tenant{
				Name: args.Input.Name,
			}})

	if err != nil {
		return nil, err
	}

	if response.Error != tenantGrpcContract.Error_NO_ERROR {
		return nil, errors.New(response.ErrorMessage)
	}

	return m.resolverCreator.NewUpdateTenantPayloadResolver(
		ctx,
		args.Input.ClientMutationId,
		tenantID,
		&tenant.TenantDetail{
			Tenant: response.Tenant,
		},
		response.Cursor)
}

// Tenant returns the updated tenant inforamtion
// ctx: Mandatory. Reference to the context
// Returns the updated tenant inforamtion
func (r *updateTenantPayloadResolver) Tenant(ctx context.Context) (tenant.TenantTypeEdgeResolverContract, error) {
	return r.resolverCreator.NewTenantTypeEdgeResolver(
		ctx,
		r.tenantID,
		r.cursor,
		r.tenantDetail)
}

// ClientMutationId returns the client mutation ID that was provided as part of the mutation request
// ctx: Mandatory. Reference to the context
// Returns the provided clientMutationId as part of mutation request
func (r *updateTenantPayloadResolver) ClientMutationId(ctx context.Context) *string {
	return r.clientMutationId
}
