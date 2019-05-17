package dept

import (
	"emprpc/client"
	"log"

	"github.com/spf13/cobra"
)

// setDeptCmd the set department command
var (
	setDeptCmd = &cobra.Command{
		Use:   "set",
		Short: "creates a departement",
		Run:   setDept,
	}
)

func setDept(cmd *cobra.Command, args []string) {
	if err := client.SetDeptRequest(); err != nil {
		log.Fatal(err.Error())
	}
}
