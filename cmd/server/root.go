package cmd

import (
	"emprpc/server"
	"log"

	"github.com/spf13/cobra"
)

// ServeCmd is the cobra Command instance for the server command
var (
	ServeCmd = &cobra.Command{
		Use:   "server",
		Short: "Starts the gRPC server",
		Run:   serve,
	}
)

func serve(cmd *cobra.Command, args []string) {
	if err := server.StartServer(); err != nil {
		log.Fatal("gRPC server failed: ", err.Error())
	}
}
