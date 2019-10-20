// package tenant implements different tenant GraphQL query resovlers required by the GraphQL transport layer
package tenant

import (
	"context"
	"errors"
	"strings"

	"github.com/decentralized-cloud/api-gateway/services/configuration"
	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types/edgecluster"
	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types/tenant"
	tenantGrpcContract "github.com/decentralized-cloud/tenant/contract/grpc/go"
	"github.com/graph-gophers/graphql-go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type tenantResolver struct {
	logger          *zap.Logger
	resolverCreator types.ResolverCreatorContract
	tenantID        graphql.ID
	name            string
}

// NewTenantResolver creates new instance of the tenantResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can create new instances of resolvers
// logger: Mandatory. Reference to the logger service
// tenantID: Mandatory. the tenant unique identifier
// Returns the new instance or error if something goes wrong
func NewTenantResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	logger *zap.Logger,
	configurationService configuration.ConfigurationContract,
	tenantID graphql.ID) (tenant.TenantResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if configurationService == nil {
		return nil, commonErrors.NewArgumentNilError("configurationService", "configurationService is required")
	}

	if strings.Trim(string(tenantID), " ") == "" {
		return nil, commonErrors.NewArgumentError("tenantID", "tenantID is required")
	}

	tenantServiceAddress, err := configurationService.GetTenantServiceAddress()
	if err != nil {
		return nil, err
	}

	connection, err := grpc.Dial(tenantServiceAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer connection.Close()

	client := tenantGrpcContract.NewTenantServiceClient(connection)
	response, err := client.ReadTenant(
		ctx,
		&tenantGrpcContract.ReadTenantRequest{
			TenantID: string(tenantID),
		})
	if err != nil {
		return nil, err
	}

	if response.Error != tenantGrpcContract.Error_NO_ERROR {
		return nil, errors.New(response.ErrorMessage)
	}

	return &tenantResolver{
		logger:          logger,
		resolverCreator: resolverCreator,
		tenantID:        tenantID,
		name:            response.Tenant.Name,
	}, nil
}

// ID returns tenant unique identifier
// ctx: Mandatory. Reference to the context
// Returns the tenant unique identifier
func (r *tenantResolver) ID(ctx context.Context) graphql.ID {
	return r.tenantID
}

// Name returns tenant name
// ctx: Mandatory. Reference to the context
// Returns the tenant name or error
func (r *tenantResolver) Name(ctx context.Context) string {
	return r.name
}

// EdgeCluster returns tenant resolver
// ctx: Mandatory. Reference to the context
// args: Mandatory. The argument list
// Returns the tenant resolver or error if something goes wrong
func (r *tenantResolver) EdgeCluster(
	ctx context.Context,
	args tenant.TenantClusterEdgeClusterInputArgument) (edgecluster.EdgeClusterResolverContract, error) {
	return r.resolverCreator.NewEdgeClusterResolver(
		ctx,
		args.EdgeClusterID)
}

// EdgeClusters returns tenant conenction compatible with graphql-relay
// ctx: Mandatory. Reference to the context
// args: Mandatory. The argument list
// Returns the tenant resolver or error if something goes wrong
func (r *tenantResolver) EdgeClusters(
	ctx context.Context,
	args tenant.TenantEdgeClustersInputArgument) (edgecluster.EdgeClusterTypeConnectionResolverContract, error) {
	return r.resolverCreator.NewEdgeClusterTypeConnectionResolver(ctx)
}
