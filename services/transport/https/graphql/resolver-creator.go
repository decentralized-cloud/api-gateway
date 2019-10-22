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
	"github.com/graph-gophers/graphql-go"
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
	userID graphql.ID) (types.UserResolverContract, error) {
	return query.NewUserResolver(
		ctx,
		creator,
		creator.logger,
		userID)
}

// NewTenantResolver creates new TenantResolverContract and returns it
// ctx: Mandatory. Reference to the context
// tenantID: Mandatory. The tenant unique identifier
// Returns the TenantResolverContract or error if something goes wrong
func (creator *resolverCreator) NewTenantResolver(
	ctx context.Context,
	tenantID graphql.ID) (tenant.TenantResolverContract, error) {
	return querytenant.NewTenantResolver(
		ctx,
		creator,
		creator.logger,
		tenantGrpcContract.NewTenantServiceClient(creator.tenantClientConnection),
		tenantID)
}

// NewTenantTypeEdgeResolver creates new TenantTypeEdgeResolverContract and returns it
// ctx: Mandatory. Reference to the context
// tenantID: Mandatory. The tenant unique identifier
// cursor: Mandatory. The cursor
// Returns the TenantTypeEdgeResolverContract or error if something goes wrong
func (creator *resolverCreator) NewTenantTypeEdgeResolver(
	ctx context.Context,
	tenantID graphql.ID,
	cursor string) (tenant.TenantTypeEdgeResolverContract, error) {
	return querytenant.NewTenantTypeEdgeResolver(
		ctx,
		creator,
		tenantID,
		cursor)
}

// NewTenantTypeConnectionResolver creates new TenantTypeConnectionResolverContract and returns it
// ctx: Mandatory. Reference to the context
// Returns the TenantTypeConnectionResolverContract or error if something goes wrong
func (creator *resolverCreator) NewTenantTypeConnectionResolver(ctx context.Context) (tenant.TenantTypeConnectionResolverContract, error) {
	return querytenant.NewTenantTypeConnectionResolver(
		ctx,
		creator)
}

// NewEdgeClusterResolver creates new EdgeClusterResolverContract and returns it
// ctx: Mandatory. Reference to the context
// tenantID: Mandatory. The tenant unique identifier
// Returns the EdgeClusterResolverContract or error if something goes wrong
func (creator *resolverCreator) NewEdgeClusterResolver(
	ctx context.Context,
	tenantID graphql.ID) (edgecluster.EdgeClusterResolverContract, error) {
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
	tenantID graphql.ID,
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
	tenantID graphql.ID) (edgecluster.EdgeClusterTenantResolverContract, error) {
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
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewCreateTenantPayloadResolver(
	ctx context.Context,
	clientMutationId *string,
	tenantID string) (tenant.CreateTenantPayloadResolverContract, error) {
	return mutationtenant.NewCreateTenantPayloadResolver(
		ctx,
		creator,
		clientMutationId,
		tenantID)
}

// NewUpdateTenant creates new instance of the updateTenant, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewUpdateTenant(ctx context.Context) (tenant.UpdateTenantContract, error) {
	return mutationtenant.NewUpdateTenant(
		ctx,
		creator,
		creator.logger)
}

// NewUpdateTenantPayloadResolver creates new instance of the updateTenantPayloadResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// clientMutationId: Optional. Reference to the client mutation ID to correlate the request and response
// Returns the new instance or error if something goes wrong
func (creator *resolverCreator) NewUpdateTenantPayloadResolver(
	ctx context.Context,
	clientMutationId *string) (tenant.UpdateTenantPayloadResolverContract, error) {
	return mutationtenant.NewUpdateTenantPayloadResolver(
		ctx,
		creator,
		clientMutationId)
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
