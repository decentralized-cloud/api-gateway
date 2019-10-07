// Package graphql implements functions to expose api-gateway service endpoint using GraphQL protocol.
package graphql

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/decentralized-cloud/api-gateway/services/configuration"
	"github.com/decentralized-cloud/api-gateway/services/transport"
	"github.com/gobuffalo/packr"
	"github.com/graph-gophers/graphql-go"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"github.com/savsgio/atreugo/v9"
	"go.uber.org/zap"
)

type transportService struct {
	logger               *zap.Logger
	configurationService configuration.ConfigurationContract
	schema               *graphql.Schema
}

// NewTransportService creates new instance of the GRPCService, setting up all dependencies and returns the instance
// logger: Mandatory. Reference to the logger service
// configurationService: Mandatory. Reference to the service that provides required configurations
// Returns the new service or error if something goes wrong
func NewTransportService(
	logger *zap.Logger,
	configurationService configuration.ConfigurationContract) (transport.TransportContract, error) {
	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if configurationService == nil {
		return nil, commonErrors.NewArgumentNilError("configurationService", "configurationService is required")
	}

	return &transportService{
		logger:               logger,
		configurationService: configurationService,
	}, nil
}

// Start starts the GraphQL transport service
// Returns error if something goes wrong
func (service *transportService) Start() error {
	config := &atreugo.Config{}
	server := atreugo.New(config)
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

	box := packr.NewBox("./schema")
	graphqlSchema, err := box.FindString("schema.graphql")
	if err != nil {
		return err
	}

	service.schema = graphql.MustParseSchema(graphqlSchema, &query{})
	server.Path("POST", "/query", service.graphQLHandler)
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

type query struct{}

func (_ *query) Hello() string { return "Hello, world!" }
