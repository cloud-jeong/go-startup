package cmd

import (
	"github.com/spf13/cobra"
	"io"
	"fmt"
	"github.com/spf13/pflag"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/kubernetes"
)

const statusHelp = `
This command creates a chart directory along with the common files and
directories used in a chart.

For example, 'helm create foo' will create a directory structure that looks
something like this:

	foo/
	  |
	  |- .helmignore   # Contains patterns to ignore when packaging Helm charts.
	  |
	  |- Chart.yaml    # Information about your chart
	  |
	  |- values.yaml   # The default values for your templates
	  |
	  |- charts/       # Charts that this chart depends on
	  |
	  |- templates/    # The template files

'helm create' takes a path for an argument. If directories in the given path
do not exist, Helm will attempt to create them as it goes. If the given
destination exists and there are files in that directory, conflicting files
will be overwritten, but other files will be left alone.
`

type statusCmd struct {
	namespace 	string 		`json:"nameSpace"`
	podName 	string 		`json:"podName"`
	out			io.Writer
}

func init() {
	fmt.Println("Called init() ..")
}

// NewCmdStatus returns "kcem status" command.
func NewCmdStatus(out io.Writer) *cobra.Command {
	status := &statusCmd{
		out:	out,
	}

	cmd := &cobra.Command{
		Use:   	"status",
		Short: 	"Run this command to get current cluster status",
		Long: 	statusHelp,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("PreRunE() callled")
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error{
			cmd.Flags().Visit(func(f *pflag.Flag) {
				fmt.Printf("option [%s][%v]\n", f.Name, f.Value)
			})

			return status.run()
		},
	}

	f := cmd.Flags()
	f.StringVarP(&status.namespace, "namespace", "n", "--all-namespaces", "Namespace name to query")
	f.StringVarP(&status.podName, "podname", "p", "", "Pod name to query")

	return cmd
}

func (s *statusCmd) run() error {
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/cloud/.kube/onap_kube_config")
	if err != nil {
		return err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	pods, err := clientset.CoreV1().Pods("kube-system").List(metav1.ListOptions{})
	if err != nil {
		return err
	}
	fmt.Printf("kube-system %d pods running\n", len(pods.Items))

	pods, err = clientset.CoreV1().Pods("cocktail-system").List(metav1.ListOptions{})
	if err != nil {
		return err
	}
	fmt.Printf("cocktail-system %d pods running\n", len(pods.Items))

	pods, err = clientset.CoreV1().Pods("monitoring").List(metav1.ListOptions{})
	if err != nil {
		return err
	}
	fmt.Printf("monitoring %d pods running\n", len(pods.Items))

	pods, err = clientset.CoreV1().Pods("onap").List(metav1.ListOptions{})
	if err != nil {
		return err
	}
	fmt.Printf("onap %d pods running\n", len(pods.Items))


	return nil
}

func int32Ptr(i int32) *int32 { return &i }