package cmd

import (
	"github.com/spf13/cobra"
	"io"
	"github.com/cloud-jeong/go-startup/cmd/kcem/app/cmd/options"
	"fmt"
	"github.com/spf13/pflag"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/kubernetes"
)

func init() {
//	fmt.Println("Called init() ..")
}

// NewCmdStatus returns "kcem status" command.
func NewCmdStatus(out io.Writer) *cobra.Command {

	queryOptions := &options.Status{}

	cmd := &cobra.Command{
		Use:   "status",
		Short: "Run this command to get current cluster status",
		Run: func(cmd *cobra.Command, args []string) {

			cmd.Flags().Visit(func(f *pflag.Flag) {
				fmt.Printf("option [%s][%v]\n", f.Name, f.Value)
			})

			config, err := clientcmd.BuildConfigFromFlags("", "/Users/cloud/.kube/acorn_kube_config")
			if err != nil {
				panic(err.Error())
			}

			clientset, err := kubernetes.NewForConfig(config)
			if err != nil {
				panic(err.Error())
			}

			pods, err := clientset.CoreV1().Pods("kube-system").List(metav1.ListOptions{})
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("kube-system %d pods running\n", len(pods.Items))

			pods, err = clientset.CoreV1().Pods("cocktail-system").List(metav1.ListOptions{})
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("cocktail-system %d pods running\n", len(pods.Items))

			pods, err = clientset.CoreV1().Pods("monitoring").List(metav1.ListOptions{})
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("monitoring %d pods running\n", len(pods.Items))

		},
	}

	cmd.Flags().StringVarP(&queryOptions.NameSpace, "namespace", "n", "--all-namespaces", "Namespace name to query")
	cmd.Flags().StringVarP(&queryOptions.PodName, "podname", "p", "", "Pod name to query")

	return cmd
}

func int32Ptr(i int32) *int32 { return &i }