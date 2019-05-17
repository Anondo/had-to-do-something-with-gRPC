package dept

import (
	"emprpc/client"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// deptGetCmd is the get command to get the department data
var (
	deptGetCmd = &cobra.Command{
		Use:   "get",
		Short: "Fetches a department info",
		Run:   getDept,
	}
)

func init() {
	deptGetCmd.Flags().IntP("id", "i", 0, "The employee id")
	viper.BindPFlag("id", deptGetCmd.Flags().Lookup("id"))
}

func getDept(cmd *cobra.Command, args []string) {
	id := viper.GetInt32("id")
	if err := client.FetchDeptRequest(id); err != nil {
		log.Fatal(err.Error())
	}
}
