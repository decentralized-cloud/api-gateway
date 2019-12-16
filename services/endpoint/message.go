// Package endpoint implements different endpoint services required by the api-gateway service
package endpoint

import (
	"github.com/graph-gophers/graphql-go"
)

// GraphQLRequest contains the request to process the GraphQL request
type GraphQLRequest struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

// GraphQLResponse contains the result of processing the GraphQL request
type GraphQLResponse struct {
	Err      error
	Response graphql.Response
}
