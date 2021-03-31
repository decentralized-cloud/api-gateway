// packae edgecluster implements used edge cluster related types in the GraphQL transport layer
package edgecluster

import (
	"context"
)

// EdgeClusterPodConditionResolverContract declares the resolver that returns the current service state of pod
type EdgeClusterPodConditionResolverContract interface {
	// Type returns the type of pod condition.
	// ctx: Mandatory. Reference to the context
	// Returns the type of pod condition.
	Type(ctx context.Context) string

	// Status returns the status of the condition, one of True, False, Unknown.
	// ctx: Mandatory. Reference to the context
	// Returns the status of the condition, one of True, False, Unknown.
	Status(ctx context.Context) string

	// LastHeartbeatTime returns the last time we got an update on a given condition.
	// ctx: Mandatory. Reference to the context
	// Returns the last time we got an update on a given condition.
	LastProbeTime(ctx context.Context) string

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

// EdgeClusterPodStatusResolverContract declares the resolver that contains the most recently observed status of the existing edge cluster pod
type EdgeClusterPodStatusResolverContract interface {
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
	Conditions(ctx context.Context) ([]EdgeClusterPodConditionResolverContract, error)
}

// EdgeClusterPodSpecResolverContract declares the resolver that contains the specification of the desired behavior of the existing edge cluster pod
type EdgeClusterPodSpecResolverContract interface {
	// NodeName returns the name of the node where the Pod is deployed into.
	// ctx: Mandatory. Reference to the context
	// Returns the name of the node where the Pod is deployed into.
	NodeName(ctx context.Context) string
}

// EdgeClusterPodResolverContract declares the resolver that contains information about the edge cluster pod
type EdgeClusterPodResolverContract interface {
	// Metadata contains the pod metadata
	// ctx: Mandatory. Reference to the context
	// Returns the pod metadata resolver or error if something goes wrong.
	Metadata(ctx context.Context) (EdgeClusterObjectMetadataResolverContract, error)

	// Status contains the most recently observed status of the pod
	// ctx: Mandatory. Reference to the context
	// Returns the most recently observed status of the pod resolver or error if something goes wrong.
	Status(ctx context.Context) (EdgeClusterPodStatusResolverContract, error)

	// Status contains the specification of the desired behavior of the pod
	// ctx: Mandatory. Reference to the context
	// Returns the specification of the desired behavior of the pod resolver or error if something goes wrong.
	Spec(ctx context.Context) (EdgeClusterPodSpecResolverContract, error)
}

type EdgeClusterPodInputArgument struct {
	NodeName  *string
	Namespace *string
}
