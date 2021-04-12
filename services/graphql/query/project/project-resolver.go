// Package project implements different project GraphQL query resovlers required by the GraphQL transport layer
package project

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
	"go.uber.org/zap"
)

type projectResolver struct {
	logger                   *zap.Logger
	resolverCreator          types.ResolverCreatorContract
	projectID                string
	projectDetail            *project.ProjectDetail
	edgeClusterClientService edgecluster.EdgeClusterClientContract
}

// NewProjectResolver creates new instance of the projectResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can create new instances of resolvers
// logger: Mandatory. Reference to the logger service
// projectClientService: Mandatory. the project client service that creates gRPC connection and client to the project
// projectID: Mandatory. the project unique identifier
// projectDetail: Optional. The tennat details, if provided, the value be used instead of contacting  the edge cluster service
// Returns the new instance or error if something goes wrong
func NewProjectResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	logger *zap.Logger,
	projectClientService project.ProjectClientContract,
	edgeClusterClientService edgecluster.EdgeClusterClientContract,
	projectID string,
	projectDetail *project.ProjectDetail) (project.ProjectResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if projectClientService == nil {
		return nil, commonErrors.NewArgumentNilError("projectClientService", "projectClientService is required")
	}

	if edgeClusterClientService == nil {
		return nil, commonErrors.NewArgumentNilError("edgeClusterClientService", "edgeClusterClientService is required")
	}

	if strings.Trim(projectID, " ") == "" {
		return nil, commonErrors.NewArgumentError("projectID", "projectID is required")
	}

	resolver := projectResolver{
		logger:                   logger,
		resolverCreator:          resolverCreator,
		edgeClusterClientService: edgeClusterClientService,
		projectID:                projectID,
	}

	if projectDetail == nil {
		connection, projectServiceClient, err := projectClientService.CreateClient()
		if err != nil {
			return nil, err
		}

		defer func() {
			_ = connection.Close()
		}()

		response, err := projectServiceClient.ReadProject(
			ctx,
			&projectGrpcContract.ReadProjectRequest{
				ProjectID: projectID,
			})
		if err != nil {
			return nil, err
		}

		if response.Error != projectGrpcContract.Error_NO_ERROR {
			return nil, errors.New(response.ErrorMessage)
		}

		resolver.projectDetail = &project.ProjectDetail{
			Project: response.Project,
		}
	} else {
		resolver.projectDetail = projectDetail
	}

	return &resolver, nil
}

// ID returns project unique identifier
// ctx: Mandatory. Reference to the context
// Returns the project unique identifier
func (r *projectResolver) ID(ctx context.Context) graphql.ID {
	return graphql.ID(r.projectID)
}

// Name returns project name
// ctx: Mandatory. Reference to the context
// Returns the project name or error
func (r *projectResolver) Name(ctx context.Context) string {
	return r.projectDetail.Project.Name
}

// EdgeCluster returns project resolver
// ctx: Mandatory. Reference to the context
// args: Mandatory. The argument list
// Returns the project resolver or error if something goes wrong
func (r *projectResolver) EdgeCluster(
	ctx context.Context,
	args project.ProjectClusterEdgeClusterInputArgument) (edgecluster.EdgeClusterResolverContract, error) {
	return r.resolverCreator.NewEdgeClusterResolver(
		ctx,
		string(args.EdgeClusterID),
		nil)
}

// EdgeClusters returns project connection compatible with graphql-relay
// ctx: Mandatory. Reference to the context
// args: Mandatory. The argument list
// Returns the project resolver or error if something goes wrong
func (r *projectResolver) EdgeClusters(
	ctx context.Context,
	args project.ProjectEdgeClustersInputArgument) (edgecluster.EdgeClusterTypeConnectionResolverContract, error) {
	projectIDs := []string{r.projectID}
	sortingOptions := []*edgeClusterGrpcContract.SortingOptionPair{}

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
			Pagination: &edgeClusterGrpcContract.Pagination{
				After:  "",
				First:  1000,
				Before: "",
				Last:   0,
			},
			SortingOptions: sortingOptions,
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
