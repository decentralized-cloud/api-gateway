// Package tenant implements tenant mutation required by the GraphQL transport layer
package tenant

import (
	"context"
	"errors"

	"github.com/decentralized-cloud/api-gateway/services/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/tenant"
	tenantGrpcContract "github.com/decentralized-cloud/tenant/contract/grpc/go"
	"github.com/graph-gophers/graphql-go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type deleteTenant struct {
	logger              *zap.Logger
	resolverCreator     types.ResolverCreatorContract
	tenantClientService tenant.TenantClientContract
}

type deleteTenantPayloadResolver struct {
	resolverCreator  types.ResolverCreatorContract
	tenantID         string
	clientMutationId *string
}

// NewDeleteTenant deletes new instance of the deleteTenant, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can delete new instances of resolvers
// logger: Mandatory. Reference to the logger service
// tenantClientService: Mandatory. the tenant client service that creates gRPC connection and client to the tenant
// Returns the new instance or error if something goes wrong
func NewDeleteTenant(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	logger *zap.Logger,
	tenantClientService tenant.TenantClientContract) (tenant.DeleteTenantContract, error) {
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

	return &deleteTenant{
		logger:              logger,
		resolverCreator:     resolverCreator,
		tenantClientService: tenantClientService,
	}, nil
}

// NewDeleteTenantPayloadResolver updates new instance of the deleteTenantPayloadResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can update new instances of resolvers
// tenantID: Mandatory. The tenant unique identifier
// clientMutationId: Optional. Reference to the client mutation ID
// Returns the new instance or error if something goes wrong
func NewDeleteTenantPayloadResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	tenantID string,
	clientMutationId *string) (tenant.DeleteTenantPayloadResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	return &deleteTenantPayloadResolver{
		resolverCreator:  resolverCreator,
		tenantID:         tenantID,
		clientMutationId: clientMutationId,
	}, nil
}

// MutateAndGetPayload delete an existing tenant and returns the payload contains the result of deleting an existing tenant
// ctx: Mandatory. Reference to the context
// args: Mandatory. Reference to the input argument contains tenant information to delete
// Returns the deleted tenant payload or error if something goes wrong
func (m *deleteTenant) MutateAndGetPayload(
	ctx context.Context,
	args tenant.DeleteTenantInputArgument) (tenant.DeleteTenantPayloadResolverContract, error) {
	tenantID := string(args.Input.TenantID)
	connection, tenantServiceClient, err := m.tenantClientService.CreateClient()
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = connection.Close()
	}()

	response, err := tenantServiceClient.DeleteTenant(
		ctx,
		&tenantGrpcContract.DeleteTenantRequest{
			TenantID: tenantID,
		})
	if err != nil {
		return nil, err
	}

	if response.Error != tenantGrpcContract.Error_NO_ERROR {
		return nil, errors.New(response.ErrorMessage)
	}

	return m.resolverCreator.NewDeleteTenantPayloadResolver(
		ctx,
		tenantID,
		args.Input.ClientMutationId,
	)
}

// DeletedTenantID returns the unique identifier of the tenant that got deleted
// ctx: Mandatory. Reference to the context
// Returns the unique identifier of the the tenant that got deleted
func (r *deleteTenantPayloadResolver) DeletedTenantID(ctx context.Context) graphql.ID {
	return graphql.ID(r.tenantID)
}

// ClientMutationId returns the client mutation ID that was provided as part of the mutation request
// ctx: Mandatory. Reference to the context
// Returns the provided clientMutationId as part of mutation request
func (r *deleteTenantPayloadResolver) ClientMutationId(ctx context.Context) *string {
	return r.clientMutationId
}