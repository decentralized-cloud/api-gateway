// Package edgecluster implements different edge cluster GraphQL query resovlers required by the GraphQL transport layer
package edgecluster

import (
	"context"
	"time"

	"github.com/decentralized-cloud/api-gateway/services/graphql/types/edgecluster"
	edgeclusterGrpcContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type edgeClusterPodConditionResolver struct {
	logger    *zap.Logger
	condition *edgeclusterGrpcContract.EdgeClusterPodCondition
}

// NewEdgeClusterPodConditionResolver creates new instance of the edgeClusterPodConditionResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// logger: Mandatory. Reference to the logger service
// condition: Mandatory. Contains condition information for a pod.
// Returns the new instance or error if something goes wrong
func NewEdgeClusterPodConditionResolver(
	ctx context.Context,
	logger *zap.Logger,
	condition *edgeclusterGrpcContract.EdgeClusterPodCondition) (edgecluster.EdgeClusterPodConditionResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if condition == nil {
		return nil, commonErrors.NewArgumentNilError("condition", "condition is required")
	}

	return &edgeClusterPodConditionResolver{
		logger:    logger,
		condition: condition,
	}, nil
}

// Type returns the type of pod condition.
// ctx: Mandatory. Reference to the context
// Returns the type of pod condition.
func (r *edgeClusterPodConditionResolver) Type(ctx context.Context) string {
	return edgeclusterGrpcContract.EdgeClusterPodConditionType_name[int32(r.condition.Type)]
}

// Status returns the status of the condition, one of True, False, Unknown.
// ctx: Mandatory. Reference to the context
// Returns the status of the condition, one of True, False, Unknown.
func (r *edgeClusterPodConditionResolver) Status(ctx context.Context) string {
	return edgeclusterGrpcContract.EdgeClusterConditionStatus_name[int32(r.condition.Status)]
}

// LastHeartbeatTime returns the last time we got an update on a given condition.
// ctx: Mandatory. Reference to the context
// Returns the last time we got an update on a given condition.
func (r *edgeClusterPodConditionResolver) LastProbeTime(ctx context.Context) string {
	return r.condition.LastProbeTime.AsTime().Format(time.RFC3339)
}

// LastTransitionTime returns the last time the condition transit from one status to another.
// ctx: Mandatory. Reference to the context
// Returns the last time the condition transit from one status to another.
func (r *edgeClusterPodConditionResolver) LastTransitionTime(ctx context.Context) string {
	return r.condition.LastTransitionTime.AsTime().Format(time.RFC3339)
}

// Reason returns the (brief) reason for the condition's last transition.
// ctx: Mandatory. Reference to the context
// Returns the (brief) reason for the condition's last transition.
func (r *edgeClusterPodConditionResolver) Reason(ctx context.Context) string {
	return r.condition.Reason
}

// Message returns the human readable message indicating details about last transition.
// ctx: Mandatory. Reference to the context
// Returns the human readable message indicating details about last transition.
func (r *edgeClusterPodConditionResolver) Message(ctx context.Context) string {
	return r.condition.Message
}
