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

type edgeClusterPodResolver struct {
	logger          *zap.Logger
	resolverCreator types.ResolverCreatorContract
	pod             *edgeclusterGrpcContract.EdgeClusterPod
}

// NewEdgeClusterPodResolver creates new instance of the edgeClusterPodResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// logger: Mandatory. Reference to the logger service
// pod: Mandatory. Contains information about the edge cluster pod.
// Returns the new instance or error if something goes wrong
func NewEdgeClusterPodResolver(
	ctx context.Context,
	logger *zap.Logger,
	resolverCreator types.ResolverCreatorContract,
	pod *edgeclusterGrpcContract.EdgeClusterPod) (edgecluster.PodResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	if pod == nil {
		return nil, commonErrors.NewArgumentNilError("pod", "pod is required")
	}

	return &edgeClusterPodResolver{
		logger:          logger,
		resolverCreator: resolverCreator,
		pod:             pod,
	}, nil
}

// Metadata contains the pod metadata
// ctx: Mandatory. Reference to the context
// Returns the pod metadata resolver or error if something goes wrong.
func (r *edgeClusterPodResolver) Metadata(ctx context.Context) (edgecluster.ObjectMetaResolverContract, error) {
	return r.resolverCreator.NewObjectMetaResolver(ctx, r.pod.Metadata)
}

// Status contains the most recently observed status of the pod
// ctx: Mandatory. Reference to the context
// Returns the most recently observed status of the pod resolver or error if something goes wrong.
func (r *edgeClusterPodResolver) Status(ctx context.Context) (edgecluster.PodStatusResolverContract, error) {
	return r.resolverCreator.NewPodStatusResolver(ctx, r.pod.Status)
}

// Status contains the specification of the desired behavior of the pod
// ctx: Mandatory. Reference to the context
// Returns the specification of the desired behavior of the pod resolver or error if something goes wrong.
func (r *edgeClusterPodResolver) Spec(ctx context.Context) (edgecluster.PodSpecResolverContract, error) {
	return r.resolverCreator.NewPodSpecResolver(ctx, r.pod.Spec)
}
