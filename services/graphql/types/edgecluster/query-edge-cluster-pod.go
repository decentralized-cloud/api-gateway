// packae edgecluster implements used edge cluster related types in the GraphQL transport layer
package edgecluster

import (
	"context"

	edgeclusterGrpcContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
)

type EdgeClusterPodResolverCreatorContract interface {
	// NewEdgeClusterPodResolver creates new instance of the PodResolverContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// pod: Mandatory. Contains information about the edge cluster pod
	// Returns the new instance or error if something goes wrong
	NewEdgeClusterPodResolver(
		ctx context.Context,
		pod *edgeclusterGrpcContract.EdgeClusterPod) (PodResolverContract, error)

	// NewPodStatusResolver creates new instance of the PodStatusResolverContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// status: Mandatory. Contains information about the edge cluster pod status
	// Returns the new instance or error if something goes wrong
	NewPodStatusResolver(
		ctx context.Context,
		status *edgeclusterGrpcContract.PodStatus) (PodStatusResolverContract, error)

	// NewPodSpecResolver creates new instance of the PodSpecResolverContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// spec: Mandatory. Contains information about the edge cluster pod specification
	// Returns the new instance or error if something goes wrong
	NewPodSpecResolver(
		ctx context.Context,
		spec *edgeclusterGrpcContract.PodSpec) (PodSpecResolverContract, error)

	// NewPodConditionResolver creates new instance of the PodConditionResolverContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// podCondition: Mandatory. Contains condition information for a pod.
	// Returns the new instance or error if something goes wrong
	NewPodConditionResolver(
		ctx context.Context,
		podCondition *edgeclusterGrpcContract.PodCondition) (PodConditionResolverContract, error)
}

// PodConditionResolverContract declares the resolver that returns the current service state of pod
type PodConditionResolverContract interface {
	// Type is the type of the condition
	// ctx: Mandatory. Reference to the context
	// Returns the type of the condition
	Type(ctx context.Context) string

	// Status returns the status of the condition
	// ctx: Mandatory. Reference to the context
	// Returns the status of the condition
	Status(ctx context.Context) string

	// LastHeartbeatTime returns the last time we got an update on a given condition
	// ctx: Mandatory. Reference to the context
	// Returns the last time we got an update on a given condition
	LastProbeTime(ctx context.Context) string

	// LastTransitionTime returns the last time the condition transitioned from one status to another
	// ctx: Mandatory. Reference to the context
	// Returns the last time the condition transitioned from one status to another
	LastTransitionTime(ctx context.Context) string

	// Reason returns the Unique, one-word, CamelCase reason for the condition's last transition
	// ctx: Mandatory. Reference to the context
	// Returns the Unique, one-word, CamelCase reason for the condition's last transition
	Reason(ctx context.Context) string

	// Message returns the human-readable message indicating details about last transition
	// ctx: Mandatory. Reference to the context
	// Returns the human-readable message indicating details about last transition
	Message(ctx context.Context) string
}

// PodStatusResolverContract declares the resolver that contains the most recently observed status of the existing edge cluster pod
type PodStatusResolverContract interface {
	// HostIP returns the IP address allocated to the pod. Routable at least within the cluster.
	// ctx: Mandatory. Reference to the context
	// Returns the IP address allocated to the pod. Routable at least within the cluster.
	HostIP(ctx context.Context) string

	// PodIP returns the IP address allocated to the pod. Routable at least within the cluster.
	// ctx: Mandatory. Reference to the context
	// Returns the IP address allocated to the pod. Routable at least within the cluster.
	PodIP(ctx context.Context) string

	// Conditions is an array of current observed node conditions.
	// ctx: Mandatory. Reference to the context
	// Returns an array of current observed node conditions resolver or error if something goes wrong.
	Conditions(ctx context.Context) ([]PodConditionResolverContract, error)
}

// PodSpecResolverContract declares the resolver that contains the specification of the desired behavior of the existing edge cluster pod
type PodSpecResolverContract interface {
	// NodeName returns the name of the node where the Pod is deployed into.
	// ctx: Mandatory. Reference to the context
	// Returns the name of the node where the Pod is deployed into.
	NodeName(ctx context.Context) string
}

// PodResolverContract declares the resolver that contains information about the edge cluster pod
type PodResolverContract interface {
	// Metadata contains the pod metadata
	// ctx: Mandatory. Reference to the context
	// Returns the pod metadata resolver or error if something goes wrong.
	Metadata(ctx context.Context) (ObjectMetaResolverContract, error)

	// Status contains the most recently observed status of the pod
	// ctx: Mandatory. Reference to the context
	// Returns the most recently observed status of the pod resolver or error if something goes wrong.
	Status(ctx context.Context) (PodStatusResolverContract, error)

	// Status contains the specification of the desired behavior of the pod
	// ctx: Mandatory. Reference to the context
	// Returns the specification of the desired behavior of the pod resolver or error if something goes wrong.
	Spec(ctx context.Context) (PodSpecResolverContract, error)
}

type EdgeClusterPodInputArgument struct {
	NodeName  *string
	Namespace *string
}
