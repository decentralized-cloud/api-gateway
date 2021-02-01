// packae edgecluster implements used edge cluster related types in the GraphQL transport layer
package edgecluster

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/graphql/types/relay"
	edgeclusterGrpcContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
	"github.com/graph-gophers/graphql-go"
)

type QueryResolverCreatorContract interface {
	// NewEdgeClusterResolver creates new EdgeClusterResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// edgeClusterID: Mandatory. The edge cluster unique identifier
	// edgeClusterDetail: Optional. The edge cluster details, if provided, the value be used instead of contacting  the edge cluster service
	// Returns the EdgeClusterResolverContract or error if something goes wrong
	NewEdgeClusterResolver(
		ctx context.Context,
		edgeClusterID string,
		edgeClusterDetail *EdgeClusterDetail) (EdgeClusterResolverContract, error)

	// NewEdgeClusterTypeEdgeResolver creates new EdgeClusterTypeEdgeResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// edgeClusterID: Mandatory. The edge cluster unique identifier
	// cursor: Mandatory. The cursor
	// edgeClusterDetail: Optional. The edge cluster details, if provided, the value be used instead of contacting  the edge cluster service
	// Returns the EdgeClusterTypeEdgeResolverContract or error if something goes wrong
	NewEdgeClusterTypeEdgeResolver(
		ctx context.Context,
		edgeClusterID string,
		cursor string,
		edgeClusterDetail *EdgeClusterDetail) (EdgeClusterTypeEdgeResolverContract, error)

	// NewEdgeClusterTypeConnectionResolver creates new EdgeClusterTypeConnectionResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// edgeClusters: Mandatory. Reference the list of edge clusters
	// hasPreviousPage: Mandatory. Indicates whether more edges exist prior to the set defined by the clients arguments
	// hasNextPage: Mandatory. Indicates whether more edges exist following the set defined by the clients arguments
	// totalCount: Mandatory. The total count of matched edge clusters
	// Returns the EdgeClusterTypeConnectionResolverContract or error if something goes wrong
	NewEdgeClusterTypeConnectionResolver(
		ctx context.Context,
		edgeClusters []*edgeclusterGrpcContract.EdgeClusterWithCursor,
		hasPreviousPage bool,
		hasNextPage bool,
		totalCount int32) (EdgeClusterTypeConnectionResolverContract, error)

	// NewEdgeClusterTenantResolver creates new EdgeClusterTenatnResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// tenantID: Mandatory. The tenant unique identifier
	// Returns the EdgeClusterTenatnResolverContract or error if something goes wrong
	NewEdgeClusterTenantResolver(
		ctx context.Context,
		tenantID string) (EdgeClusterTenantResolverContract, error)

	// NewEdgeClusterProvisionDetailResolver creates new EdgeClusterProvisionDetailResolverContract and returns it
	// ctx: Mandatory. Reference to the context
	// provisionDetail: Optional. The edge cluster provisioning details
	// Returns the EdgeClusterProvisionDetailResolverContract or error if something goes wrong
	NewEdgeClusterProvisionDetailResolver(
		ctx context.Context,
		provisionDetail *edgeclusterGrpcContract.EdgeClusterProvisionDetail) (EdgeClusterProvisionDetailResolverContract, error)

	// NewIngressResolver creates new instance of the ingressResolver, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// ingress: Mandatory. The ingress details
	// Returns the new instance or error if something goes wrong
	NewIngressResolver(
		ctx context.Context,
		ingress *edgeclusterGrpcContract.Ingress) (IngressResolverContract, error)

	// NewPortResolver creates new instance of the PortResolver, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// logger: Mandatory. Reference to the logger service
	// port: Mandatory. The PostStatus details
	// Returns the new instance or error if something goes wrong
	NewPortResolver(
		ctx context.Context,
		port *edgeclusterGrpcContract.Port) (PortResolverContract, error)

	// NewEdgeClusterNodeStatusResolver creates new instance of the edgeClusterNodeStatusResolver, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// logger: Mandatory. Reference to the logger service
	// node: Mandatory. Contains information about the current status of a node.
	// Returns the new instance or error if something goes wrong
	NewEdgeClusterNodeStatusResolver(
		ctx context.Context,
		node *edgeclusterGrpcContract.EdgeClusterNodeStatus) (EdgeClusterNodeStatusResolverContract, error)

	// NewEdgeClusterNodeConditionResolver creates new instance of the edgeClusterNodeConditionResolver, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// condition: Mandatory. Contains condition information for a node.
	// Returns the new instance or error if something goes wrong
	NewEdgeClusterNodeConditionResolver(
		ctx context.Context,
		condition *edgeclusterGrpcContract.EdgeClusterNodeCondition) (EdgeClusterNodeConditionResolverContract, error)

	// NewEdgeClusterNodeAddressResolverContract creates new instance of the edgeClusterNodeAddressResolverContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// nodeAddresss: Mandatory. Contains information for the edge cluster node's address.
	// Returns the new instance or error if something goes wrong
	NewEdgeClusterNodeAddressResolverContract(
		ctx context.Context,
		nodeAddresss *edgeclusterGrpcContract.EdgeClusterNodeAddress) (EdgeClusterNodeAddressResolverContract, error)

	// NewEdgeClusterNodeSystemInfoResolverContract creates new instance of the edgeClusterNodeSystemInfoResolverContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// nodeInfo: Mandatory. Contains a set of ids/uuids to uniquely identify the node.
	// Returns the new instance or error if something goes wrong
	NewEdgeClusterNodeSystemInfoResolverContract(
		ctx context.Context,
		nodeInfo *edgeclusterGrpcContract.EdgeClusterNodeSystemInfo) (EdgeClusterNodeSystemInfoResolverContract, error)
}

