package cmd

import (
	"github.com/spf13/cobra"
	_	"github.com/spf13/pflag"
	"fmt"
	"io"
)

const (
	// Kcem component name
	componentKcem = "kcem"
)

func NewKcemCommand(in io.Reader, out, err io.Writer) *cobra.Command {

	cmd := &cobra.Command{
		Use: componentKcem,
		Short: "Kubernetes-cem",
		Long: `Provide kubernetes cluster information remotely.
It seems like kubectl.
`,
		DisableFlagParsing: false,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Starting ...")
		},

	}

	cmd.AddCommand(NewCmdInit(out))
	cmd.AddCommand(NewCmdStatus(out))

	return cmd
}