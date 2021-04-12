// Package edgecluster implements different edge cluster GraphQL query resovlers required by the GraphQL transport layer
package edgecluster

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/graphql/types/edgecluster"
	edgeclusterGrpcContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type servicePortResolver struct {
	logger      *zap.Logger
	servicePort *edgeclusterGrpcContract.ServicePort
}

// NewServicePortResolver creates new instance of the servicePortResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// logger: Mandatory. Reference to the logger service
// servicePort: Mandatory. Contains condition information for a pod.
// Returns the new instance or error if something goes wrong
func NewServicePortResolver(
	ctx context.Context,
	logger *zap.Logger,
	servicePort *edgeclusterGrpcContract.ServicePort) (edgecluster.ServicePortResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if servicePort == nil {
		return nil, commonErrors.NewArgumentNilError("servicePort", "servicePort is required")
	}

	return &servicePortResolver{
		logger:      logger,
		servicePort: servicePort,
	}, nil
}

// Type is the name of this port within the service
// ctx: Mandatory. Reference to the context
// Returns he name of this port within the service
func (r *servicePortResolver) Name(ctx context.Context) string {
	return r.servicePort.Name
}

// Protocol is the IP protocol for this port
// ctx: Mandatory. Reference to the context
// Returns the IP protocol for this port
func (r *servicePortResolver) Protocol(ctx context.Context) string {
	return r.servicePort.Protcol.String()
}

// Port is the port that will be exposed by this service
// ctx: Mandatory. Reference to the context
// Returns the port that will be exposed by this service
func (r *servicePortResolver) Port(ctx context.Context) int32 {
	return r.servicePort.Port
}

// TargetPort is the number or name of the port to access on the pods targeted by the service
// ctx: Mandatory. Reference to the context
// Returns the number or name of the port to access on the pods targeted by the service
func (r *servicePortResolver) TargetPort(ctx context.Context) string {
	return r.servicePort.TargetPort
}

// NodePort is the port on each node on which this service is exposed when type is
// NodePort or LoadBalancer
// ctx: Mandatory. Reference to the context
// Returns the port on each node on which this service is exposed
func (r *servicePortResolver) NodePort(ctx context.Context) int32 {
	return r.servicePort.NodePort
}
