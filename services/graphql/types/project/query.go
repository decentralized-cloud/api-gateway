// packae project implements used project related types in the GraphQL transport layer
package project

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/graphql/types/edgecluster"
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/relay"
	projectGrpcContract "github.com/decentralized-cloud/project/contract/grpc/go"
	"github.com/graph-gophers/graphql-go"
)

type QueryResolverCreatorContract interface {
	// NewProjectResolver creates new ProjectResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// projectID: Mandatory. The project unique identifier
	// projectDetail: Optional. The tennat details, if provided, the value be used instead of contacting  the edge cluster service
	// Returns the ProjectResolverContract or error if something goes wrong
	NewProjectResolver(
		ctx context.Context,
		projectID string,
		projectDetail *ProjectDetail) (ProjectResolverContract, error)

	// NewProjectTypeEdgeResolver creates new ProjectTypeEdgeResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// projectID: Mandatory. The project unique identifier
	// projectDetail: Optional. The tennat details, if provided, the value be used instead of contacting  the edge cluster service
	// cursor: Mandatory. The cursor
	// Returns the ProjectTypeEdgeResolverContract or error if something goes wrong
	NewProjectTypeEdgeResolver(
		ctx context.Context,
		projectID string,
		cursor string,
		projectDetail *ProjectDetail) (ProjectTypeEdgeResolverContract, error)

	// NewProjectTypeConnectionResolver creates new ProjectTypeConnectionResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// projects: Mandatory. Reference the list of projects
	// hasPreviousPage: Mandatory. Indicates whether more edges exist prior to the set defined by the clients arguments
	// hasNextPage: Mandatory. Indicates whether more edges exist following the set defined by the clients arguments
	// totalCount: Mandatory. The total count of matched projects
	// Returns the ProjectTypeConnectionResolverContract or error if something goes wrong
	NewProjectTypeConnectionResolver(
		ctx context.Context,
		projects []*projectGrpcContract.ProjectWithCursor,
		hasPreviousPage bool,
		hasNextPage bool,
		totalCount int32) (ProjectTypeConnectionResolverContract, error)
}

// ProjectResolverContract declares the resolver that can retrieve project information
type ProjectResolverContract interface {
	// ID returns project unique identifier
	// ctx: Mandatory. Reference to the context
	// Returns the project unique identifier
	ID(ctx context.Context) graphql.ID

	// Name returns project name
	// ctx: Mandatory. Reference to the context
	// Returns the project name
	Name(ctx context.Context) string

	// EdgeCluster returns project resolver
	// ctx: Mandatory. Reference to the context
	// args: Mandatory. The argument list
	// Returns the project resolver or error if something goes wrong
	EdgeCluster(
		ctx context.Context,
		args ProjectClusterEdgeClusterInputArgument) (edgecluster.EdgeClusterResolverContract, error)

	// EdgeClusters returns project connection compatible with graphql-relay
	// ctx: Mandatory. Reference to the context
	// args: Mandatory. The argument list
	// Returns the project resolver or error if something goes wrong
	EdgeClusters(
		ctx context.Context,
		args ProjectEdgeClustersInputArgument) (edgecluster.EdgeClusterTypeConnectionResolverContract, error)
}

// ProjectTypeConnectionResolverContract declares the resolver that returns project edge compatible with graphql-relay
type ProjectTypeConnectionResolverContract interface {
	// PageInfo returns the paging information compatible with graphql-relay
	// ctx: Mandatory. Reference to the context
	// Returns the paging information
	PageInfo(ctx context.Context) (relay.PageInfoResolverContract, error)

	// Edges returns the project edges compatible with graphql-relay
	// ctx: Mandatory. Reference to the context
	// Returns the project edges
	Edges(ctx context.Context) (*[]ProjectTypeEdgeResolverContract, error)

	// TotalCount returns total count of the matched projects
	// ctx: Mandatory. Reference to the context
	// Returns the total count of the matched projects
	TotalCount(ctx context.Context) *int32
}

// ProjectTypeEdgeResolverContract declares the resolver that returns project edge compatible with graphql-relay
type ProjectTypeEdgeResolverContract interface {
	// Node returns the project resolver
	// ctx: Mandatory. Reference to the context
	// Returns the project resolver or error if something goes wrong
	Node(ctx context.Context) (ProjectResolverContract, error)

	// Cursor returns the cursor for the project edge compatible with graphql-relay
	// ctx: Mandatory. Reference to the context
	// Returns the cursor
	Cursor(ctx context.Context) string
}

type ProjectDetail struct {
	Project *projectGrpcContract.Project
}

type ProjectClusterEdgeClusterInputArgument struct {
	EdgeClusterID graphql.ID
}

type ProjectEdgeClustersInputArgument struct {
	relay.ConnectionArgument
	EdgeClusterIDs *[]graphql.ID
	SortOption     *string
}
