// Package edgecluster implements different edge cluster GraphQL query resovlers required by the GraphQL transport layer
package edgecluster

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/graphql/types/edgecluster"
	edgeclusterGrpcContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type portStatusResolver struct {
	logger     *zap.Logger
	PortStatus *edgeclusterGrpcContract.PortStatus
}

// NewPortStatusResolver creates new instance of the portStatusResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// logger: Mandatory. Reference to the logger service
// portStatus: Mandatory. Contains the port status.
// Returns the new instance or error if something goes wrong
func NewPortStatusResolver(
	ctx context.Context,
	logger *zap.Logger,
	PortStatus *edgeclusterGrpcContract.PortStatus) (edgecluster.PortStatusResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if PortStatus == nil {
		return nil, commonErrors.NewArgumentNilError("PortStatus", "PortStatus is required")
	}

	return &portStatusResolver{
		logger:     logger,
		PortStatus: PortStatus,
	}, nil
}

// Port is the port number of the service port of which status is recorded here
// ctx: Mandatory. Reference to the context
// Returns the port number of the service port of which status is recorded here
func (r *portStatusResolver) Port(ctx context.Context) int32 {
	return r.PortStatus.Port
}

// Protocol is the protocol of the service port of which status is recorded here
// ctx: Mandatory. Reference to the context
// Returns the protocol of the service port of which status is recorded here
func (r *portStatusResolver) Protocol(ctx context.Context) string {
	return r.PortStatus.Protocol.String()

}

// Error is to record the problem with the service port
// ctx: Mandatory. Reference to the context
// Returns the problem with the service port
func (r *portStatusResolver) Error(ctx context.Context) *string {
	if r.PortStatus.Error == "" {
		return nil
	}

	return &r.PortStatus.Error
}
