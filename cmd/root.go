package cmd

import (
	"emprpc/db"
	"log"

	"github.com/spf13/cobra"
)

// RootCmd is the root command
var (
	RootCmd = &cobra.Command{
		Use:   "emprpc",
		Short: "emprpc is the desperation of doing something in gRPC",
	}
)

func init() {
	db.Init()
	RootCmd.AddCommand(serveCmd)
	RootCmd.AddCommand(clientCmd)
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
