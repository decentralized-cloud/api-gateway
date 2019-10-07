// Package cmd implements different commands that can be executed against API GatewaAPI Gateway service
package cmd

import (
	"github.com/decentralized-cloud/api-gateway/pkg/util"
	"github.com/spf13/cobra"
)

func newStartCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start the API Gateway service",
		Run: func(cmd *cobra.Command, args []string) {
			util.StartService()
		},
	}
}
