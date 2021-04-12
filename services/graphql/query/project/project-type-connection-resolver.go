// Package project implements different project GraphQL query resovlers required by the GraphQL transport layer
package project

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/project"
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/relay"
	projectGrpcContract "github.com/decentralized-cloud/project/contract/grpc/go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"github.com/thoas/go-funk"
)

type projectTypeConnectionResolver struct {
	resolverCreator types.ResolverCreatorContract
	projects        []*projectGrpcContract.ProjectWithCursor
	hasPreviousPage bool
	hasNextPage     bool
	totalCount      int32
}

// NewProjectTypeConnectionResolver creates new instance of the projectTypeConnectionResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// resolverCreator: Mandatory. Reference to the resolver creator service that can create new instances of resolvers
// projects: Mandatory. Reference the list of projects
// hasPreviousPage: Mandatory. Indicates whether more edges exist prior to the set defined by the clients arguments
// hasNextPage: Mandatory. Indicates whether more edges exist following the set defined by the clients arguments
// totalCount: Mandatory. The total count of matched projects
// Returns the new instance or error if something goes wrong
func NewProjectTypeConnectionResolver(
	ctx context.Context,
	resolverCreator types.ResolverCreatorContract,
	projects []*projectGrpcContract.ProjectWithCursor,
	hasPreviousPage bool,
	hasNextPage bool,
	totalCount int32) (project.ProjectTypeConnectionResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	return &projectTypeConnectionResolver{
		resolverCreator: resolverCreator,
		projects:        projects,
		hasPreviousPage: hasPreviousPage,
		hasNextPage:     hasNextPage,
		totalCount:      totalCount,
	}, nil
}

// PageInfo returns the paging information compatible with graphql-relay
// ctx: Mandatory. Reference to the context
// Returns the paging information
func (r *projectTypeConnectionResolver) PageInfo(ctx context.Context) (relay.PageInfoResolverContract, error) {
	var startCursor, endCursor string

	if len(r.projects) > 0 {
		startCursor = r.projects[0].Cursor
		endCursor = r.projects[len(r.projects)-1].Cursor
	}

	return r.resolverCreator.NewPageInfoResolver(
		ctx,
		&startCursor,
		&endCursor,
		r.hasNextPage,
		r.hasPreviousPage)
}

// Edges returns the project edges compatible with graphql-relay
// ctx: Mandatory. Reference to the context
// Returns the project edges
func (r *projectTypeConnectionResolver) Edges(ctx context.Context) (*[]project.ProjectTypeEdgeResolverContract, error) {
	projects := funk.Filter(r.projects, func(project *projectGrpcContract.ProjectWithCursor) bool {
		return project != nil
	}).([]*projectGrpcContract.ProjectWithCursor)

	edges := []project.ProjectTypeEdgeResolverContract{}
	for _, item := range projects {
		if edge, err := r.resolverCreator.NewProjectTypeEdgeResolver(
			ctx,
			item.ProjectID,
			item.Cursor,
			&project.ProjectDetail{
				Project: item.Project,
			}); err != nil {
			return nil, err
		} else {
			edges = append(edges, edge)
		}
	}

	return &edges, nil
}

// TotalCount returns total count of the matched projects
// ctx: Mandatory. Reference to the context
// Returns the total count of the matched projects
func (r *projectTypeConnectionResolver) TotalCount(ctx context.Context) *int32 {
	return &r.totalCount
}
