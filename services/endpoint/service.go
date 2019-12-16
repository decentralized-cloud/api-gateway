// Package endpoint implements different endpoint services required by the api-gateway service
package endpoint

import (
	"context"

	"github.com/decentralized-cloud/api-gateway/services/transport/https/graphql/types"
	"github.com/go-kit/kit/endpoint"
	"github.com/gobuffalo/packr"
	"github.com/graph-gophers/graphql-go"
	commonErrors "github.com/micro-business/go-core/system/errors"
)

type endpointCreatorService struct {
	schema *graphql.Schema
}

// NewEndpointCreatorService creates new instance of the EndpointCreatorService, setting up all dependencies and returns the instance
// Returns the new service or error if something goes wrong
func NewEndpointCreatorService(
	resolverCreator types.ResolverCreatorContract) (EndpointCreatorContract, error) {

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	box := packr.NewBox("../../contract/graphql/schema")
	graphqlSchema, err := box.FindString("schema.graphql")
	if err != nil {
		return nil, NewUnknownErrorWithError("Failed to find schema.graphql", err)
	}

	// Adding the schema part as this is required by the graphql-go library and is not generated by the schema-generator
	graphqlSchema = `
		schema {
		  query: Query
		  mutation: Mutation
		}
	` + "\n" + graphqlSchema

	rootResolver, err := resolverCreator.NewRootResolver(context.Background())
	if err != nil {
		return nil, NewUnknownErrorWithError("Failed to create the root resolver", err)
	}

	schema := graphql.MustParseSchema(graphqlSchema, rootResolver)

	return &endpointCreatorService{
		schema: schema,
	}, nil
}

// GraphQLEndpoint creates GraphQL endpoint
// Returns the GraphQL endpoint
func (service *endpointCreatorService) GraphQLEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		if ctx == nil {
			return &GraphQLResponse{
				Err: commonErrors.NewArgumentNilError("ctx", "ctx is required"),
			}, nil
		}

		if request == nil {
			return &GraphQLResponse{
				Err: commonErrors.NewArgumentNilError("request", "request is required"),
			}, nil
		}

		castedRequest := request.(*GraphQLRequest)

		return service.schema.Exec(ctx, castedRequest.Query, castedRequest.OperationName, castedRequest.Variables), nil
	}
}
