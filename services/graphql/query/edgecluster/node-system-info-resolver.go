// Package edgecluster implements different edge cluster GraphQL query resovlers required by the GraphQL transport layer
package edgecluster

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/graphql/types/edgecluster"
	edgeclusterGrpcContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type nodeSystemInfoResolver struct {
	logger   *zap.Logger
	nodeInfo *edgeclusterGrpcContract.NodeSystemInfo
}

// NewNodeSystemInfoResolver creates new instance of the nodeSystemInfoResolver, setting up all dependencies and returns the instance
// ctx: Mandatory. Reference to the context
// logger: Mandatory. Reference to the logger service
// nodeInfo: Mandatory. Contains a set of ids/uuids to uniquely identify the node.
// Returns the new instance or error if something goes wrong
func NewNodeSystemInfoResolver(
	ctx context.Context,
	logger *zap.Logger,
	nodeInfo *edgeclusterGrpcContract.NodeSystemInfo) (edgecluster.NodeSystemInfoResolverContract, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if nodeInfo == nil {
		return nil, commonErrors.NewArgumentNilError("nodeInfo", "nodeInfo is required")
	}

	return &nodeSystemInfoResolver{
		logger:   logger,
		nodeInfo: nodeInfo,
	}, nil
}

// MachineID reported by the node. For unique machine identification in the cluster this field is preferred.
// ctx: Mandatory. Reference to the context
// Returns the MachineID reported by the node. For unique machine identification in the cluster this field is preferred.
func (r *nodeSystemInfoResolver) MachineID(ctx context.Context) string {
	return r.nodeInfo.MachineID
}

// SystemUUID reported by the node. For unique machine identification MachineID is preferred. This field is specific to Red Hat hosts
// ctx: Mandatory. Reference to the context
// Returns the SystemUUID reported by the node. For unique machine identification MachineID is preferred. This field is specific to Red Hat hosts
func (r *nodeSystemInfoResolver) SystemUUID(ctx context.Context) string {
	return r.nodeInfo.SystemUUID
}

// BootID reported by the node.
// ctx: Mandatory. Reference to the context
// Returns the Boot ID reported by the node.
func (r *nodeSystemInfoResolver) BootID(ctx context.Context) string {
	return r.nodeInfo.BootID
}

// KernelVersion reported by the node from 'uname -r' (e.g. 3.16.0-0.bpo.4-amd64).
// ctx: Mandatory. Reference to the context
// Returns the Kernel Version reported by the node from 'uname -r' (e.g. 3.16.0-0.bpo.4-amd64).
func (r *nodeSystemInfoResolver) KernelVersion(ctx context.Context) string {
	return r.nodeInfo.KernelVersion
}

// OSImage reported by the node from /etc/os-release (e.g. Debian GNU/Linux 7 (wheezy)).
// ctx: Mandatory. Reference to the context
// Returns the OS Image reported by the node from /etc/os-release (e.g. Debian GNU/Linux 7 (wheezy)).
func (r *nodeSystemInfoResolver) OSImage(ctx context.Context) string {
	return r.nodeInfo.OsImage
}

// ContainerRuntimeVersion reported by the node through runtime remote API (e.g. docker://1.5.0).
// ctx: Mandatory. Reference to the context
// Returns the ContainerRuntime Version reported by the node through runtime remote API (e.g. docker://1.5.0).
func (r *nodeSystemInfoResolver) ContainerRuntimeVersion(ctx context.Context) string {
	return r.nodeInfo.ContainerRuntimeVersion
}

// KubeletVersion reported by the node.
// ctx: Mandatory. Reference to the context
// Returns the Kubelet Version reported by the node.
func (r *nodeSystemInfoResolver) KubeletVersion(ctx context.Context) string {
	return r.nodeInfo.KubeletVersion
}

// KubeProxyVersion reported by the node.
// ctx: Mandatory. Reference to the context
// Returns the KubeProxy Version reported by the node.
func (r *nodeSystemInfoResolver) KubeProxyVersion(ctx context.Context) string {
	return r.nodeInfo.KubeProxyVersion
}

// OperatingSystem System reported by the node
// ctx: Mandatory. Reference to the context
// Returns the Operating System reported by the node
func (r *nodeSystemInfoResolver) OperatingSystem(ctx context.Context) string {
	return r.nodeInfo.OperatingSystem
}

// The Architecture reported by the node
// ctx: Mandatory. Reference to the context
// Returns the Architecture reported by the node
func (r *nodeSystemInfoResolver) Architecture(ctx context.Context) string {
	return r.nodeInfo.Architecture
}
