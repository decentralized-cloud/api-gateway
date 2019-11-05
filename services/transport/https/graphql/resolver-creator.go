// Package graphql implements functions to expose api-gateway service endpoint using GraphQL protocol.
package graphql

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/configuration"
	mutationedgecluster "github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/mutation/edgecluster"
	mutationtenant "github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/mutation/tenant"
	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/query"
	queryedgecluster "github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/query/edgecluster"
	queryrelay "github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/query/relay"
	querytenant "github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/query/tenant"
	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/root"
	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types/edgecluster"
	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types/relay"
	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types/tenant"
	tenantGrpcContract "github.com/decentralized-cloud/tenant/contract/grpc/go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type resolverCreator struct {
	logger                 *zap.Logger
	tenantClientConnection *grpc.ClientConn
}

// NewResolverCreator creates new instance of the resolverCreator, setting up all dependencies and returns the instance
// logger: Mandatory. Reference to the logger service
// configurationService: Mandatory. Reference to the configuration service
// Returns the new instance or error if something goes wrong
func NewResolverCreator(
	logger *zap.Logger,
	configurationService configuration.ConfigurationContract) (types.ResolverCreatorContract, error) {
	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if configurationService == nil {
		return nil, commonErrors.NewArgumentNilError("configurationService", "configurationService is required")
	}

	tenantServiceAddress, err := configurationService.GetTenantServiceAddress()
	if err != nil {
		return nil, err
	}

	// TODO: 20/10/2019 - Morteza, two things here, first we need to find out whether the below created connection can recover from failure
	// and second when to call connection.Close function.
	tenantClientConnection, err := grpc.Dial(tenantServiceAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &resolverCreator{
		logger:                 logger,
		tenantClientConnection: tenantClientConnection,
	}, nil
}

// NewPageInfoResolver creates new PageInfoResolverContract and returns it
// ctx: Mandatory. Reference to the context
// startCursor: Mandatory. Reference to the start cursor
// endCursor: Mandatory. Reference to the end cursor
// hasNextPage: Mandatory. Reference to the value indicates whether returned page has next page to be retrieved
// hasPreviousPage: Mandatory. Reference to the value indicates whether returned page has previous page to be retrieved
// Returns the PageInfoResolverContract or error if something goes wrong
func (creator *resolverCreator) NewPageInfoResolver(
	ctx context.Context,
	startCursor *string,
	endCursor *string,
	hasNextPage bool,
	hasPreviousPage bool) (relay.PageInfoResolverContract, error) {
	return queryrelay.NewPageInfoResolver(
		startCursor,
		endCursor,
		hasNextPage,
		hasPreviousPage)
}

// NewRootResolver creates new RootResolverContract and returns it
// ctx: Mandatory. Reference to the context
// Returns the RootResolverContract or error if something goes wrong
func (creator *resolverCreator) NewRootResolver(ctx context.Context) (types.RootResolverContract, error) {
	return root.NewRootResolver(
		ctx,
		creator,
		creator.logger)
}

// NewUserResolver creates new UserResolverContract and returns it
// ctx: Mandatory. Reference to the context
// userID: Mandatory. The user unique identifier
// Returns the UserResolverContract or error if something goes wrong
func (creator *resolverCreator) NewUserResolver(
	ctx context.Context,
	userID string) (types.UserResolverContract, error) {
	return query.NewUserResolver(
		ctx,
		creator,
		creator.logger,
		userID,
		tenantGrpcContract.NewTenantServiceClient(creator.tenantClientConnection))
}

// NewTenantResolver creates new TenantResolverContract and returns it
// ctx: Mandatory. Reference to the context
// tenantID: Mandatory. The tenant unique identifier
// tenant: Optional. The tenant details
// Returns the TenantResolverContract or error if something goes wrong
func (creator *resolverCreator) NewTenantResolver(
	ctx context.Context,
	tenantID string,
	tenant *tenantGrpcContract.Tenant) (tenant.TenantResolverContract, error) {
	return querytenant.NewTenantResolver(
		ctx,
		creator,
		creator.logger,
		tenantGrpcContract.NewTenantServiceClient(creator.tenantClientConnection),
		tenantID,
		tenant)
}

// NewTenantTypeEdgeResolver creates new TenantTypeEdgeResolverContract and returns it
// ctx: Mandatory. Reference to the context
// tenantID: Mandatory. The tenant unique identifier
// tenant: Optional. The tenant details
// cursor: Mandatory. The cursor
// Returns the TenantTypeEdgeResolverContract or error if something goes wrong
func (creator *resolverCreator) NewTenantTypeEdgeResolver(
	ctx context.Context,
	tenantID string,
	cursor string,
	tenant *tenantGrpcContract.Tenant) (tenant.TenantTypeEdgeResolverContract, error) {
	return querytenant.NewTenantTypeEdgeResolver(
		ctx,
		creator,
		tenantID,
		tenant,
		cursor)
}

// NewTenantTypeConnectionResolver creates new TenantTypeConnectionResolverContract and returns it
// ctx: Mandatory. Reference to the context
// tenants: Mandatory. Reference the list of tenants
// hasPreviousPage: Mandatory. Indicates whether more edges exist prior to the set defined by the clients arguments
// hasNextPage: Mandatory. Indicates whether more edges exist following the set defined by the clients arguments
// Returns the TenantTypeConnectionResolverContract or error if something goes wrong
func (creator *resolverCreator) NewTenantTypeConnectionResolver(
	ctx context.Context,
	tenants []*tenantGrpcContract.TenantWithCursor,
	hasPreviousPage bool,
	hasNextPage bool) (tenant.TenantTypeConnectionResolverContract, error) {
	return querytenant.NewTenantTypeConnectionResolver(
		ctx,
		creator,
		tenants,
		hasPreviousPage,
		hasNextPage)
}

// NewEdgeClusterResolver creates new EdgeClusterResolverContract and returns it
// ctx: Mandatory. Reference to the context
// tenantID: Mandatory. The tenant unique identifier
// Returns the EdgeClusterResolverContract or error if something goes wrong
func (creator *resolverCreator) NewEdgeClusterResolver(
	ctx context.Context,
	tenantID string) (edgecluster.EdgeClusterResolverContract, error) {
	return queryedgecluster.NewEdgeClusterResolver(
		ctx,
		creator,
		creator.logger,
		tenantID)
}

// NewEdgeClusterTypeEdgeResolver creates new EdgeClusterTypeEdgeResolverContract and returns it
// ctx: Mandatory. Reference to the context
// tenantID: Mandatory. The tenant unique identifier
// cursor: Mandatory. The cursor
// Returns the EdgeClusterTypeEdgeResolverContract or error if something goes wrong
func (creator *resolverCreator) NewEdgeClusterTypeEdgeResolver(
	ctx context.Context,
	tenantID string,
	cursor string) (edgecluster.EdgeClusterTypeEdgeResolverContract, error) {
	return queryedgecluster.NewEdgeClusterTypeEdgeResolver(
		ctx,
		creator,
		tenantID,
		cursor)
}

// NewEdgeClusterTypeConnectionResolver creates new EdgeClusterTypeConnectionResolverContract and returns it
// ctx: Mandatory. Reference to the context
// Returns the EdgeClusterTypeConnectionResolverContract or error if something goes wrong
func (creator *resolverCreator) NewEdgeClusterTypeConnectionResolver(ctx context.Context) (edgecluster.EdgeClusterTypeConnectionResolverContract, error) {
	return queryedgecluster.NewEdgeClusterTypeConnectionResolver(
		ctx,
		creator)
}

// NewEdgeClusterTenantResolver creates new EdgeClusterTenatnResolverContract and returns it
// ctx: Mandatory. Reference to the context
// tenantID: Mandatory. The tenant unique identifier
// Returns the EdgeClusterTenatnResolverContract or error if something goes wrong
func (creator *resolverCreator) NewEdgeClusterTenantResolver(
	ctx context.Context,
	tenantID string) (edgecluster.EdgeClusterTenantResolverContract, error) {
	return queryedgecluster.NewEdgeClusterTenantResolver(
		ctx,
		creator,
		creator.logger,
		tenantID)
}

// NewCreateTenant creates new instance of the createTenant, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewCreateTenant(ctx context.Context) (tenant.CreateTenantContract, error) {
	return mutationtenant.NewCreateTenant(
		ctx,
		creator,
		creator.logger,
		tenantGrpcContract.NewTenantServiceClient(creator.tenantClientConnection))
}

