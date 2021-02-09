// Package edgecluster implements different edge cluster GraphQL query resovlers required by the GraphQL transport layer
package edgecluster

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/graphql/types/edgecluster"
	edgeclusterGrpcContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type ingressResolver struct {
	logger  *zap.Logger
	ingress *edgeclusterGrpcContract.Ingress
}

// NewIngressResolver creates new instance of the ingressResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// logger: Mandatory. Reference to the logger service
// ingress: Mandatory. The ingress details
// Returns the new instance or error if something goes wrong
func NewIngressResolver(
	ctx context.Context,
	logger *zap.Logger,
	ingress *edgeclusterGrpcContract.Ingress) (edgecluster.IngressResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if ingress == nil {
		return nil, commonErrors.NewArgumentNilError("ingress", "ingress is required")
	}

	return &ingressResolver{
		logger:  logger,
		ingress: ingress,
	}, nil
}

// IP is set for load-balancer ingress points that are IP based
// (typically GCE or OpenStack load-balancers)
// ctx: Mandatory. Reference to the context
// Returns the IP is set for load-balancer ingress points that are IP based
func (r *ingressResolver) IP(ctx context.Context) *string {
	return &r.ingress.Ip
}

// Hostname is set for load-balancer ingress points that are DNS based
// (typically AWS load-balancers)
// ctx: Mandatory. Reference to the context
// Returns the Hostname is set for load-balancer ingress points that are DNS based
func (r *ingressResolver) Hostname(ctx context.Context) *string {
	return &r.ingress.Hostname
}