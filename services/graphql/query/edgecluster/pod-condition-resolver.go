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

type podConditionResolver struct {
	logger       *zap.Logger
	podCondition *edgeclusterGrpcContract.PodCondition
}

// NewPodConditionResolver creates new instance of the podConditionResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// logger: Mandatory. Reference to the logger service
// podCondition: Mandatory. Contains condition information for a pod.
// Returns the new instance or error if something goes wrong
func NewPodConditionResolver(
	ctx context.Context,
	logger *zap.Logger,
	podCondition *edgeclusterGrpcContract.PodCondition) (edgecluster.PodConditionResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if podCondition == nil {
		return nil, commonErrors.NewArgumentNilError("podCondition", "podCondition is required")
	}

	return &podConditionResolver{
		logger:       logger,
		podCondition: podCondition,
	}, nil
}

// Type is the type of the condition
// ctx: Mandatory. Reference to the context
// Returns the type of the condition
func (r *podConditionResolver) Type(ctx context.Context) string {
	return edgeclusterGrpcContract.PodConditionType_name[int32(r.podCondition.Type)]
}

// Status returns the status of the condition
// ctx: Mandatory. Reference to the context
// Returns the status of the condition
func (r *podConditionResolver) Status(ctx context.Context) string {
	return edgeclusterGrpcContract.ConditionStatus_name[int32(r.podCondition.Status)]
}

// LastHeartbeatTime returns the last time we got an update on a given condition
// ctx: Mandatory. Reference to the context
// Returns the last time we got an update on a given condition
func (r *podConditionResolver) LastProbeTime(ctx context.Context) string {
	return r.podCondition.LastProbeTime.AsTime().Format(time.RFC3339)
}

// LastTransitionTime returns the last time the condition transitioned from one status to another
// ctx: Mandatory. Reference to the context
// Returns the last time the condition transitioned from one status to another
func (r *podConditionResolver) LastTransitionTime(ctx context.Context) string {
	return r.podCondition.LastTransitionTime.AsTime().Format(time.RFC3339)
}

// Reason returns the Unique, one-word, CamelCase reason for the condition's last transition
// ctx: Mandatory. Reference to the context
// Returns the Unique, one-word, CamelCase reason for the condition's last transition
func (r *podConditionResolver) Reason(ctx context.Context) string {
	return r.podCondition.Reason
}

// Message returns the human-readable message indicating details about last transition
// ctx: Mandatory. Reference to the context
// Returns the human-readable message indicating details about last transition
func (r *podConditionResolver) Message(ctx context.Context) string {
	return r.podCondition.Message
}
