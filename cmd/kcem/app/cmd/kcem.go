package cmd

import (
	"github.com/spf13/cobra"
	_	"github.com/spf13/pflag"
	"fmt"
	"io"
	"github.com/cloud-jeong/go-startup/cmd/kcem/app/cmd/options"
	"os"
	"path/filepath"
)

var globalUsage = `The Kubernetes package manager

To begin working with Helm, run the 'helm init' command:

	$ helm init

This will install Tiller to your running Kubernetes cluster.
It will also set up any necessary local configuration.

Common actions from this point include:

- helm search:    search for charts
- helm fetch:     download a chart to your local directory to view
- helm install:   upload the chart to Kubernetes
- helm list:      list releases of charts

Environment:
  $HELM_HOME          set an alternative location for Helm files. By default, these are stored in ~/.helm
  $HELM_HOST          set an alternative Tiller host. The format is host:port
  $HELM_NO_PLUGINS    disable plugins. Set HELM_NO_PLUGINS=1 to disable plugins.
  $TILLER_NAMESPACE   set an alternative Tiller namespace (default "kube-system")
  $KUBECONFIG         set an alternative Kubernetes configuration file (default "~/.kube/config")
`

const (
	// Kcem component name
	componentKcem = "kcem"
)

func NewKcemCommand(in io.Reader, out, err io.Writer) *cobra.Command {

	kcemOption := &options.KcemFlags{}

	rootCmd := &cobra.Command{
		Use: componentKcem,
		Short: "Kubernetes-cem",
		Long: globalUsage,
		DisableFlagParsing: false,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Starting ...")
		},

	}

	rootCmd.PersistentFlags().StringVar(&kcemOption.KubeConfig, "kubeconfig", filepath.Join(homeDir(), ".kube", "config"), "absolute path to the kubeconfig file")
	rootCmd.PersistentFlags().StringVar(&kcemOption.LogLevel, "loglevel", "", "Log level")

	rootCmd.AddCommand(
		NewCmdInit(out),
		NewCmdStatus(out),
		NewCmdCreate(out),
	)

	return rootCmd
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}