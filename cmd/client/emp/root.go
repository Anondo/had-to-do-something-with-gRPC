package emp

import "github.com/spf13/cobra"

// The employee command
var (
	EmpCmd = &cobra.Command{
		Use:   "emp",
		Short: "Everything related to employee",
	}
)

func init() {
	EmpCmd.AddCommand(empGetCmd)
	EmpCmd.AddCommand(setEmpCmd)
}
