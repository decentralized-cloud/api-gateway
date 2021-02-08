// Package project implements project mutation required by the GraphQL transport layer
package project

import (
	"context"
	"errors"
	"strings"

	"github.com/decentralized-cloud/api-gateway/services/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/project"
	projectGrpcContract "github.com/decentralized-cloud/project/contract/grpc/go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type updateProject struct {
	logger               *zap.Logger
	resolverCreator      types.ResolverCreatorContract
	projectClientService project.ProjectClientContract
}

type updateProjectPayloadResolver struct {
	resolverCreator  types.ResolverCreatorContract
	clientMutationId *string
	projectID        string
	projectDetail    *project.ProjectDetail
	cursor           string
}

// NewUpdateProject updates new instance of the updateProject, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can update new instances of resolvers
// logger: Mandatory. Reference to the logger service
// projectClientService: Mandatory. the project client service that creates gRPC connection and client to the project
// Returns the new instance or error if something goes wrong
func NewUpdateProject(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	logger *zap.Logger,
	projectClientService project.ProjectClientContract) (project.UpdateProjectContract, error) {
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

	return &updateProject{
		logger:               logger,
		resolverCreator:      resolverCreator,
		projectClientService: projectClientService,
	}, nil
}

// NewUpdateProjectPayloadResolver updates new instance of the updateProjectPayloadResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can update new instances of resolvers
// clientMutationId: Optional. Reference to the client mutation ID
// projectID: Mandatory. The project unique identifier
// projectDetail: Mandatory. The project details
// cursor: Mandatory. The edge cluster cursor
// Returns the new instance or error if something goes wrong
func NewUpdateProjectPayloadResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	clientMutationId *string,
	projectID string,
	projectDetail *project.ProjectDetail,
	cursor string) (project.UpdateProjectPayloadResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if strings.Trim(projectID, " ") == "" {
		return nil, commonErrors.NewArgumentError("projectID", "projectID is required")
	}

	if projectDetail == nil {
		return nil, commonErrors.NewArgumentNilError("projectDetail", "projectDetail is required")
	}

	if strings.Trim(cursor, " ") == "" {
		return nil, commonErrors.NewArgumentError("cursor", "cursor is required")
	}

	return &updateProjectPayloadResolver{
		resolverCreator:  resolverCreator,
		clientMutationId: clientMutationId,
		projectID:        projectID,
		projectDetail:    projectDetail,
		cursor:           cursor,
	}, nil
}

// MutateAndGetPayload update an existing project and returns the payload contains the result of updating an existing project
// ctx: Mandatory. Reference to the context
// args: Mandatory. Reference to the input argument contains project information to update
// Returns the updated project payload or error if something goes wrong
func (m *updateProject) MutateAndGetPayload(
	ctx context.Context,
	args project.UpdateProjectInputArgument) (project.UpdateProjectPayloadResolverContract, error) {
	projectID := string(args.Input.ProjectID)
	connection, projectServiceClient, err := m.projectClientService.CreateClient()
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = connection.Close()
	}()

	response, err := projectServiceClient.UpdateProject(
		ctx,
		&projectGrpcContract.UpdateProjectRequest{
			ProjectID: projectID,
			Project: &projectGrpcContract.Project{
				Name: args.Input.Name,
			}})

	if err != nil {
		return nil, err
	}

	if response.Error != projectGrpcContract.Error_NO_ERROR {
		return nil, errors.New(response.ErrorMessage)
	}

	return m.resolverCreator.NewUpdateProjectPayloadResolver(
		ctx,
		args.Input.ClientMutationId,
		projectID,
		&project.ProjectDetail{
			Project: response.Project,
		},
		response.Cursor)
}

// Project returns the updated project inforamtion
// ctx: Mandatory. Reference to the context
// Returns the updated project inforamtion
func (r *updateProjectPayloadResolver) Project(ctx context.Context) (project.ProjectTypeEdgeResolverContract, error) {
	return r.resolverCreator.NewProjectTypeEdgeResolver(
		ctx,
		r.projectID,
		r.cursor,
		r.projectDetail)
}

// ClientMutationId returns the client mutation ID that was provided as part of the mutation request
// ctx: Mandatory. Reference to the context
// Returns the provided clientMutationId as part of mutation request
func (r *updateProjectPayloadResolver) ClientMutationId(ctx context.Context) *string {
	return r.clientMutationId
}
