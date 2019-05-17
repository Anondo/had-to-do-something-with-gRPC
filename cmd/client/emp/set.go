package emp

import (
	"emprpc/client"
	"log"

	"github.com/spf13/cobra"
)

// setEmpCmd the set emplpoyee command
var (
	setEmpCmd = &cobra.Command{
		Use:   "set",
		Short: "creates a employee",
		Run:   setEmp,
	}
)

func setEmp(cmd *cobra.Command, args []string) {
	if err := client.SetEmpRequest(); err != nil {
		log.Fatal(err.Error())
	}
}
