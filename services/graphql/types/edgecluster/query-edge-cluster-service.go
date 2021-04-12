// packae edgecluster implements used edge cluster related types in the GraphQL transport layer
package edgecluster

import (
	"context"

	edgeclusterGrpcContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
)

type EdgeClusterServiceResolverCreatorContract interface {
	// NewEdgeClusterServiceResolver creates new instance of the ServiceResolverContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// service: Mandatory. Contains information about the edge cluster service
	// Returns the new instance or error if something goes wrong
	NewEdgeClusterServiceResolver(
		ctx context.Context,
		service *edgeclusterGrpcContract.EdgeClusterService) (ServiceResolverContract, error)

	// NewServiceStatusResolver creates new instance of the ServiceStatusResolverContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// serviceStatus: Mandatory. Contains information about the service status
	// Returns the new instance or error if something goes wrong
	NewServiceStatusResolver(
		ctx context.Context,
		serviceStatus *edgeclusterGrpcContract.ServiceStatus) (ServiceStatusResolverContract, error)

	// NewServicePortResolver creates new instance of the ServicePortResolverContract, setting up all dependencies and returns the instance
	// ctx: Mandatory. Reference to the context
	// servicePort: Mandatory. Contains information about the service port
	// Returns the new instance or error if something goes wrong
	NewServicePortResolver(
		ctx context.Context,
		servicePort *edgeclusterGrpcContract.ServicePort) (ServicePortResolverContract, error)
}

// ServiceStatusResolverContract declares the resolver that returns the most recently observed
// status of the existing edge cluster service
type ServiceStatusResolverContract interface {
	// LoadBalancer contains the current status of the load-balancer
	// ctx: Mandatory. Reference to the context
	// Returns the load balancer status resolver or error if something goes wrong.
	LoadBalancer(ctx context.Context) (LoadBalancerStatusResolverContract, error)

	// KubeconfigContent returns the edge cluster Kubeconfig content
	// ctx: Mandatory. Reference to the context
	// Returns the edge cluster Kubeconfig content
	Conditions(ctx context.Context) ([]ServiceConditionResolverContract, error)
}

// ServicePortResolverContract declares the resolver that contains information on service's port.
type ServicePortResolverContract interface {
	// Type is the name of this port within the service
	// ctx: Mandatory. Reference to the context
	// Returns he name of this port within the service
	Name(ctx context.Context) string

	// Protocol is the IP protocol for this port
	// ctx: Mandatory. Reference to the context
	// Returns the IP protocol for this port
	Protocol(ctx context.Context) string

	// Port is the port that will be exposed by this service
	// ctx: Mandatory. Reference to the context
	// Returns the port that will be exposed by this service
	Port(ctx context.Context) int32

	// TargetPort is the number or name of the port to access on the pods targeted by the service
	// ctx: Mandatory. Reference to the context
	// Returns the number or name of the port to access on the pods targeted by the service
	TargetPort(ctx context.Context) string

	// NodePort is the port on each node on which this service is exposed when type is
	// NodePort or LoadBalancer
	// ctx: Mandatory. Reference to the context
	// Returns the port on each node on which this service is exposed
	NodePort(ctx context.Context) int32
}

// ServiceResolverContract declares the resolver that contains information about the edge cluster service
type ServiceResolverContract interface {
	// Metadata contains the service metadata
	// ctx: Mandatory. Reference to the context
	// Returns the service metadata resolver or error if something goes wrong.
	Metadata(ctx context.Context) (ObjectMetaResolverContract, error)

	// Status contains the most recently observed status of the service
	// ctx: Mandatory. Reference to the context
	// Returns the most recently observed status of the service resolver or error if something goes wrong.
	Status(ctx context.Context) (ServiceStatusResolverContract, error)

	// Status contains the specification of the desired behavior of the service
	// ctx: Mandatory. Reference to the context
	// Returns the specification of the desired behavior of the service resolver or error if something goes wrong.
	Spec(ctx context.Context) (ServiceSpecResolverContract, error)
}

type EdgeClusterServiceInputArgument struct {
	Namespace *string
}
