// Package endpoint implements different endpoint services required by the api-gateway service
package endpoint

import "github.com/go-kit/kit/endpoint"

// EndpointCreatorContract declares the contract that creates endpoints to create new edgeCluster,
// read, update and delete existing edgeClusters.
type EndpointCreatorContract interface {
	// GraphQLEndpoint creates GraphQL endpoint
	// Returns the GraphQL endpoint
	GraphQLEndpoint() endpoint.Endpoint
}