// EdgeClusterNodeConditionResolverContract declares the resolver that returns Node Condition
type EdgeClusterNodeConditionResolverContract interface {
	// Type returns the type of node condition.
	// ctx: Mandatory. Reference to the context
	// Returns the type of node condition.
	Type(ctx context.Context) string

	// Status returns the status of the condition, one of True, False, Unknown.
	// ctx: Mandatory. Reference to the context
	// Returns the status of the condition, one of True, False, Unknown.
	Status(ctx context.Context) string

	// LastHeartbeatTime returns the last time we got an update on a given condition.
	// ctx: Mandatory. Reference to the context
	// Returns the last time we got an update on a given condition.
	LastHeartbeatTime(ctx context.Context) string

	// LastTransitionTime returns the last time the condition transit from one status to another.
	// ctx: Mandatory. Reference to the context
	// Returns the last time the condition transit from one status to another.
	LastTransitionTime(ctx context.Context) string

	// Reason returns the (brief) reason for the condition's last transition.
	// ctx: Mandatory. Reference to the context
	// Returns the (brief) reason for the condition's last transition.
	Reason(ctx context.Context) string

	// Message returns the human readable message indicating details about last transition.
	// ctx: Mandatory. Reference to the context
	// Returns the human readable message indicating details about last transition.
	Message(ctx context.Context) string
}

// EdgeClusterNodeAddressResolverContract declares the resolver that contains information for the edge cluster node's address.
type EdgeClusterNodeAddressResolverContract interface {
	// Type returns the edge cluster node address type, one of Hostname, ExternalIP or InternalIP.
	// ctx: Mandatory. Reference to the context
	// Returns the edge cluster node address type, one of Hostname, ExternalIP or InternalIP.
	NodeAddressType(ctx context.Context) string

	// Address returns the node address.
	// ctx: Mandatory. Reference to the context
	// Returns the node address.
	Address(ctx context.Context) string
}

