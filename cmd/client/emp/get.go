package emp

import (
	"emprpc/client"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// empGetCmd is the get command to get the employee data
var (
	empGetCmd = &cobra.Command{
		Use:   "get",
		Short: "Fetches a employee info",
		Run:   getEmp,
	}
)

func init() {
	empGetCmd.Flags().IntP("id", "i", 0, "The employee id")
	viper.BindPFlag("id", empGetCmd.Flags().Lookup("id"))
}

func getEmp(cmd *cobra.Command, args []string) {
	id := viper.GetInt32("id")
	if err := client.FetchEmpRequest(id); err != nil {
		log.Fatal(err.Error())
	}
}
