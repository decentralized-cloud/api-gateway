// Package types defines the different interfaces used in the GraphQL implementation
package types

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/graphql/types/edgecluster"
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/project"
	"github.com/graph-gophers/graphql-go"
)

// UserResolverContract declares the resolver that can retrieve user information
type UserResolverContract interface {
	// ID returns user unique identifier
	// ctx: Mandatory. Reference to the context
	// Returns the user unique identifier
	ID(ctx context.Context) graphql.ID

	// Project returns project resolver
	// ctx: Mandatory. Reference to the context
	// args: Mandatory. The argument list
	// Returns the project resolver or error if something goes wrong
	Project(
		ctx context.Context,
		args UserProjectInputArgument) (project.ProjectResolverContract, error)

	// Projects returns project connection compatible with graphql-relay
	// ctx: Mandatory. Reference to the context
	// args: Mandatory. The argument list
	// Returns the project resolver or error if something goes wrong
	Projects(
		ctx context.Context,
		args UserProjectsInputArgument) (project.ProjectTypeConnectionResolverContract, error)

	// EdgeCluster returns project resolver
	// ctx: Mandatory. Reference to the context
	// args: Mandatory. The argument list
	// Returns the project resolver or error if something goes wrong
	EdgeCluster(
		ctx context.Context,
		args UserEdgeClusterInputArgument) (edgecluster.EdgeClusterResolverContract, error)

	// EdgeClusters returns project connection compatible with graphql-relay
	// ctx: Mandatory. Reference to the context
	// args: Mandatory. The argument list
	// Returns the project resolver or error if something goes wrong
	EdgeClusters(
		ctx context.Context,
		args UserEdgeClustersInputArgument) (edgecluster.EdgeClusterTypeConnectionResolverContract, error)
}