// EdgeClusterNodeSystemInfoResolverContract declares the resolver that contains a set of ids/uuids to uniquely identify the node.
type EdgeClusterNodeSystemInfoResolverContract interface {
	// MachineID reported by the node. For unique machine identification in the cluster this field is preferred.
	// ctx: Mandatory. Reference to the context
	// Returns the MachineID reported by the node. For unique machine identification in the cluster this field is preferred.
	MachineID(ctx context.Context) string

	// SystemUUID reported by the node. For unique machine identification MachineID is preferred. This field is specific to Red Hat hosts
	// ctx: Mandatory. Reference to the context
	// Returns the SystemUUID reported by the node. For unique machine identification MachineID is preferred. This field is specific to Red Hat hosts
	SystemUUID(ctx context.Context) string

	// BootID reported by the node.
	// ctx: Mandatory. Reference to the context
	// Returns the Boot ID reported by the node.
	BootID(ctx context.Context) string

	// KernelVersion reported by the node from 'uname -r' (e.g. 3.16.0-0.bpo.4-amd64).
	// ctx: Mandatory. Reference to the context
	// Returns the Kernel Version reported by the node from 'uname -r' (e.g. 3.16.0-0.bpo.4-amd64).
	KernelVersion(ctx context.Context) string

	// OSImage reported by the node from /etc/os-release (e.g. Debian GNU/Linux 7 (wheezy)).
	// ctx: Mandatory. Reference to the context
	// Returns the OS Image reported by the node from /etc/os-release (e.g. Debian GNU/Linux 7 (wheezy)).
	OSImage(ctx context.Context) string

	// ContainerRuntimeVersion reported by the node through runtime remote API (e.g. docker://1.5.0).
	// ctx: Mandatory. Reference to the context
	// Returns the ContainerRuntime Version reported by the node through runtime remote API (e.g. docker://1.5.0).
	ContainerRuntimeVersion(ctx context.Context) string

	// KubeletVersion reported by the node.
	// ctx: Mandatory. Reference to the context
	// Returns the Kubelet Version reported by the node.
	KubeletVersion(ctx context.Context) string

	// KubeProxyVersion reported by the node.
	// ctx: Mandatory. Reference to the context
	// Returns the KubeProxy Version reported by the node.
	KubeProxyVersion(ctx context.Context) string

	// OperatingSystem System reported by the node
	// ctx: Mandatory. Reference to the context
	// Returns the Operating System reported by the node
	OperatingSystem(ctx context.Context) string

	// The Architecture reported by the node
	// ctx: Mandatory. Reference to the context
	// Returns the Architecture reported by the node
	Architecture(ctx context.Context) string
}

// EdgeClusterNodeStatusResolverContract declares the resolver that contains information about the current status of a node.
type EdgeClusterNodeStatusResolverContract interface {
	// Conditions is an array of current observed node conditions.
	// ctx: Mandatory. Reference to the context
	// Returns an array of current observed node conditions.
	Conditions(ctx context.Context) (*[]EdgeClusterNodeConditionResolverContract, error)

	// Addresses is the list of addresses reachable to the node.
	// ctx: Mandatory. Reference to the context
	// Returns the list of addresses reachable to the node.
	Addresses(ctx context.Context) (*[]EdgeClusterNodeAddressResolverContract, error)

	// NodeInfo is the set of ids/uuids to uniquely identify the node.
	// ctx: Mandatory. Reference to the context
	// Returns the set of ids/uuids to uniquely identify the node.
	NodeInfo(ctx context.Context) (EdgeClusterNodeSystemInfoResolverContract, error)
}

// EdgeClusterResolverContract declares the resolver that can retrieve edge cluster information
type EdgeClusterResolverContract interface {
	// ID returns edge cluster unique identifier
	// ctx: Mandatory. Reference to the context
	// Returns the edge cluster unique identifier
	ID(ctx context.Context) graphql.ID

	// Name returns edge cluster name
	// ctx: Mandatory. Reference to the context
	// Returns the edge cluster name
	Name(ctx context.Context) string

	// ClusterSecret returns edge cluster secret
	// ctx: Mandatory. Reference to the context
	// Returns the edge cluster secret
	ClusterSecret(ctx context.Context) string

	// ClusterType returns the edge cluster current type
	// ctx: Mandatory. Reference to the context
	// Returns the edge cluster current type or error if something went wrong
	ClusterType(ctx context.Context) (string, error)

	// Tenant returns edge cluster tenant
	// ctx: Mandatory. Reference to the context
	// Returns the edge cluster tenant
	Tenant(ctx context.Context) (EdgeClusterTenantResolverContract, error)

	// ProvisionDetail returns edge cluster provisioning detail
	// ctx: Mandatory. Reference to the context
	// Returns the edge cluster provisioning detail
	ProvisionDetail(ctx context.Context) (EdgeClusterProvisionDetailResolverContract, error)

	// Ingress returns the ingress details of the edge cluster master node
	// ctx: Mandatory. Reference to the context
	// Returns the ingress details of the edge cluster master node
	Nodes(ctx context.Context) (*[]EdgeClusterNodeStatusResolverContract, error)
}

