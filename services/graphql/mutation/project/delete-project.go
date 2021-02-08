// Package project implements project mutation required by the GraphQL transport layer
package project

import (
	"context"
	"errors"

	"github.com/decentralized-cloud/api-gateway/services/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/project"
	projectGrpcContract "github.com/decentralized-cloud/project/contract/grpc/go"
	"github.com/graph-gophers/graphql-go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type deleteProject struct {
	logger               *zap.Logger
	resolverCreator      types.ResolverCreatorContract
	projectClientService project.ProjectClientContract
}

type deleteProjectPayloadResolver struct {
	resolverCreator  types.ResolverCreatorContract
	projectID        string
	clientMutationId *string
}

// NewDeleteProject deletes new instance of the deleteProject, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can delete new instances of resolvers
// logger: Mandatory. Reference to the logger service
// projectClientService: Mandatory. the project client service that creates gRPC connection and client to the project
// Returns the new instance or error if something goes wrong
func NewDeleteProject(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	logger *zap.Logger,
	projectClientService project.ProjectClientContract) (project.DeleteProjectContract, error) {
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

	return &deleteProject{
		logger:               logger,
		resolverCreator:      resolverCreator,
		projectClientService: projectClientService,
	}, nil
}

// NewDeleteProjectPayloadResolver updates new instance of the deleteProjectPayloadResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can update new instances of resolvers
// projectID: Mandatory. The project unique identifier
// clientMutationId: Optional. Reference to the client mutation ID
// Returns the new instance or error if something goes wrong
func NewDeleteProjectPayloadResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	projectID string,
	clientMutationId *string) (project.DeleteProjectPayloadResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	return &deleteProjectPayloadResolver{
		resolverCreator:  resolverCreator,
		projectID:        projectID,
		clientMutationId: clientMutationId,
	}, nil
}

// MutateAndGetPayload delete an existing project and returns the payload contains the result of deleting an existing project
// ctx: Mandatory. Reference to the context
// args: Mandatory. Reference to the input argument contains project information to delete
// Returns the deleted project payload or error if something goes wrong
func (m *deleteProject) MutateAndGetPayload(
	ctx context.Context,
	args project.DeleteProjectInputArgument) (project.DeleteProjectPayloadResolverContract, error) {
	projectID := string(args.Input.ProjectID)
	connection, projectServiceClient, err := m.projectClientService.CreateClient()
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = connection.Close()
	}()

	response, err := projectServiceClient.DeleteProject(
		ctx,
		&projectGrpcContract.DeleteProjectRequest{
			ProjectID: projectID,
		})
	if err != nil {
		return nil, err
	}

	if response.Error != projectGrpcContract.Error_NO_ERROR {
		return nil, errors.New(response.ErrorMessage)
	}

	return m.resolverCreator.NewDeleteProjectPayloadResolver(
		ctx,
		projectID,
		args.Input.ClientMutationId,
	)
}

// DeletedProjectID returns the unique identifier of the project that got deleted
// ctx: Mandatory. Reference to the context
// Returns the unique identifier of the the project that got deleted
func (r *deleteProjectPayloadResolver) DeletedProjectID(ctx context.Context) graphql.ID {
	return graphql.ID(r.projectID)
}

// ClientMutationId returns the client mutation ID that was provided as part of the mutation request
// ctx: Mandatory. Reference to the context
// Returns the provided clientMutationId as part of mutation request
func (r *deleteProjectPayloadResolver) ClientMutationId(ctx context.Context) *string {
	return r.clientMutationId
}
