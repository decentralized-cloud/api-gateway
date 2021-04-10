// Package edgecluster implements different edge cluster GraphQL query resovlers required by the GraphQL transport layer
package edgecluster

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/graphql/types"
	"github.com/decentralized-cloud/api-gateway/services/graphql/types/edgecluster"
	edgeclusterGrpcContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type edgeClusterServiceResolver struct {
	logger          *zap.Logger
	resolverCreator types.ResolverCreatorContract
	service         *edgeclusterGrpcContract.EdgeClusterService
}

// NewEdgeClusterServiceResolver creates new instance of the edgeClusterServiceResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// logger: Mandatory. Reference to the logger service
// service: Mandatory. Contains information about the edge cluster service.
// Returns the new instance or error if something goes wrong
func NewEdgeClusterServiceResolver(
	ctx context.Context,
	logger *zap.Logger,
	resolverCreator types.ResolverCreatorContract,
	service *edgeclusterGrpcContract.EdgeClusterService) (edgecluster.ServiceResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if service == nil {
		return nil, commonErrors.NewArgumentNilError("service", "service is required")
	}

	return &edgeClusterServiceResolver{
		logger:          logger,
		resolverCreator: resolverCreator,
		service:         service,
	}, nil
}

// Metadata contains the service metadata
// ctx: Mandatory. Reference to the context
// Returns the service metadata resolver or error if something goes wrong.
func (r *edgeClusterServiceResolver) Metadata(ctx context.Context) (edgecluster.ObjectMetaResolverContract, error) {
	return r.resolverCreator.NewObjectMetaResolver(ctx, r.service.Metadata)
}

// Status contains the most recently observed status of the service
// ctx: Mandatory. Reference to the context
// Returns the most recently observed status of the service resolver or error if something goes wrong.
func (r *edgeClusterServiceResolver) Status(ctx context.Context) (edgecluster.ServiceStatusResolverContract, error) {
	return r.resolverCreator.NewServiceStatusResolver(ctx, r.service.Status)
}

// Status contains the specification of the desired behavior of the service
// ctx: Mandatory. Reference to the context
// Returns the specification of the desired behavior of the service resolver or error if something goes wrong.
func (r *edgeClusterServiceResolver) Spec(ctx context.Context) (edgecluster.ServiceSpecResolverContract, error) {
	return r.resolverCreator.NewServiceSpecResolver(ctx, r.service.Spec)
}
