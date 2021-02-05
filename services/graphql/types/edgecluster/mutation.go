// packae edgecluster implements used edge cluster related types in the GraphQL transport layer
package edgecluster

import (
	"context"

	"github.com/graph-gophers/graphql-go"
)

type MutationResolverCreatorContract interface {
	// NewCreateEdgeCluster creates new instance of the CreateEdgeClusterContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// Returns the new instance or error if something goes wrong
	NewCreateEdgeCluster(ctx context.Context) (CreateEdgeClusterContract, error)

	// NewCreateEdgeClusterPayloadResolver creates new instance of the CreateEdgeClusterPayloadResolverContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// clientMutationId: Optional. Reference to the client mutation ID to correlate the request and response
	// edgeClusterID: Mandatory. The edge cluster unique identifier
	// edgeClusterDetail: Mandatory. The edge cluster details
	// cursor: Mandatory. The edge cluster cursor
	// Returns the new instance or error if something goes wrong
	NewCreateEdgeClusterPayloadResolver(
		ctx context.Context,
		clientMutationId *string,
		edgeClusterID string,
		edgeClusterDetail *EdgeClusterDetail,
		cursor string) (CreateEdgeClusterPayloadResolverContract, error)

	// NewUpdateEdgeCluster creates new instance of the UpdateEdgeClusterContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// Returns the new instance or error if something goes wrong
	NewUpdateEdgeCluster(ctx context.Context) (UpdateEdgeClusterContract, error)

	// NewUpdateEdgeClusterPayloadResolver creates new instance of the UpdateEdgeClusterPayloadResolverContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// clientMutationId: Optional. Reference to the client mutation ID to correlate the request and response
	// edgeClusterID: Mandatory. The edge cluster unique identifier
	// edgeClusterDetail: Mandatory. The edge cluster details
	// cursor: Mandatory. The edge cluster cursor
	// Returns the new instance or error if something goes wrong
	NewUpdateEdgeClusterPayloadResolver(
		ctx context.Context,
		clientMutationId *string,
		edgeClusterID string,
		edgeClusterDetail *EdgeClusterDetail,
		cursor string) (UpdateEdgeClusterPayloadResolverContract, error)

	// NewDeleteEdgeCluster creates new instance of the DeleteEdgeClusterContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// Returns the new instance or error if something goes wrong
	NewDeleteEdgeCluster(ctx context.Context) (DeleteEdgeClusterContract, error)

	// NewDeleteEdgeClusterPayloadResolver creates new instance of the DeleteEdgeClusterPayloadResolverContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// edgeClusterID: Mandatory. The edge cluster unique identifier
	// clientMutationId: Optional. Reference to the client mutation ID to correlate the request and response
	// Returns the new instance or error if something goes wrong
	NewDeleteEdgeClusterPayloadResolver(
		ctx context.Context,
		edgeClusterID string,
		clientMutationId *string) (DeleteEdgeClusterPayloadResolverContract, error)
}

// RootResolverContract declares the root resolver
type RootResolverContract interface {
	// CreateEdgeCluster returns create edge cluster mutator
	// ctx: Mandatory. Reference to the context
	// Returns the create edge cluster mutator or error if something goes wrong
	CreateEdgeCluster(
		ctx context.Context,
		args CreateEdgeClusterInputArgument) (CreateEdgeClusterPayloadResolverContract, error)

	// UpdateEdgeCluster returns update edge cluster mutator
	// ctx: Mandatory. Reference to the context
	// Returns the update edge cluster mutator or error if something goes wrong
	UpdateEdgeCluster(
		ctx context.Context,
		args UpdateEdgeClusterInputArgument) (UpdateEdgeClusterPayloadResolverContract, error)

	// DeleteEdgeCluster returns delete edge cluster mutator
	// ctx: Mandatory. Reference to the context
	// Returns the delete edge cluster mutator or error if something goes wrong
	DeleteEdgeCluster(
		ctx context.Context,
		args DeleteEdgeClusterInputArgument) (DeleteEdgeClusterPayloadResolverContract, error)
}

