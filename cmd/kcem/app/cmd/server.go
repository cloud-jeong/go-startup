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
		Long: `The kubelet is the primary "node agent" that runs on each
node. The kubelet works in terms of a PodSpec. A PodSpec is a YAML or JSON object
that describes a pod. The kubelet takes a set of PodSpecs that are provided through
various mechanisms (primary through the apiserver) and ensures that the containers
described in those PodSpecs are running and healthy. The kubelet doesn't manage
containers which were not created by kubernetes.

`,
		DisableFlagParsing: false,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Starting ...")
		},

	}

	cmd.AddCommand(NewCmdInit(out))

	return cmd
}