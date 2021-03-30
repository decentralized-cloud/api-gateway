// Package edgecluster implements different edge cluster GraphQL query resovlers required by the GraphQL transport layer
package edgecluster

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/graphql/types/edgecluster"
	edgeclusterGrpcContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
	"github.com/graph-gophers/graphql-go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type edgeClusterObjectMetadataResolver struct {
	logger   *zap.Logger
	metadata *edgeclusterGrpcContract.EdgeClusterObjectMetadata
}

// NewEdgeClusterObjectMetadataResolver creates new instance of the edgeClusterObjectMetadataResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// logger: Mandatory. Reference to the logger service
// metadata: Mandatory. Contains the object metadata.
// Returns the new instance or error if something goes wrong
func NewEdgeClusterObjectMetadataResolver(
	ctx context.Context,
	logger *zap.Logger,
	metadata *edgeclusterGrpcContract.EdgeClusterObjectMetadata) (edgecluster.EdgeClusterObjectMetadataResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if metadata == nil {
		return nil, commonErrors.NewArgumentNilError("metadata", "metadata is required")
	}

	return &edgeClusterObjectMetadataResolver{
		logger:   logger,
		metadata: metadata,
	}, nil
}

// ID returns the object unique identitfier
// ctx: Mandatory. Reference to the context
// Returns the object unique identitfier
func (r *edgeClusterObjectMetadataResolver) ID(ctx context.Context) graphql.ID {
	return graphql.ID(r.metadata.Id)
}

// Name returns the object name
// ctx: Mandatory. Reference to the context
// Returns the object name
func (r *edgeClusterObjectMetadataResolver) Name(ctx context.Context) string {
	return r.metadata.Name
}

// Namespace returns the object namespace
// ctx: Mandatory. Reference to the context
// Returns the object namespace
func (r *edgeClusterObjectMetadataResolver) Namespace(ctx context.Context) *string {
	return &r.metadata.Namespace
}
