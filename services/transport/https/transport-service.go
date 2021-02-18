// Package https implements functions to expose api-gateway service endpoint using HTTPS/GraphQL protocol.
package https

import (
	"fmt"
	"net/http"

	"github.com/decentralized-cloud/api-gateway/services/configuration"
	"github.com/decentralized-cloud/api-gateway/services/endpoint"
	"github.com/decentralized-cloud/api-gateway/services/transport"
	"github.com/friendsofgo/graphiql"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/micro-business/go-core/gokit/middleware"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/savsgio/atreugo/v11"
	"go.uber.org/zap"
)

type transportService struct {
	logger                    *zap.Logger
	configurationService      configuration.ConfigurationContract
	endpointCreatorService    endpoint.EndpointCreatorContract
	middlewareProviderService middleware.MiddlewareProviderContract
	jwksURL                   string
	graphQLHandler            *httpTransport.Server
}

// NewTransportService creates new instance of the transportService, setting up all dependencies and returns the instance
// logger: Mandatory. Reference to the logger service
// configurationService: Mandatory. Reference to the service that provides required configurations
// middlewareProviderService: Mandatory. Reference to the service that provides different go-kit middlewares
// Returns the new service or error if something goes wrong
func NewTransportService(
	logger *zap.Logger,
	configurationService configuration.ConfigurationContract,
	endpointCreatorService endpoint.EndpointCreatorContract,
	middlewareProviderService middleware.MiddlewareProviderContract) (transport.TransportContract, error) {
	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if configurationService == nil {
		return nil, commonErrors.NewArgumentNilError("configurationService", "configurationService is required")
	}

	if endpointCreatorService == nil {
		return nil, commonErrors.NewArgumentNilError("endpointCreatorService", "endpointCreatorService is required")
	}

	if middlewareProviderService == nil {
		return nil, commonErrors.NewArgumentNilError("middlewareProviderService", "middlewareProviderService is required")
	}

	jwksURL, err := configurationService.GetJwksURL()
	if err != nil {
		return nil, err
	}

	return &transportService{
		logger:                    logger,
		configurationService:      configurationService,
		endpointCreatorService:    endpointCreatorService,
		middlewareProviderService: middlewareProviderService,
		jwksURL:                   jwksURL,
	}, nil
}

// Start starts the GraphQL transport service
// Returns error if something goes wrong
func (service *transportService) Start() error {
	service.setupHandlers()

	config := atreugo.Config{GracefulShutdown: true}
	var err error

	host, err := service.configurationService.GetHttpHost()
	if err != nil {
		return err
	}

	port, err := service.configurationService.GetHttpPort()
	if err != nil {
		return err
	}

	config.Addr = fmt.Sprintf("%s:%d", host, port)
	server := atreugo.New(config)

	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		return err
	}

	server.UseBefore()

	server.NetHTTPPath("POST", "/graphql", service.graphQLHandler)
	server.NetHTTPPath("GET", "/graphiql", graphiqlHandler)

	server.Path("GET", "/live", service.livenessCheckHandler)
	server.Path("GET", "/ready", service.readinessCheckHandler)

	server.NetHTTPPath("GET", "/metrics", promhttp.Handler())

	service.logger.Info("HTTPS service started", zap.String("address", config.Addr))

	return server.ListenAndServe()
}

// Stop stops the GraphQL transport service
// Returns error if something goes wrong
func (service *transportService) Stop() error {
	return nil
}

func (service *transportService) setupHandlers() {
	endpoint := service.endpointCreatorService.GraphQLEndpoint()
	endpoint = service.middlewareProviderService.CreateLoggingMiddleware("GraphQL")(endpoint)
	endpoint = service.createAuthMiddleware("GraphQL")(endpoint)
	service.graphQLHandler = httpTransport.NewServer(
		endpoint,
		decodeGraphQLRequest,
		encodeGraphQLResponse,
	)

}

func (service *transportService) readinessCheckHandler(ctx *atreugo.RequestCtx) error {
	ctx.Response.SetStatusCode(http.StatusOK)

	return nil
}

func (service *transportService) livenessCheckHandler(ctx *atreugo.RequestCtx) error {
	ctx.Response.SetStatusCode(http.StatusOK)

	return nil
}
