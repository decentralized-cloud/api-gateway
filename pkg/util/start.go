// Package util implements different utilities required by the tenant service
package util

import (
	"log"
	"os"
	"os/signal"

	"go.uber.org/zap"
)

// StartService setups all dependecies required to start the API Gateway service and
// start the service
func StartService() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	defer logger.Sync()

	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan struct{})
	signal.Notify(signalChan, os.Interrupt)

	go func() {
		<-signalChan
		logger.Info("Received an interrupt, stopping services...")

		close(cleanupDone)
	}()
	<-cleanupDone
}
