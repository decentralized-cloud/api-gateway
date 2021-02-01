// Package edgecluster implements different edge cluster GraphQL query resovlers required by the GraphQL transport layer
package edgecluster

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/graphql/types/edgecluster"
	edgeclusterGrpcContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type edgeClusterNodeAddressResolverContract struct {
	logger       *zap.Logger
	nodeAddresss *edgeclusterGrpcContract.EdgeClusterNodeAddress
}

// NewEdgeClusterNodeAddressResolverContract creates new instance of the edgeClusterNodeAddressResolverContract, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// logger: Mandatory. Reference to the logger service
// nodeAddresss: Mandatory. Contains information for the edge cluster node's address.
// Returns the new instance or error if something goes wrong
func NewEdgeClusterNodeAddressResolverContract(
	ctx context.Context,
	logger *zap.Logger,
	nodeAddresss *edgeclusterGrpcContract.EdgeClusterNodeAddress) (edgecluster.EdgeClusterNodeAddressResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if nodeAddresss == nil {
		return nil, commonErrors.NewArgumentNilError("nodeAddresss", "nodeAddresss is required")
	}

	return &edgeClusterNodeAddressResolverContract{
		logger:       logger,
		nodeAddresss: nodeAddresss,
	}, nil
}

// Type returns the edge cluster node address type, one of Hostname, ExternalIP or InternalIP.
// ctx: Mandatory. Reference to the context
// Returns the edge cluster node address type, one of Hostname, ExternalIP or InternalIP.
func (r *edgeClusterNodeAddressResolverContract) NodeAddressType(ctx context.Context) string {
	return edgeclusterGrpcContract.EdgeClusterNodeAddressType_name[int32(r.nodeAddresss.NodeAddressType)]
}

// Address returns the node address.
// ctx: Mandatory. Reference to the context
// Returns the node address.
func (r *edgeClusterNodeAddressResolverContract) Address(ctx context.Context) string {
	return r.nodeAddresss.Address
}
