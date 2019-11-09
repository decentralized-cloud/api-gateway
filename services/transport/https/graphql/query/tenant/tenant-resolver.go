// Package tenant implements different tenant GraphQL query resovlers required by the GraphQL transport layer
package tenant

import (
	"context"
	"errors"
	"strings"

	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types/edgecluster"
	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types/tenant"
	edgeClusterGrpcContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
	tenantGrpcContract "github.com/decentralized-cloud/tenant/contract/grpc/go"
	"github.com/graph-gophers/graphql-go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type tenantResolver struct {
	logger                   *zap.Logger
	resolverCreator          types.ResolverCreatorContract
	tenantID                 string
	tenant                   *tenantGrpcContract.Tenant
	edgeClusterClientService edgecluster.EdgeClusterClientContract
}

// NewTenantResolver creates new instance of the tenantResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can create new instances of resolvers
// logger: Mandatory. Reference to the logger service
// tenantClientService: Mandatory. the tenant client service that creates gRPC connection and client to the tenant
// tenantID: Mandatory. the tenant unique identifier
// tenant: Optional. The tenant details
// Returns the new instance or error if something goes wrong
func NewTenantResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	logger *zap.Logger,
	tenantClientService tenant.TenantClientContract,
	edgeClusterClientService edgecluster.EdgeClusterClientContract,
	tenantID string,
	tenant *tenantGrpcContract.Tenant) (tenant.TenantResolverContract, error) {
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

	if edgeClusterClientService == nil {
		return nil, commonErrors.NewArgumentNilError("edgeClusterClientService", "edgeClusterClientService is required")
	}

	if strings.Trim(tenantID, " ") == "" {
		return nil, commonErrors.NewArgumentError("tenantID", "tenantID is required")
	}

	resolver := tenantResolver{
		logger:                   logger,
		resolverCreator:          resolverCreator,
		edgeClusterClientService: edgeClusterClientService,
		tenantID:                 tenantID,
	}

	if tenant == nil {
		connection, tenantServiceClient, err := tenantClientService.CreateClient()
		if err != nil {
			return nil, err
		}

		defer func() {
			_ = connection.Close()
		}()

		response, err := tenantServiceClient.ReadTenant(
			ctx,
			&tenantGrpcContract.ReadTenantRequest{
				TenantID: tenantID,
			})
		if err != nil {
			return nil, err
		}

		if response.Error != tenantGrpcContract.Error_NO_ERROR {
			return nil, errors.New(response.ErrorMessage)
		}

		resolver.tenant = response.Tenant
	} else {
		resolver.tenant = tenant
	}

	return &resolver, nil
}

// ID returns tenant unique identifier
// ctx: Mandatory. Reference to the context
// Returns the tenant unique identifier
func (r *tenantResolver) ID(ctx context.Context) graphql.ID {
	return graphql.ID(r.tenantID)
}

// Name returns tenant name
// ctx: Mandatory. Reference to the context
// Returns the tenant name or error
func (r *tenantResolver) Name(ctx context.Context) string {
	return r.tenant.Name
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
		string(args.EdgeClusterID),
		nil)
}

// EdgeClusters returns tenant connection compatible with graphql-relay
// ctx: Mandatory. Reference to the context
// args: Mandatory. The argument list
// Returns the tenant resolver or error if something goes wrong
func (r *tenantResolver) EdgeClusters(
	ctx context.Context,
	args tenant.TenantEdgeClustersInputArgument) (edgecluster.EdgeClusterTypeConnectionResolverContract, error) {
	tenantIDs := []string{r.tenantID}
	sortingOptions := []*edgeClusterGrpcContract.SortingOptionPair{}

	connection, edgeClusterServiceClient, err := r.edgeClusterClientService.CreateClient()
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = connection.Close()
	}()

	response, err := edgeClusterServiceClient.Search(
		ctx,
		&edgeClusterGrpcContract.SearchRequest{
			Pagination: &edgeClusterGrpcContract.Pagination{
				After:  "",
				First:  1000,
				Before: "",
				Last:   0,
			},
			SortingOptions: sortingOptions,
			TenantIDs:      tenantIDs,
		})
	if err != nil {
		return nil, err
	}

	if response.Error != edgeClusterGrpcContract.Error_NO_ERROR {
		return nil, errors.New(response.ErrorMessage)
	}

	return r.resolverCreator.NewEdgeClusterTypeConnectionResolver(
		ctx,
		response.EdgeClusters,
		response.HasPreviousPage,
		response.HasNextPage,
	)
}
