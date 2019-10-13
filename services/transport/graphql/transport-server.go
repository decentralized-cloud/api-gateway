// Package graphql implements functions to expose api-gateway service endpoint using GraphQL protocol.
package graphql

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/decentralized-cloud/api-gateway/services/configuration"
	"github.com/decentralized-cloud/api-gateway/services/transport"
	"github.com/decentralized-cloud/api-gateway/services/transport/graphql/types"
	"github.com/friendsofgo/graphiql"
	"github.com/gobuffalo/packr"
	"github.com/graph-gophers/graphql-go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"github.com/savsgio/atreugo/v9"
	"go.uber.org/zap"
)

type transportService struct {
	logger               *zap.Logger
	configurationService configuration.ConfigurationContract
	resolverCreator      types.ResolverCreatorContract
	schema               *graphql.Schema
}

// NewTransportService creates new instance of the GRPCService, setting up all dependencies and returns the instance
// logger: Mandatory. Reference to the logger service
// configurationService: Mandatory. Reference to the service that provides required configurations
// Returns the new service or error if something goes wrong
func NewTransportService(
	logger *zap.Logger,
	configurationService configuration.ConfigurationContract,
	resolverCreator types.ResolverCreatorContract) (transport.TransportContract, error) {
	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if configurationService == nil {
		return nil, commonErrors.NewArgumentNilError("configurationService", "configurationService is required")
	}

	if resolverCreator == nil {
		return nil, commonErrors.NewArgumentNilError("resolverCreator", "resolverCreator is required")
	}

	return &transportService{
		logger:               logger,
		configurationService: configurationService,
		resolverCreator:      resolverCreator,
	}, nil
}

// Start starts the GraphQL transport service
// Returns error if something goes wrong
func (service *transportService) Start() error {
	config := &atreugo.Config{GracefulShutdown: true}
	var err error

	host, err := service.configurationService.GetHost()
	if err != nil {
		return err
	}

	port, err := service.configurationService.GetPort()
	if err != nil {
		return err
	}

	config.Addr = fmt.Sprintf("%s:%d", host, port)
	server := atreugo.New(config)

	box := packr.NewBox("../../../contract/graphql/schema")
	graphqlSchema, err := box.FindString("schema.graphql")
	if err != nil {
		return err
	}

	// Adding the schema part as this is required by the graphql-go library and is not generated by the schema-generator
	graphqlSchema = `
		schema {
		  query: Query
		  mutation: Mutation
		}
	` + "\n" + graphqlSchema

	rootResolver, err := service.resolverCreator.NewRootResolver(context.Background())
	if err != nil {
		return err
	}

	service.schema = graphql.MustParseSchema(graphqlSchema, rootResolver)

	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/query")
	if err != nil {
		return err
	}

	server.Path("POST", "/query", service.graphQLHandler)
	server.NetHTTPPath("GET", "/graphiql", graphiqlHandler)
	service.logger.Info("GraphQL server started", zap.String("address", config.Addr))

	return server.ListenAndServe()
}

// Stop stops the GraphQL transport service
// Returns error if something goes wrong
func (service *transportService) Stop() error {
	return nil
}

func (service *transportService) graphQLHandler(ctx *atreugo.RequestCtx) error {
	var params struct {
		Query         string                 `json:"query"`
		OperationName string                 `json:"operationName"`
		Variables     map[string]interface{} `json:"variables"`
	}

	if err := json.Unmarshal(ctx.PostBody(), &params); err != nil {
		ctx.Error(err.Error(), http.StatusBadRequest)

		return nil
	}

	response := service.schema.Exec(ctx, params.Query, params.OperationName, params.Variables)

	return ctx.JSONResponse(response, http.StatusOK)
}
