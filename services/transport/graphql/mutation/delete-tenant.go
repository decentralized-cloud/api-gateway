// package resolver implements different GraphQL resolvers required by the GraphQL transport layer
package mutation

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/transport/graphql/types"
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
	logger *zap.Logger) (types.DeleteTenantContract, error) {
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

// Tenant returns the deleted tenant inforamtion
// ctx: Mandatory. Reference to the context
// Returns the deleted tenant inforamtion
func NewDeleteTenantPayloadResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	clientMutationId *string) (types.DeleteTenantPayloadResolverContract, error) {
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
	args types.DeleteTenantInputArgument) (types.DeleteTenantPayloadResolverContract, error) {
	return m.resolverCreator.NewDeleteTenantPayloadResolver(ctx, args.Input.ClientMutationId)
}

// ClientMutationId returns the client mutation ID that was provided as part of the mutation request
// ctx: Mandatory. Reference to the context
// Returns the provided clientMutationId as part of mutation request
func (r *deleteTenantPayloadResolver) ClientMutationId(ctx context.Context) *string {
	return r.clientMutationId
}
