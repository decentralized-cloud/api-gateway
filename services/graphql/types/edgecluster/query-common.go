// packae edgecluster implements used edge cluster related types in the GraphQL transport layer
package edgecluster

import (
	"context"

	edgeclusterGrpcContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
	"github.com/graph-gophers/graphql-go"
)

type CommonResolverCreatorContract interface {
	// NewObjectMetaResolver creates new instance of the ObjectMetaResolverContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// metadata: Mandatory. Contains the object metadata.
	// Returns the new instance or error if something goes wrong
	NewObjectMetaResolver(
		ctx context.Context,
		metadata *edgeclusterGrpcContract.ObjectMeta) (ObjectMetaResolverContract, error)

	// NewPortStatusResolverContract creates new instance of the PortStatusResolverContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// portStatus: Mandatory. Contains the object portStatus.
	// Returns the new instance or error if something goes wrong
	NewPortStatusResolverContract(
		ctx context.Context,
		portStatus *edgeclusterGrpcContract.PortStatus) (PortStatusResolverContract, error)

	// NewLoadBalancerIngressResolver creates new instance of the loadBalancerIngressResolverContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// loadBalancerIngress: Mandatory. The load balancer ingress details
	// Returns the new instance or error if something goes wrong
	NewLoadBalancerIngressResolver(
		ctx context.Context,
		loadBalancerIngress *edgeclusterGrpcContract.LoadBalancerIngress) (LoadBalancerIngressResolverContract, error)

	// NewLoadBalancerStatusResolver creates new instance of the loadBalancerStatusResolverContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// Returns the new instance or error if something goes wrong
	NewLoadBalancerStatusResolver(
		ctx context.Context,
		loadBalancerStatus *edgeclusterGrpcContract.LoadBalancerStatus) (LoadBalancerStatusResolverContract, error)

	// NewServiceConditionResolver creates new instance of the ServiceConditionResolverContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// serviceCondition: Mandatory. Contains condition information for a service.
	// Returns the new instance or error if something goes wrong
	NewServiceConditionResolver(
		ctx context.Context,
		serviceCondition *edgeclusterGrpcContract.ServiceCondition) (ServiceConditionResolverContract, error)

	// NewServiceSpecResolver creates new instance of the ServiceSpecResolverContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// serviceSpec: Mandatory. Contains spec information for a service.
	// Returns the new instance or error if something goes wrong
	NewServiceSpecResolver(
		ctx context.Context,
		serviceSpec *edgeclusterGrpcContract.ServiceSpec) (ServiceSpecResolverContract, error)
}

// ObjectMetaResolverContract declares the standard edge cluster object's metadata.
type ObjectMetaResolverContract interface {
	// ID returns the object unique identitfier
	// ctx: Mandatory. Reference to the context
	// Returns the object unique identitfier
	ID(ctx context.Context) graphql.ID

	// Name returns the object name
	// ctx: Mandatory. Reference to the context
	// Returns the object name
	Name(ctx context.Context) string

	// Namespace returns the object namespace
	// ctx: Mandatory. Reference to the context
	// Returns the object namespace
	Namespace(ctx context.Context) string
}

// PortStatusResolverContract represents the error condition of a service port
type PortStatusResolverContract interface {
	// Port is the port number of the service port of which status is recorded here
	// ctx: Mandatory. Reference to the context
	// Returns the port number of the service port of which status is recorded here
	Port(ctx context.Context) int32

	// Protocol is the protocol of the service port of which status is recorded here
	// ctx: Mandatory. Reference to the context
	// Returns the protocol of the service port of which status is recorded here
	Protocol(ctx context.Context) string

	// Error is to record the problem with the service port
	// ctx: Mandatory. Reference to the context
	// Returns the problem with the service port
	Error(ctx context.Context) *string
}

// LoadBalancerIngressResolverContract represents the status of a load-balancer ingress point:
// traffic intended for the service should be sent to an ingress point.
type LoadBalancerIngressResolverContract interface {
	// IP is set for load-balancer ingress points that are IP based
	// ctx: Mandatory. Reference to the context
	// Returns the IP that is set for load-balancer ingress points that are IP based
	IP(ctx context.Context) string

	// Hostname is set for load-balancer ingress points that are DNS based
	// ctx: Mandatory. Reference to the context
	// Returns the hostname that is set for load-balancer ingress points that are DNS based
	Hostname(ctx context.Context) string

	// PortStatus is a list of records of service ports
	// If used, every port defined in the service should have an entry in it
	// ctx: Mandatory. Reference to the context
	// Returns an array of port status resolver or error if something goes wrong.
	PortStatus(ctx context.Context) ([]PortStatusResolverContract, error)
}

// LoadBalancerStatusResolverContract represents the status of a load-balancer.
type LoadBalancerStatusResolverContract interface {
	// Ingress is a list containing ingress points for the load-balancer.
	// Traffic intended for the service should be sent to these ingress points.
	// ctx: Mandatory. Reference to the context
	// Returns an array of load balancer ingress resolver or error if something goes wrong.
	Ingress(ctx context.Context) ([]LoadBalancerIngressResolverContract, error)
}

// ServiceConditionResolverContract declares the resolver that returns the current service state of service
type ServiceConditionResolverContract interface {
	// Type is the type of the condition
	// ctx: Mandatory. Reference to the context
	// Returns the type of the condition
	Type(ctx context.Context) string

	// Status returns the status of the condition
	// ctx: Mandatory. Reference to the context
	// Returns the status of the condition
	Status(ctx context.Context) string

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

// ServiceSpecResolverContract declares the resolver that returns the the specification of the desired behavior
// of the existing edge cluster service
type ServiceSpecResolverContract interface {
	// Ports is the list of ports that are exposed by this service
	// ctx: Mandatory. Reference to the context
	// Returns an array of service port resolver or error if something goes wrong.
	Ports(ctx context.Context) ([]ServicePortResolverContract, error)

	// ClusterIPs is a list of IP addresses assigned to this service
	// ctx: Mandatory. Reference to the context
	// Returns the list of IP addresses assigned to this service
	ClusterIPs(ctx context.Context) []string

	// Type determines how the Service is exposed
	// ctx: Mandatory. Reference to the context
	// Returns how the service is exposed
	Type(ctx context.Context) string

	// ExternalIPs is a list of IP addresses for which nodes in the cluster will also accept traffic for this service
	// ctx: Mandatory. Reference to the context
	// Returns the list of IP addresses for which nodes in the cluster will also accept traffic for this service
	ExternalIPs(ctx context.Context) []string

	// ExternalName is the external reference that discovery mechanisms will return as an alias for this service (e.g. a DNS CNAME record)
	// ctx: Mandatory. Reference to the context
	// Returns the the external reference that discovery mechanisms will return as an alias for this service (e.g. a DNS CNAME record)
	ExternalName(ctx context.Context) *string
}
