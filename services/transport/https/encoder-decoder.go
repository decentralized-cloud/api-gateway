// Package https implements functions to expose api-gateway service endpoint using HTTPS/GraphQL protocol.
package https

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/decentralized-cloud/api-gateway/services/endpoint"
)

// decodeGraphQLRequest decodes GraphQL request message from GRPC object to business object
// context: Mandatory The reference to the context
// request: Mandatory. The reference to the GRPC request
// Returns either the decoded request or error if something goes wrong
func decodeGraphQLRequest(
	ctx context.Context,
	request *http.Request) (interface{}, error) {
	var graphqlRequest endpoint.GraphQLRequest

	query, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(query, &graphqlRequest); err != nil {
		return nil, err
	}

	return &graphqlRequest, nil
}

// encodeGraphQLResponse encodes GraphQL response from business object to GRPC object
// context: Optional The reference to the context
// request: Mandatory. The reference to the business response
// Returns either the decoded response or error if something goes wrong
func encodeGraphQLResponse(
	ctx context.Context,
	writer http.ResponseWriter, response interface{}) error {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	writer.Header().Set("Access-Control-Allow-Headers", "Origin,DNT,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Range,Authorization")

	return json.NewEncoder(writer).Encode(response)
}
