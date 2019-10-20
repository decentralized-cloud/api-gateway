// Package cmd implements different commands that can be executed against API GatewaAPI Gateway service
package cmd

import (
	"github.com/decentralized-cloud/api-gateway/pkg/util"
	gocoreUtil "github.com/micro-business/go-core/pkg/util"
	"github.com/spf13/cobra"
)

func newStartCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start the API Gateway service",
		Run: func(cmd *cobra.Command, args []string) {
			gocoreUtil.PrintInfo("Copyright (C) 2019, Micro Business Ltd.\n")
			gocoreUtil.PrintYAML(gocoreUtil.GetVersion())
			util.StartService()
		},
	}
}
