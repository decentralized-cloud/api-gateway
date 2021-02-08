// Package project implements different project GraphQL query resovlers required by the GraphQL transport layer
package project

import (
	"context"
	"strings"

	"github.com/decentralized-cloud/api-gateway/services/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/project"
	commonErrors "github.com/micro-business/go-core/system/errors"
)

type projectTypeEdgeResolver struct {
	resolverCreator types.ResolverCreatorContract
	projectID       string
	projectDetail   *project.ProjectDetail
	cursor          string
}

// NewProjectTypeEdgeResolver creates new instance of the projectTypeEdgeResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can create new instances of resolvers
// projectID: Mandatory. the project unique identifier
// projectDetail: Optional. The project details, if provided, the value be used instead of contacting  the edge cluster service
// cursor: Mandatory. the cursor
// Returns the new instance or error if something goes wrong
func NewProjectTypeEdgeResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	projectID string,
	projectDetail *project.ProjectDetail,
	cursor string) (project.ProjectTypeEdgeResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if strings.Trim(projectID, " ") == "" {
		return nil, commonErrors.NewArgumentError("projectID", "projectID is required")
	}

	if strings.Trim(cursor, " ") == "" {
		return nil, commonErrors.NewArgumentError("cursor", "cursor is required")
	}

	return &projectTypeEdgeResolver{
		resolverCreator: resolverCreator,
		projectID:       projectID,
		projectDetail:   projectDetail,
		cursor:          cursor,
	}, nil
}

// Node returns the project resolver
// ctx: Mandatory. Reference to the context
// Returns the project resolver or error if something goes wrong
func (r *projectTypeEdgeResolver) Node(ctx context.Context) (project.ProjectResolverContract, error) {
	return r.resolverCreator.NewProjectResolver(
		ctx,
		r.projectID,
		r.projectDetail)
}

// Cursor returns the cursor for the project edge compatible with graphql-relay
// ctx: Mandatory. Reference to the context
// Returns the cursor
func (r *projectTypeEdgeResolver) Cursor(ctx context.Context) string {
	return r.cursor
}
