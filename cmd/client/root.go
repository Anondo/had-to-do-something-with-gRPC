package cmd

import (
	"emprpc/cmd/client/dept"
	"emprpc/cmd/client/emp"

	"github.com/spf13/cobra"
)

// The client command
var (
	ClientCmd = &cobra.Command{
		Use:   "client",
		Short: "Makes a request to the gRPC server",
	}
)

func init() {
	ClientCmd.AddCommand(dept.DeptCmd)
	ClientCmd.AddCommand(emp.EmpCmd)
}