// CreateEdgeClusterPayloadResolverContract declares the resolver that can return the payload contains the result of creating a new edge cluster
type CreateEdgeClusterPayloadResolverContract interface {
	// EdgeCluster returns the new edge cluster inforamtion
	// ctx: Mandatory. Reference to the context
	// Returns the new edge cluster inforamtion
	EdgeCluster(ctx context.Context) (EdgeClusterTypeEdgeResolverContract, error)

	// ClientMutationId returns the client mutation ID that was provided as part of the mutation request
	// ctx: Mandatory. Reference to the context
	// Returns the provided clientMutationId as part of mutation request
	ClientMutationId(ctx context.Context) *string
}

// UpdateEdgeClusterPayloadResolverContract declares the resolver that can return the payload contains the result of updating an existing edge cluster
type UpdateEdgeClusterPayloadResolverContract interface {
	// EdgeCluster returns the updated edge cluster inforamtion
	// ctx: Mandatory. Reference to the context
	// Returns the updated edge cluster inforamtion
	EdgeCluster(ctx context.Context) (EdgeClusterTypeEdgeResolverContract, error)

	// ClientMutationId returns the client mutation ID that was provided as part of the mutation request
	// ctx: Mandatory. Reference to the context
	// Returns the provided clientMutationId as part of mutation request
	ClientMutationId(ctx context.Context) *string
}

// DeleteEdgeClusterPayloadResolverContract declares the resolver that can return the payload contains the result of deleting an existing edge cluster
type DeleteEdgeClusterPayloadResolverContract interface {
	// DeletedEdgeClusterID returns the unique identifier of the edge cluster that got deleted
	// ctx: Mandatory. Reference to the context
	// Returns the unique identifier of the the edge cluster that got deleted
	DeletedEdgeClusterID(ctx context.Context) graphql.ID

	// ClientMutationId returns the client mutation ID that was provided as part of the mutation request
	// ctx: Mandatory. Reference to the context
	// Returns the provided clientMutationId as part of mutation request
	ClientMutationId(ctx context.Context) *string
}

// CreateEdgeClusterContract declares the type to use when creating a new edge cluster
type CreateEdgeClusterContract interface {
	// MutateAndGetPayload creates a new edge cluster and returns the payload contains the result of creating a new edge cluster
	// ctx: Mandatory. Reference to the context
	// args: Mandatory. Reference to the input argument contains edge cluster information to create
	// Returns the new edge cluster payload or error if something goes wrong
	MutateAndGetPayload(
		ctx context.Context,
		args CreateEdgeClusterInputArgument) (CreateEdgeClusterPayloadResolverContract, error)
}

// UpdateEdgeClusterContract declares the type to use when updating an existing edge cluster
type UpdateEdgeClusterContract interface {
	// MutateAndGetPayload update an existing edge cluster and returns the payload contains the result of updating an existing edge cluster
	// ctx: Mandatory. Reference to the context
	// args: Mandatory. Reference to the input argument contains edge cluster information to update
	// Returns the updated edge cluster payload or error if something goes wrong
	MutateAndGetPayload(
		ctx context.Context,
		args UpdateEdgeClusterInputArgument) (UpdateEdgeClusterPayloadResolverContract, error)
}

// DeleteEdgeClusterContract declares the type to use when updating an existing edge cluster
type DeleteEdgeClusterContract interface {
	// MutateAndGetPayload update an existing edge cluster and returns the payload contains the result of deleting an existing edge cluster
	// ctx: Mandatory. Reference to the context
	// args: Mandatory. Reference to the input argument contains edge cluster information to update
	// Returns the deleted edge cluster payload or error if something goes wrong
	MutateAndGetPayload(
		ctx context.Context,
		args DeleteEdgeClusterInputArgument) (DeleteEdgeClusterPayloadResolverContract, error)
}

type CreateEdgeClusterInput struct {
	TenantID         graphql.ID
	Name             string
	ClusterSecret    string
	ClusterType      string
	ClientMutationId *string
}

type CreateEdgeClusterInputArgument struct {
	Input CreateEdgeClusterInput
}

type UpdateEdgeClusterInput struct {
	ID               graphql.ID
	TenantID         graphql.ID
	Name             string
	ClusterSecret    string
	ClusterType      string
	ClientMutationId *string
}

type UpdateEdgeClusterInputArgument struct {
	Input UpdateEdgeClusterInput
}

type DeleteEdgeClusterInput struct {
	ID               graphql.ID
	ClientMutationId *string
}

type DeleteEdgeClusterInputArgument struct {
	Input DeleteEdgeClusterInput
}
