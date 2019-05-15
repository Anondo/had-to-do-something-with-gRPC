package cmd

import (
	"emprpc/client"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	clientCmd = &cobra.Command{
		Use:   "client",
		Short: "Makes a request to the gRPC server",
		Run:   request,
	}
)

func init() {
	clientCmd.Flags().IntP("id", "i", 0, "The employee id")
	viper.BindPFlag("id", clientCmd.Flags().Lookup("id"))
}

func request(cmd *cobra.Command, args []string) {
	if err := client.MakeRequest(); err != nil {
		log.Fatal("gRPC client failed: ", err.Error())
	}
}
