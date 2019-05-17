package cmd

import (
	clnt "emprpc/cmd/client"
	srvr "emprpc/cmd/server"
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
	RootCmd.AddCommand(srvr.ServeCmd)
	RootCmd.AddCommand(clnt.ClientCmd)
}

// Execute executes the root command
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
