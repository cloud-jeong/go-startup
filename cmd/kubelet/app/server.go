package app

import (
	"github.com/spf13/cobra"
	_	"github.com/spf13/pflag"
	"fmt"
)

const (
	// Kubelet component name
	componentKubelet = "kubelet"
)

func NewKubeletCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use: componentKubelet,
		Long: `The kubelet is the primary "node agent" that runs on each
node. The kubelet works in terms of a PodSpec. A PodSpec is a YAML or JSON object
that describes a pod. The kubelet takes a set of PodSpecs that are provided through
various mechanisms (primary through the apiserver) and ensures that the containers
described in those PodSpecs are running and healthy. The kubelet doesn't manage
containers which were not created by kubernetes.

`,
		DisableFlagParsing: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Starting KubeletCommand ...")
		},

	}

	return cmd
}