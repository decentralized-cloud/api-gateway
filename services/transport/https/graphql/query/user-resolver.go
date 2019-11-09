// Package query implements different GraphQL query resovlers required by the GraphQL transport layer
package query

import (
	"context"
	"errors"
	"strings"

	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types/edgecluster"
	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types/tenant"
	tenantGrpcContract "github.com/decentralized-cloud/tenant/contract/grpc/go"
	"github.com/graph-gophers/graphql-go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"github.com/thoas/go-funk"
	"go.uber.org/zap"
)

type userResolver struct {
	logger              *zap.Logger
	resolverCreator     types.ResolverCreatorContract
	userID              string
	tenantClientService tenant.TenantClientContract
}

// NewUserResolver creates new instance of the userResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can create new instances of resolvers
// logger: Mandatory. Reference to the logger service
// userID: Mandatory. the tenant unique identifier
// tenantClientService: Mandatory. the tenant client service that creates gRPC connection and client to the tenant
// Returns the new instance or error if something goes wrong
func NewUserResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	logger *zap.Logger,
	userID string,
	tenantClientService tenant.TenantClientContract) (types.UserResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if strings.Trim(userID, " ") == "" {
		return nil, commonErrors.NewArgumentError("userID", "userID is required")
	}

	if tenantClientService == nil {
		return nil, commonErrors.NewArgumentNilError("tenantClientService", "tenantClientService is required")
	}

	return &userResolver{
		logger:              logger,
		resolverCreator:     resolverCreator,
		userID:              userID,
		tenantClientService: tenantClientService,
	}, nil
}

// ID returns user unique identifier
// ctx: Mandatory. Reference to the context
// Returns the user unique identifier
func (r *userResolver) ID(ctx context.Context) graphql.ID {
	return graphql.ID(r.userID)
}

// Tenant returns tenant resolver
// ctx: Mandatory. Reference to the context
// args: Mandatory. The argument list
// Returns the tenant resolver or error if something goes wrong
func (r *userResolver) Tenant(
	ctx context.Context,
	args types.UserTenantInputArgument) (tenant.TenantResolverContract, error) {
	return r.resolverCreator.NewTenantResolver(
		ctx,
		string(args.TenantID),
		nil)
}

// Tenants returns tenant connection compatible with graphql-relay
// ctx: Mandatory. Reference to the context
// args: Mandatory. The argument list
// Returns the tenant resolver or error if something goes wrong
func (r *userResolver) Tenants(
	ctx context.Context,
	args types.UserTenantsInputArgument) (tenant.TenantTypeConnectionResolverContract, error) {
	after := ""
	if args.After != nil {
		after = *args.After
	}

	var first int32 = 0
	if args.First != nil {
		first = *args.First
	}

	before := ""
	if args.Before != nil {
		before = *args.Before
	}

	var last int32 = 0
	if args.Last != nil {
		last = *args.Last
	}

	sortingOptions := []*tenantGrpcContract.SortingOptionPair{}

	if args.SortingOptions != nil {
		sortingOptions = funk.Map(*args.SortingOptions, func(sortingOption types.SortingOptionPair) *tenantGrpcContract.SortingOptionPair {
			direction := tenantGrpcContract.SortingDirection_ASCENDING

			if sortingOption.Direction == "DESCENDING" {
				direction = tenantGrpcContract.SortingDirection_DESCENDING
			}

			return &tenantGrpcContract.SortingOptionPair{
				Name:      sortingOption.Name,
				Direction: direction,
			}
		}).([]*tenantGrpcContract.SortingOptionPair)
	}

	tenantIDs := []string{}
	if args.TenantIDs != nil {
		tenantIDs = funk.Map(*args.TenantIDs, func(tenantID graphql.ID) string {
			return string(tenantID)
		}).([]string)
	}

	connection, tenantServiceClient, err := r.tenantClientService.CreateClient()
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = connection.Close()
	}()

	response, err := tenantServiceClient.Search(
		ctx,
		&tenantGrpcContract.SearchRequest{
			Pagination: &tenantGrpcContract.Pagination{
				After:  after,
				First:  first,
				Before: before,
				Last:   last,
			},
			SortingOptions: sortingOptions,
			TenantIDs:      tenantIDs,
		})
	if err != nil {
		return nil, err
	}

	if response.Error != tenantGrpcContract.Error_NO_ERROR {
		return nil, errors.New(response.ErrorMessage)
	}

	return r.resolverCreator.NewTenantTypeConnectionResolver(
		ctx,
		response.Tenants,
		response.HasPreviousPage,
		response.HasNextPage)
}

// EdgeCluster returns tenant resolver
// ctx: Mandatory. Reference to the context
// args: Mandatory. The argument list
// Returns the tenant resolver or error if something goes wrong
func (r *userResolver) EdgeCluster(
	ctx context.Context,
	args types.UserEdgeClusterInputArgument) (edgecluster.EdgeClusterResolverContract, error) {
	return r.resolverCreator.NewEdgeClusterResolver(
		ctx,
		string(args.EdgeClusterID))
}

// EdgeClusters returns tenant connection compatible with graphql-relay
// ctx: Mandatory. Reference to the context
// args: Mandatory. The argument list
// Returns the tenant resolver or error if something goes wrong
func (r *userResolver) EdgeClusters(
	ctx context.Context,
	args types.UserEdgeClustersInputArgument) (edgecluster.EdgeClusterTypeConnectionResolverContract, error) {
	return r.resolverCreator.NewEdgeClusterTypeConnectionResolver(ctx)
}
