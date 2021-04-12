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

type serviceConditionResolverContract struct {
	logger    *zap.Logger
	condition *edgeclusterGrpcContract.ServiceCondition
}

// NewServiceConditionResolver creates new instance of the serviceConditionResolverContract, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// logger: Mandatory. Reference to the logger service
// condition: Mandatory. Contains condition information for a service.
// Returns the new instance or error if something goes wrong
func NewServiceConditionResolver(
	ctx context.Context,
	logger *zap.Logger,
	condition *edgeclusterGrpcContract.ServiceCondition) (edgecluster.ServiceConditionResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if condition == nil {
		return nil, commonErrors.NewArgumentNilError("condition", "condition is required")
	}

	return &serviceConditionResolverContract{
		logger:    logger,
		condition: condition,
	}, nil
}

// Type is the type of the condition
// ctx: Mandatory. Reference to the context
// Returns the type of the condition
func (r *serviceConditionResolverContract) Type(ctx context.Context) string {
	return edgeclusterGrpcContract.ConditionStatus_name[int32(r.condition.Status)]
}

// Status returns the status of the condition
// ctx: Mandatory. Reference to the context
// Returns the status of the condition
func (r *serviceConditionResolverContract) Status(ctx context.Context) string {
	return edgeclusterGrpcContract.ConditionStatus_name[int32(r.condition.Status)]
}

// LastTransitionTime returns the last time the condition transitioned from one status to another
// ctx: Mandatory. Reference to the context
// Returns the last time the condition transitioned from one status to another
func (r *serviceConditionResolverContract) LastTransitionTime(ctx context.Context) string {
	return r.condition.LastTransitionTime.AsTime().Format(time.RFC3339)
}

// Reason returns the Unique, one-word, CamelCase reason for the condition's last transition
// ctx: Mandatory. Reference to the context
// Returns the Unique, one-word, CamelCase reason for the condition's last transition
func (r *serviceConditionResolverContract) Reason(ctx context.Context) string {
	return r.condition.Reason
}

// Message returns the human-readable message indicating details about last transition
// ctx: Mandatory. Reference to the context
// Returns the human-readable message indicating details about last transition
func (r *serviceConditionResolverContract) Message(ctx context.Context) string {
	return r.condition.Message
}