// NewCreateTenantPayloadResolver creates new instance of the createTenantPayloadResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// clientMutationId: Optional. Reference to the client mutation ID to correlate the request and response
// tenantID: Mandatory. The tenant unique identifier
// tenant: Optional. The tenant details
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewCreateTenantPayloadResolver(
	ctx context.Context,
	clientMutationId *string,
	tenantID string,
	tenant *tenantGrpcContract.Tenant) (tenant.CreateTenantPayloadResolverContract, error) {
	return mutationtenant.NewCreateTenantPayloadResolver(
		ctx,
		creator,
		clientMutationId,
		tenantID,
		tenant)
}

// NewUpdateTenant creates new instance of the updateTenant, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewUpdateTenant(ctx context.Context) (tenant.UpdateTenantContract, error) {
	return mutationtenant.NewUpdateTenant(
		ctx,
		creator,
		creator.logger,
		tenantGrpcContract.NewTenantServiceClient(creator.tenantClientConnection))
}

// NewUpdateTenantPayloadResolver creates new instance of the updateTenantPayloadResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// clientMutationId: Optional. Reference to the client mutation ID to correlate the request and response
// tenantID: Mandatory. The tenant unique identifier
// tenant: Optional. The tenant details
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewUpdateTenantPayloadResolver(
	ctx context.Context,
	clientMutationId *string,
	tenantID string,
	tenant *tenantGrpcContract.Tenant) (tenant.UpdateTenantPayloadResolverContract, error) {
	return mutationtenant.NewUpdateTenantPayloadResolver(
		ctx,
		creator,
		clientMutationId,
		tenantID,
		tenant)
}

