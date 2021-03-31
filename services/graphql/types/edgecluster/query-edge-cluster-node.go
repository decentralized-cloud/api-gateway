// packae edgecluster implements used edge cluster related types in the GraphQL transport layer
package edgecluster

import "context"

// EdgeClusterNodeConditionResolverContract declares the resolver that returns the current service state of node
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

// EdgeClusterNodeStatusResolverContract declares the resolver that contains information about the status of a node.
type EdgeClusterNodeStatusResolverContract interface {
	// Conditions is an array of current observed node conditions.
	// ctx: Mandatory. Reference to the context
	// Returns an array of current observed node conditions resolver or error if something goes wrong.
	Conditions(ctx context.Context) ([]EdgeClusterNodeConditionResolverContract, error)

	// Addresses is the list of addresses reachable to the node.
	// ctx: Mandatory. Reference to the context
	// Returns the list of addresses reachable to the node resolver or error if something goes wrong.
	Addresses(ctx context.Context) ([]EdgeClusterNodeAddressResolverContract, error)

	// NodeInfo is the set of ids/uuids to uniquely identify the node.
	// ctx: Mandatory. Reference to the context
	// Returns the set of ids/uuids to uniquely identify the node resolver or error if something goes wrong.
	NodeInfo(ctx context.Context) (EdgeClusterNodeSystemInfoResolverContract, error)
}

// EdgeClusterNodeResolverContract declares the resolver that contains information about the edge cluster node.
type EdgeClusterNodeResolverContract interface {
	// Metadata contains the node metadata
	// ctx: Mandatory. Reference to the context
	// Returns the node metadata resolver or error if something goes wrong.
	Metadata(ctx context.Context) (EdgeClusterObjectMetadataResolverContract, error)

	// Status contains the most recently observed status of the node
	// ctx: Mandatory. Reference to the context
	// Returns the most recently observed status of the node resolver or error if something goes wrong.
	Status(ctx context.Context) (EdgeClusterNodeStatusResolverContract, error)
}
