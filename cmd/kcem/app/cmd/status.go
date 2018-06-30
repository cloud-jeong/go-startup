package cmd

import (
	"github.com/spf13/cobra"
	"io"
)

// NewCmdStatus returns "kcem status" command.
func NewCmdStatus(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Run this command to get current cluster status",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}