// packae project implements used project related types in the GraphQL transport layer
package project

import (
	"context"

	"github.com/graph-gophers/graphql-go"
)

type MutationResolverCreatorContract interface {
	// NewCreateProject creates new instance of the CreateProjectContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// Returns the new instance or error if something goes wrong
	NewCreateProject(ctx context.Context) (CreateProjectContract, error)

	// NewCreateProjectPayloadResolver creates new instance of the CreateProjectPayloadResolverContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// clientMutationId: Optional. Reference to the client mutation ID to correlate the request and response
	// projectID: Mandatory. The project unique identifier
	// projectDetail: Mandatory. The project details
	// cursor: Mandatory. The edge cluster cursor
	// Returns the new instance or error if something goes wrong
	NewCreateProjectPayloadResolver(
		ctx context.Context,
		clientMutationId *string,
		projectID string,
		projectDetail *ProjectDetail,
		cursor string) (CreateProjectPayloadResolverContract, error)

	// NewUpdateProject creates new instance of the UpdateProjectContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// clientMutationId: Optional. Reference to the client mutation ID to correlate the request and response
	// Returns the new instance or error if something goes wrong
	NewUpdateProject(ctx context.Context) (UpdateProjectContract, error)

	// NewUpdateProjectPayloadResolver creates new instance of the UpdateProjectPayloadResolverContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// clientMutationId: Optional. Reference to the client mutation ID to correlate the request and response
	// projectID: Mandatory. The project unique identifier
	// projectDetail: Mandatory. The project details
	// cursor: Mandatory. The edge cluster cursor
	// Returns the new instance or error if something goes wrong
	NewUpdateProjectPayloadResolver(
		ctx context.Context,
		clientMutationId *string,
		projectID string,
		projectDetail *ProjectDetail,
		cursor string) (UpdateProjectPayloadResolverContract, error)

	// NewDeleteProject creates new instance of the DeleteProjectContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// clientMutationId: Optional. Reference to the client mutation ID to correlate the request and response
	// Returns the new instance or error if something goes wrong
	NewDeleteProject(ctx context.Context) (DeleteProjectContract, error)

	// NewDeleteProjectPayloadResolver creates new instance of the DeleteProjectPayloadResolverContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// projectID: Mandatory. The project unique identifier
	// clientMutationId: Optional. Reference to the client mutation ID to correlate the request and response
	// Returns the new instance or error if something goes wrong
	NewDeleteProjectPayloadResolver(
		ctx context.Context,
		projectID string,
		clientMutationId *string) (DeleteProjectPayloadResolverContract, error)
}

// RootResolverContract declares the root resolver
type RootResolverContract interface {
	// CreateProject returns create project mutator
	// ctx: Mandatory. Reference to the context
	// Returns the create project mutator or error if something goes wrong
	CreateProject(
		ctx context.Context,
		args CreateProjectInputArgument) (CreateProjectPayloadResolverContract, error)

	// UpdateProject returns update project mutator
	// ctx: Mandatory. Reference to the context
	// Returns the update project mutator or error if something goes wrong
	UpdateProject(
		ctx context.Context,
		args UpdateProjectInputArgument) (UpdateProjectPayloadResolverContract, error)

	// DeleteProject returns delete project mutator
	// ctx: Mandatory. Reference to the context
	// Returns the delete project mutator or error if something goes wrong
	DeleteProject(
		ctx context.Context,
		args DeleteProjectInputArgument) (DeleteProjectPayloadResolverContract, error)
}

// CreateProjectPayloadResolverContract declares the resolver that can return the payload contains the result of creating a new project
type CreateProjectPayloadResolverContract interface {
	// Project returns the new project inforamtion
	// ctx: Mandatory. Reference to the context
	// Returns the new project inforamtion
	Project(ctx context.Context) (ProjectTypeEdgeResolverContract, error)

	// ClientMutationId returns the client mutation ID that was provided as part of the mutation request
	// ctx: Mandatory. Reference to the context
	// Returns the provided clientMutationId as part of mutation request
	ClientMutationId(ctx context.Context) *string
}

// UpdateProjectPayloadResolverContract declares the resolver that can return the payload contains the result of updating an existing project
type UpdateProjectPayloadResolverContract interface {
	// Project returns the updated project inforamtion
	// ctx: Mandatory. Reference to the context
	// Returns the updated project inforamtion
	Project(ctx context.Context) (ProjectTypeEdgeResolverContract, error)

	// ClientMutationId returns the client mutation ID that was provided as part of the mutation request
	// ctx: Mandatory. Reference to the context
	// Returns the provided clientMutationId as part of mutation request
	ClientMutationId(ctx context.Context) *string
}

// DeleteProjectPayloadResolverContract declares the resolver that can return the payload contains the result of deleting an existing project
type DeleteProjectPayloadResolverContract interface {
	// DeletedProjectID returns the unique identifier of the project that got deleted
	// ctx: Mandatory. Reference to the context
	// Returns the unique identifier of the the project that got deleted
	DeletedProjectID(ctx context.Context) graphql.ID

	// ClientMutationId returns the client mutation ID that was provided as part of the mutation request
	// ctx: Mandatory. Reference to the context
	// Returns the provided clientMutationId as part of mutation request
	ClientMutationId(ctx context.Context) *string
}

// CreateProjectContract declares the type to use when creating a new project
type CreateProjectContract interface {
	// MutateAndGetPayload creates a new project and returns the payload contains the result of creating a new project
	// ctx: Mandatory. Reference to the context
	// args: Mandatory. Reference to the input argument contains project information to create
	// Returns the new project payload or error if something goes wrong
	MutateAndGetPayload(
		ctx context.Context,
		args CreateProjectInputArgument) (CreateProjectPayloadResolverContract, error)
}

// UpdateProjectContract declares the type to use when updating an existing project
type UpdateProjectContract interface {
	// MutateAndGetPayload update an existing project and returns the payload contains the result of updating an existing project
	// ctx: Mandatory. Reference to the context
	// args: Mandatory. Reference to the input argument contains project information to update
	// Returns the updated project payload or error if something goes wrong
	MutateAndGetPayload(
		ctx context.Context,
		args UpdateProjectInputArgument) (UpdateProjectPayloadResolverContract, error)
}

// DeleteProjectContract declares the type to use when updating an existing project
type DeleteProjectContract interface {
	// MutateAndGetPayload update an existing project and returns the payload contains the result of deleting an existing project
	// ctx: Mandatory. Reference to the context
	// args: Mandatory. Reference to the input argument contains project information to update
	// Returns the deleted project payload or error if something goes wrong
	MutateAndGetPayload(
		ctx context.Context,
		args DeleteProjectInputArgument) (DeleteProjectPayloadResolverContract, error)
}

type CreateProjectInput struct {
	Name             string
	ClientMutationId *string
}

type CreateProjectInputArgument struct {
	Input CreateProjectInput
}

type UpdateProjectInput struct {
	ProjectID        graphql.ID
	Name             string
	ClientMutationId *string
}

type UpdateProjectInputArgument struct {
	Input UpdateProjectInput
}

type DeleteProjectInput struct {
	ProjectID        graphql.ID
	Name             string
	ClientMutationId *string
}

type DeleteProjectInputArgument struct {
	Input DeleteProjectInput
}
