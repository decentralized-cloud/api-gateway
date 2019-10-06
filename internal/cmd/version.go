// Package cmd implements different commands that can be executed against API Gateway service
package cmd

import (
	"github.com/micro-business/go-core/pkg/util"
	"github.com/spf13/cobra"
)

func newVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Get API Gateway CLI version",
		Run: func(cmd *cobra.Command, args []string) {
			util.PrintInfo("API Gateway CLI\n")
			util.PrintInfo("Copyright (C) 2019, Micro Business Ltd.\n")
			util.PrintYAML(util.GetVersion())
		},
	}
}
