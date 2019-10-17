// Package util implements different utilities required by the API Gateway service
package util

import (
	"log"
	"os"
	"os/signal"

	"github.com/decentralized-cloud/api-gateway/services/configuration"
	"github.com/decentralized-cloud/api-gateway/services/transport/graphql"
	"github.com/decentralized-cloud/api-gateway/services/transport/graphql/types"
	"go.uber.org/zap"
)

var configurationService configuration.ConfigurationContract
var resolverCreator types.ResolverCreatorContract

// StartService setups all dependecies required to start the API Gateway service and
// start the service
func StartService() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	defer logger.Sync()

	if err = setupDependencies(logger); err != nil {
		logger.Fatal("Failed to setup dependecies", zap.Error(err))
	}

	graphqlTransportService, err := graphql.NewTransportService(
		logger,
		configurationService,
		resolverCreator)
	if err != nil {
		logger.Fatal("Failed to create GraphQL transport service", zap.Error(err))
	}

	go graphqlTransportService.Start()

	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan struct{})
	signal.Notify(signalChan, os.Interrupt)

	go func() {
		<-signalChan
		logger.Info("Received an interrupt, stopping services...")

		if err := graphqlTransportService.Stop(); err != nil {
			logger.Error("Failed to stop GraphQL transport service", zap.Error(err))
		}

		close(cleanupDone)
	}()
	<-cleanupDone
}

func setupDependencies(logger *zap.Logger) (err error) {
	if configurationService, err = configuration.NewEnvConfigurationService(); err != nil {
		return
	}

	if resolverCreator, err = graphql.NewResolverCreator(
		logger,
		configurationService); err != nil {
		return
	}

	return
}