// EdgeClusterTypeConnectionResolverContract declares the resolver that returns edge cluster edge compatible with graphql-relay
type EdgeClusterTypeConnectionResolverContract interface {
	// PageInfo returns the paging information compatible with graphql-relay
	// ctx: Mandatory. Reference to the context
	// Returns the paging information
	PageInfo(ctx context.Context) (relay.PageInfoResolverContract, error)

	// Edges returns the edge cluster edges compatible with graphql-relay
	// ctx: Mandatory. Reference to the context
	// Returns the edge cluster edges
	Edges(ctx context.Context) (*[]EdgeClusterTypeEdgeResolverContract, error)

	// TotalCount returns total count of the matched edge clusters
	// ctx: Mandatory. Reference to the context
	// Returns the total count of the matched edge cluster
	TotalCount(ctx context.Context) *int32
}

// EdgeClusterTypeEdgeResolverContract declares the resolver that returns edge cluster edge compatible with graphql-relay
type EdgeClusterTypeEdgeResolverContract interface {
	// Node returns the edge cluster resolver
	// ctx: Mandatory. Reference to the context
	// Returns the edge cluster resolver or error if something goes wrong
	Node(ctx context.Context) (EdgeClusterResolverContract, error)

	// Cursor returns the cursor for the edge cluster edge compatible with graphql-relay
	// ctx: Mandatory. Reference to the context
	// Returns the cursor
	Cursor(ctx context.Context) string
}

// EdgeClusterTenantResolverContract declares the resolver that returns edge cluster tenant
type EdgeClusterTenantResolverContract interface {
	// ID returns tenant unique identifier
	// ctx: Mandatory. Reference to the context
	// Returns the tenant  unique identifier
	ID(ctx context.Context) graphql.ID

	// Name returns tenant name
	// ctx: Mandatory. Reference to the context
	// Returns the tenant name
	Name(ctx context.Context) string
}

// IngressResolverContract declares the resolver that returns Ingress
type IngressResolverContract interface {
	// IP is set for load-balancer ingress points that are IP based
	// (typically GCE or OpenStack load-balancers)
	// ctx: Mandatory. Reference to the context
	// Returns the IP is set for load-balancer ingress points that are IP based
	IP(ctx context.Context) *string

	// Hostname is set for load-balancer ingress points that are DNS based
	// (typically AWS load-balancers)
	// ctx: Mandatory. Reference to the context
	// Returns the Hostname is set for load-balancer ingress points that are DNS based
	Hostname(ctx context.Context) *string
}

// PortResolverContract declares the resolver that returns Port
type PortResolverContract interface {
	// Port returns the port number of the edge-cluster master port of which status is recorded here
	// ctx: Mandatory. Reference to the context
	// Returns the port number of the edge-cluster master port of which status is recorded here
	Port(ctx context.Context) int32

	// Protocol returns the protocol of the edge-cluster master port of which status is recorded here
	// ctx: Mandatory. Reference to the context
	// Returns the protocol of the edge-cluster master port of which status is recorded here
	Protocol(ctx context.Context) string
}

// EdgeClusterProvisionDetailResolverContract declares the resolver that returns edge cluster provisioning details
type EdgeClusterProvisionDetailResolverContract interface {
	// Ingress returns the ingress details of the edge cluster master node
	// ctx: Mandatory. Reference to the context
	// Returns the ingress details of the edge cluster master node
	Ingress(ctx context.Context) (*[]IngressResolverContract, error)

	// Ports is a list of records of edge-cluster master ports
	// ctx: Mandatory. Reference to the context
	// Returns the Ports is a list of records of edge-cluster master ports
	Ports(ctx context.Context) (*[]PortResolverContract, error)

	// KubeconfigContent returns the edge cluster Kubeconfig content
	// ctx: Mandatory. Reference to the context
	// Returns the edge cluster Kubeconfig content
	KubeconfigContent(ctx context.Context) *string
}

type EdgeClusterDetail struct {
	EdgeCluster     *edgeclusterGrpcContract.EdgeCluster
	ProvisionDetail *edgeclusterGrpcContract.EdgeClusterProvisionDetail
}

type EdgeClusterClusterEdgeClusterInputArgument struct {
	EdgeClusterID graphql.ID
}

type EdgeClusterEdgeClustersInputArgument struct {
	relay.ConnectionArgument
	EdgeClusterIDs *[]graphql.ID
	SortOption     *string
}
