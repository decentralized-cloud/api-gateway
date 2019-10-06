package main

import (
	"github.com/decentralized-cloud/api-gateway/internal/cmd"
	"github.com/micro-business/go-core/pkg/util"
)

func main() {
	rootCmd := cmd.NewRootCommand()
	util.PrintIfError(rootCmd.Execute())
}
