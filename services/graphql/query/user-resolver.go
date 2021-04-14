// Package query implements different GraphQL query resovlers required by the GraphQL transport layer
package query

import (
	"context"
	"errors"
	"strings"

	"github.com/decentralized-cloud/api-gateway/services/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/edgecluster"
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/project"
	edgeClusterGrpcContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
	projectGrpcContract "github.com/decentralized-cloud/project/contract/grpc/go"
	"github.com/graph-gophers/graphql-go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"github.com/thoas/go-funk"
	"go.uber.org/zap"
)

type userResolver struct {
	logger                   *zap.Logger
	resolverCreator          types.ResolverCreatorContract
	userID                   string
	projectClientService     project.ProjectClientContract
	edgeClusterClientService edgecluster.EdgeClusterClientContract
}

// NewUserResolver creates new instance of the userResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can create new instances of resolvers
// logger: Mandatory. Reference to the logger service
// userID: Mandatory. the project unique identifier
// projectClientService: Mandatory. the project client service that creates gRPC connection and client to the project
// Returns the new instance or error if something goes wrong
func NewUserResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	logger *zap.Logger,
	userID string,
	projectClientService project.ProjectClientContract,
	edgeClusterClientService edgecluster.EdgeClusterClientContract) (types.UserResolverContract, error) {
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

	if projectClientService == nil {
		return nil, commonErrors.NewArgumentNilError("projectClientService", "projectClientService is required")
	}

	if edgeClusterClientService == nil {
		return nil, commonErrors.NewArgumentNilError("edgeClusterClientService", "edgeClusterClientService is required")
	}

	return &userResolver{
		logger:                   logger,
		resolverCreator:          resolverCreator,
		userID:                   userID,
		projectClientService:     projectClientService,
		edgeClusterClientService: edgeClusterClientService,
	}, nil
}

// ID returns user unique identifier
// ctx: Mandatory. Reference to the context
// Returns the user unique identifier
func (r *userResolver) ID(ctx context.Context) graphql.ID {
	return graphql.ID(r.userID)
}

// Project returns project resolver
// ctx: Mandatory. Reference to the context
// args: Mandatory. The argument list
// Returns the project resolver or error if something goes wrong
func (r *userResolver) Project(
	ctx context.Context,
	args types.UserProjectInputArgument) (project.ProjectResolverContract, error) {
	return r.resolverCreator.NewProjectResolver(
		ctx,
		string(args.ProjectID),
		nil)
}

// Projects returns project connection compatible with graphql-relay
// ctx: Mandatory. Reference to the context
// args: Mandatory. The argument list
// Returns the project resolver or error if something goes wrong
func (r *userResolver) Projects(
	ctx context.Context,
	args types.UserProjectsInputArgument) (project.ProjectTypeConnectionResolverContract, error) {

	pagination := projectGrpcContract.Pagination{}

	if args.After != nil {
		pagination.HasAfter = true
		pagination.After = *args.After
	}

	if args.First != nil {
		pagination.HasFirst = true
		pagination.First = *args.First
	}

	if args.Before != nil {
		pagination.HasBefore = true
		pagination.Before = *args.Before
	}

	if args.Last != nil {
		pagination.HasLast = true
		pagination.Last = *args.Last
	}

	sortingOptions := []*projectGrpcContract.SortingOptionPair{}

	if args.SortingOptions != nil {
		sortingOptions = funk.Map(*args.SortingOptions, func(sortingOption types.SortingOptionPair) *projectGrpcContract.SortingOptionPair {
			direction := projectGrpcContract.SortingDirection_ASCENDING

			if sortingOption.Direction == "DESCENDING" {
				direction = projectGrpcContract.SortingDirection_DESCENDING
			}

			return &projectGrpcContract.SortingOptionPair{
				Name:      sortingOption.Name,
				Direction: direction,
			}
		}).([]*projectGrpcContract.SortingOptionPair)
	}

	projectIDs := []string{}
	if args.ProjectIDs != nil {
		projectIDs = funk.Map(*args.ProjectIDs, func(projectID graphql.ID) string {
			return string(projectID)
		}).([]string)
	}

	connection, projectServiceClient, err := r.projectClientService.CreateClient()
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = connection.Close()
	}()

	response, err := projectServiceClient.ListProjects(
		ctx,
		&projectGrpcContract.ListProjectsRequest{
			Pagination:     &pagination,
			SortingOptions: sortingOptions,
			ProjectIDs:     projectIDs,
		})
	if err != nil {
		return nil, err
	}

	if response.Error != projectGrpcContract.Error_NO_ERROR {
		return nil, errors.New(response.ErrorMessage)
	}

	return r.resolverCreator.NewProjectTypeConnectionResolver(
		ctx,
		response.Projects,
		response.HasPreviousPage,
		response.HasNextPage,
		int32(response.TotalCount),
	)
}

