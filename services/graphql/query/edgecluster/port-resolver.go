// Package edgecluster implements different edge cluster GraphQL query resovlers required by the GraphQL transport layer
package edgecluster

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/graphql/types/edgecluster"
	edgeclusterGrpcContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type portResolver struct {
	logger *zap.Logger
	port   *edgeclusterGrpcContract.Port
}

// NewPortResolver creates new instance of the PortResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// logger: Mandatory. Reference to the logger service
// Port: Mandatory. The PostStatus details
// Returns the new instance or error if something goes wrong
func NewPortResolver(
	ctx context.Context,
	logger *zap.Logger,
	port *edgeclusterGrpcContract.Port) (edgecluster.PortResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if port == nil {
		return nil, commonErrors.NewArgumentNilError("port", "port is required")
	}

	return &portResolver{
		logger: logger,
		port:   port,
	}, nil
}

// Port returns the port number of the edge-cluster master port of which status is recorded here
// ctx: Mandatory. Reference to the context
// Returns the port number of the edge-cluster master port of which status is recorded here
func (r *portResolver) Port(ctx context.Context) int32 {
	return r.port.Port
}

// Protocol returns the protocol of the edge-cluster master port of which status is recorded here
// ctx: Mandatory. Reference to the context
// Returns the protocol of the edge-cluster master port of which status is recorded here
func (r *portResolver) Protocol(ctx context.Context) string {
	return r.port.Protcol.String()
}
