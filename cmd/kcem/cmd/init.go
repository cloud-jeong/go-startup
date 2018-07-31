package cmd

import (
	"github.com/spf13/cobra"
	"io"
)

// NewCmdInit returns "kcem init" command.
func NewCmdInit(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Run this command in order to set up the Kubernetes master.",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}