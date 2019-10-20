// Package tenant implements tenant mutation required by the GraphQL transport layer
package tenant

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types/tenant"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type deleteTenant struct {
	logger          *zap.Logger
	resolverCreator types.ResolverCreatorContract
}

type deleteTenantPayloadResolver struct {
	resolverCreator  types.ResolverCreatorContract
	clientMutationId *string
}

// NewDeleteTenant deletes new instance of the deleteTenant, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can delete new instances of resolvers
// logger: Mandatory. Reference to the logger service
// Returns the new instance or error if something goes wrong
func NewDeleteTenant(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	logger *zap.Logger) (tenant.DeleteTenantContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	return &deleteTenant{
		logger:          logger,
		resolverCreator: resolverCreator,
	}, nil
}

// NewDeleteTenantPayloadResolver updates new instance of the deleteTenantPayloadResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can update new instances of resolvers
// clientMutationId: Optional. Reference to the client mutation ID
// Returns the new instance or error if something goes wrong
func NewDeleteTenantPayloadResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	clientMutationId *string) (tenant.DeleteTenantPayloadResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	return &deleteTenantPayloadResolver{
		resolverCreator:  resolverCreator,
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
	return m.resolverCreator.NewDeleteTenantPayloadResolver(ctx, args.Input.ClientMutationId)
}

// ClientMutationId returns the client mutation ID that was provided as part of the mutation request
// ctx: Mandatory. Reference to the context
// Returns the provided clientMutationId as part of mutation request
func (r *deleteTenantPayloadResolver) ClientMutationId(ctx context.Context) *string {
	return r.clientMutationId
}