// EdgeCluster returns project resolver
// ctx: Mandatory. Reference to the context
// args: Mandatory. The argument list
// Returns the project resolver or error if something goes wrong
func (r *userResolver) EdgeCluster(
	ctx context.Context,
	args types.UserEdgeClusterInputArgument) (edgecluster.EdgeClusterResolverContract, error) {
	return r.resolverCreator.NewEdgeClusterResolver(
		ctx,
		string(args.EdgeClusterID),
		nil)
}

// EdgeClusters returns project connection compatible with graphql-relay
// ctx: Mandatory. Reference to the context
// args: Mandatory. The argument list
// Returns the project resolver or error if something goes wrong
func (r *userResolver) EdgeClusters(
	ctx context.Context,
	args types.UserEdgeClustersInputArgument) (edgecluster.EdgeClusterTypeConnectionResolverContract, error) {
	pagination := edgeClusterGrpcContract.Pagination{}

	if args.After != nil {
		pagination.HasAfter = true
		pagination.After = *args.After
	}

	if args.First != nil {
		pagination.HasFirst = true
		pagination.First = *args.First
	}

	if args.Before != nil {
		pagination.HasBefore = true
		pagination.Before = *args.Before
	}

	if args.Last != nil {
		pagination.HasLast = true
		pagination.Last = *args.Last
	}

	sortingOptions := []*edgeClusterGrpcContract.SortingOptionPair{}

	if args.SortingOptions != nil {
		sortingOptions = funk.Map(*args.SortingOptions, func(sortingOption types.SortingOptionPair) *edgeClusterGrpcContract.SortingOptionPair {
			direction := edgeClusterGrpcContract.SortingDirection_ASCENDING

			if sortingOption.Direction == "DESCENDING" {
				direction = edgeClusterGrpcContract.SortingDirection_DESCENDING
			}

			return &edgeClusterGrpcContract.SortingOptionPair{
				Name:      sortingOption.Name,
				Direction: direction,
			}
		}).([]*edgeClusterGrpcContract.SortingOptionPair)
	}

	projectIDs := []string{}
	if args.ProjectIDs != nil {
		projectIDs = funk.Map(*args.ProjectIDs, func(projectID graphql.ID) string {
			return string(projectID)
		}).([]string)
	}

	edgeClusterIDs := []string{}
	if args.EdgeClusterIDs != nil {
		edgeClusterIDs = funk.Map(*args.EdgeClusterIDs, func(edgeClusterIDs graphql.ID) string {
			return string(edgeClusterIDs)
		}).([]string)
	}

	connection, edgeClusterServiceClient, err := r.edgeClusterClientService.CreateClient()
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = connection.Close()
	}()

	response, err := edgeClusterServiceClient.ListEdgeClusters(
		ctx,
		&edgeClusterGrpcContract.ListEdgeClustersRequest{
			Pagination:     &pagination,
			SortingOptions: sortingOptions,
			EdgeClusterIDs: edgeClusterIDs,
			ProjectIDs:     projectIDs,
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
		int32(response.TotalCount),
	)
}