// NewDeleteTenant creates new instance of the deleteTenant, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewDeleteTenant(ctx context.Context) (tenant.DeleteTenantContract, error) {
	return mutationtenant.NewDeleteTenant(
		ctx,
		creator,
		creator.logger,
		tenantGrpcContract.NewTenantServiceClient(creator.tenantClientConnection))
}

// NewDeleteTenantPayloadResolver creates new instance of the deleteTenantPayloadResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// clientMutationId: Optional. Reference to the client mutation ID to correlate the request and response
// tenantID: Mandatory. The tenant unique identifier
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewDeleteTenantPayloadResolver(
	ctx context.Context,
	clientMutationId *string,
	tenantID string) (tenant.DeleteTenantPayloadResolverContract, error) {
	return mutationtenant.NewDeleteTenantPayloadResolver(
		ctx,
		creator,
		clientMutationId,
		tenantID)
}

// NewCreateEdgeCluster creates new instance of the createEdgeCluster, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewCreateEdgeCluster(ctx context.Context) (edgecluster.CreateEdgeClusterContract, error) {
	return mutationedgecluster.NewCreateEdgeCluster(
		ctx,
		creator,
		creator.logger)
}

// NewCreateEdgeClusterPayloadResolver creates new instance of the createEdgeClusterPayloadResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// clientMutationId: Optional. Reference to the client mutation ID to correlate the request and response
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewCreateEdgeClusterPayloadResolver(
	ctx context.Context,
	clientMutationId *string) (edgecluster.CreateEdgeClusterPayloadResolverContract, error) {
	return mutationedgecluster.NewCreateEdgeClusterPayloadResolver(
		ctx,
		creator,
		clientMutationId)
}

// NewUpdateEdgeCluster creates new instance of the updateEdgeCluster, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewUpdateEdgeCluster(ctx context.Context) (edgecluster.UpdateEdgeClusterContract, error) {
	return mutationedgecluster.NewUpdateEdgeCluster(
		ctx,
		creator,
		creator.logger)
}

// NewUpdateEdgeClusterPayloadResolver creates new instance of the updateEdgeClusterPayloadResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// clientMutationId: Optional. Reference to the client mutation ID to correlate the request and response
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewUpdateEdgeClusterPayloadResolver(
	ctx context.Context,
	clientMutationId *string) (edgecluster.UpdateEdgeClusterPayloadResolverContract, error) {
	return mutationedgecluster.NewUpdateEdgeClusterPayloadResolver(
		ctx,
		creator,
		clientMutationId)
}

// NewDeleteEdgeCluster creates new instance of the deleteEdgeCluster, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewDeleteEdgeCluster(ctx context.Context) (edgecluster.DeleteEdgeClusterContract, error) {
	return mutationedgecluster.NewDeleteEdgeCluster(
		ctx,
		creator,
		creator.logger)
}

// NewDeleteEdgeClusterPayloadResolver creates new instance of the deleteEdgeClusterPayloadResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// clientMutationId: Optional. Reference to the client mutation ID to correlate the request and response
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewDeleteEdgeClusterPayloadResolver(
	ctx context.Context,
	clientMutationId *string) (edgecluster.DeleteEdgeClusterPayloadResolverContract, error) {
	return mutationedgecluster.NewDeleteEdgeClusterPayloadResolver(
		ctx,
		creator,
		clientMutationId)
}
