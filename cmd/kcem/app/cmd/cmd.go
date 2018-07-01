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

const (
	// Kcem component name
	componentKcem = "kcem"
)

func NewKcemCommand(in io.Reader, out, err io.Writer) *cobra.Command {

	kcemOption := &options.KcemFlags{}

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

	cmd.PersistentFlags().StringVar(&kcemOption.KubeConfig, "kubeconfig", filepath.Join(homeDir(), ".kube", "config"), "absolute path to the kubeconfig file")
	cmd.PersistentFlags().StringVar(&kcemOption.LogLevel, "loglevel", "", "Log level")

	cmd.AddCommand(NewCmdInit(out))
	cmd.AddCommand(NewCmdStatus(out))

	return cmd
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}