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

type createProject struct {
	logger               *zap.Logger
	resolverCreator      types.ResolverCreatorContract
	projectClientService project.ProjectClientContract
}

type createProjectPayloadResolver struct {
	resolverCreator  types.ResolverCreatorContract
	clientMutationId *string
	projectID        string
	projectDetail    *project.ProjectDetail
	cursor           string
}

// NewCreateProject creates new instance of the createProject, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can create new instances of resolvers
// logger: Mandatory. Reference to the logger service
// projectClientService: Mandatory. the project client service that creates gRPC connection and client to the project
// Returns the new instance or error if something goes wrong
func NewCreateProject(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	logger *zap.Logger,
	projectClientService project.ProjectClientContract) (project.CreateProjectContract, error) {
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

	return &createProject{
		logger:               logger,
		resolverCreator:      resolverCreator,
		projectClientService: projectClientService,
	}, nil
}

// NewCreateProjectPayloadResolver updates new instance of the createProjectPayloadResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can update new instances of resolvers
// clientMutationId: Optional. Reference to the client mutation ID
// projectID: Mandatory. The project unique identifier
// projectDetail: Mandatory. The project details
// cursor: Mandatory. The edge cluster cursor
// Returns the new instance or error if something goes wrong
func NewCreateProjectPayloadResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	clientMutationId *string,
	projectID string,
	projectDetail *project.ProjectDetail,
	cursor string) (project.CreateProjectPayloadResolverContract, error) {
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

	return &createProjectPayloadResolver{
		resolverCreator:  resolverCreator,
		clientMutationId: clientMutationId,
		projectID:        projectID,
		projectDetail:    projectDetail,
		cursor:           cursor,
	}, nil
}

// MutateAndGetPayload creates a new project and returns the payload contains the result of creating a new project
// ctx: Mandatory. Reference to the context
// args: Mandatory. Reference to the input argument contains project information to create
// Returns the new project payload or error if something goes wrong
func (m *createProject) MutateAndGetPayload(
	ctx context.Context,
	args project.CreateProjectInputArgument) (project.CreateProjectPayloadResolverContract, error) {
	connection, projectServiceClient, err := m.projectClientService.CreateClient()
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = connection.Close()
	}()

	response, err := projectServiceClient.CreateProject(
		ctx,
		&projectGrpcContract.CreateProjectRequest{
			Project: &projectGrpcContract.Project{
				Name: args.Input.Name,
			}})
	if err != nil {
		return nil, err
	}

	if response.Error != projectGrpcContract.Error_NO_ERROR {
		return nil, errors.New(response.ErrorMessage)
	}

	return m.resolverCreator.NewCreateProjectPayloadResolver(
		ctx,
		args.Input.ClientMutationId,
		response.ProjectID,
		&project.ProjectDetail{
			Project: response.Project,
		},
		response.Cursor)
}

// Project returns the new project inforamtion
// ctx: Mandatory. Reference to the context
// Returns the new project inforamtion
func (r *createProjectPayloadResolver) Project(ctx context.Context) (project.ProjectTypeEdgeResolverContract, error) {
	return r.resolverCreator.NewProjectTypeEdgeResolver(
		ctx,
		r.projectID,
		r.cursor,
		r.projectDetail)
}

// ClientMutationId returns the client mutation ID that was provided as part of the mutation request
// ctx: Mandatory. Reference to the context
// Returns the provided clientMutationId as part of mutation request
func (r *createProjectPayloadResolver) ClientMutationId(ctx context.Context) *string {
	return r.clientMutationId
}
