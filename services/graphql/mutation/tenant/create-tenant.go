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

type createTenant struct {
	logger              *zap.Logger
	resolverCreator     types.ResolverCreatorContract
	tenantClientService tenant.TenantClientContract
}

type createTenantPayloadResolver struct {
	resolverCreator  types.ResolverCreatorContract
	clientMutationId *string
	tenantID         string
	tenantDetail     *tenant.TenantDetail
	cursor           string
}

// NewCreateTenant creates new instance of the createTenant, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can create new instances of resolvers
// logger: Mandatory. Reference to the logger service
// tenantClientService: Mandatory. the tenant client service that creates gRPC connection and client to the tenant
// Returns the new instance or error if something goes wrong
func NewCreateTenant(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	logger *zap.Logger,
	tenantClientService tenant.TenantClientContract) (tenant.CreateTenantContract, error) {
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

	return &createTenant{
		logger:              logger,
		resolverCreator:     resolverCreator,
		tenantClientService: tenantClientService,
	}, nil
}

// NewCreateTenantPayloadResolver updates new instance of the createTenantPayloadResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can update new instances of resolvers
// clientMutationId: Optional. Reference to the client mutation ID
// tenantID: Mandatory. The tenant unique identifier
// tenantDetail: Mandatory. The tenant details
// cursor: Mandatory. The edge cluster cursor
// Returns the new instance or error if something goes wrong
func NewCreateTenantPayloadResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	clientMutationId *string,
	tenantID string,
	tenantDetail *tenant.TenantDetail,
	cursor string) (tenant.CreateTenantPayloadResolverContract, error) {
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

	return &createTenantPayloadResolver{
		resolverCreator:  resolverCreator,
		clientMutationId: clientMutationId,
		tenantID:         tenantID,
		tenantDetail:     tenantDetail,
		cursor:           cursor,
	}, nil
}

// MutateAndGetPayload creates a new tenant and returns the payload contains the result of creating a new tenant
// ctx: Mandatory. Reference to the context
// args: Mandatory. Reference to the input argument contains tenant information to create
// Returns the new tenant payload or error if something goes wrong
func (m *createTenant) MutateAndGetPayload(
	ctx context.Context,
	args tenant.CreateTenantInputArgument) (tenant.CreateTenantPayloadResolverContract, error) {
	connection, tenantServiceClient, err := m.tenantClientService.CreateClient()
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = connection.Close()
	}()

	response, err := tenantServiceClient.CreateTenant(
		ctx,
		&tenantGrpcContract.CreateTenantRequest{
			Tenant: &tenantGrpcContract.Tenant{
				Name: args.Input.Name,
			}})
	if err != nil {
		return nil, err
	}

	if response.Error != tenantGrpcContract.Error_NO_ERROR {
		return nil, errors.New(response.ErrorMessage)
	}

	return m.resolverCreator.NewCreateTenantPayloadResolver(
		ctx,
		args.Input.ClientMutationId,
		response.TenantID,
		&tenant.TenantDetail{
			Tenant: response.Tenant,
		},
		response.Cursor)
}

// Tenant returns the new tenant inforamtion
// ctx: Mandatory. Reference to the context
// Returns the new tenant inforamtion
func (r *createTenantPayloadResolver) Tenant(ctx context.Context) (tenant.TenantTypeEdgeResolverContract, error) {
	return r.resolverCreator.NewTenantTypeEdgeResolver(
		ctx,
		r.tenantID,
		r.cursor,
		r.tenantDetail)
}

// ClientMutationId returns the client mutation ID that was provided as part of the mutation request
// ctx: Mandatory. Reference to the context
// Returns the provided clientMutationId as part of mutation request
func (r *createTenantPayloadResolver) ClientMutationId(ctx context.Context) *string {
	return r.clientMutationId
}
