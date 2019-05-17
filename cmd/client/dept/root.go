package dept

import "github.com/spf13/cobra"

// The Department command
var (
	DeptCmd = &cobra.Command{
		Use:   "dept",
		Short: "Everything related to departments",
	}
)

func init() {
	DeptCmd.AddCommand(deptGetCmd)
	DeptCmd.AddCommand(setDeptCmd)
}
