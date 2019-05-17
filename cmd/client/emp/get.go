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
	empGetCmd.Flags().Int32P("eid", "e", 0, "The employee id")
	viper.BindPFlag("eid", empGetCmd.Flags().Lookup("eid"))
}

func getEmp(cmd *cobra.Command, args []string) {
	id := viper.GetInt32("eid")
	if err := client.FetchEmpRequest(id); err != nil {
		log.Fatal(err.Error())
	}
}
