package cmd

import (
	"github.com/spf13/cobra"
	"io"
	"github.com/cloud-jeong/go-startup/cmd/kcem/app/cmd/options"
	"fmt"
	"github.com/spf13/pflag"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
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

			deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)


			deployment := &appsv1.Deployment{
				ObjectMeta: metav1.ObjectMeta{
					Name: "demo-deployment",
				},
				Spec: appsv1.DeploymentSpec{
					Replicas: int32Ptr(2),
					Selector: &metav1.LabelSelector{
						MatchLabels: map[string]string{
							"app": "demo",
						},
					},
					Template: apiv1.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Labels: map[string]string{
								"app": "demo",
							},
						},
						Spec: apiv1.PodSpec{
							Containers: []apiv1.Container{
								{
									Name:  "web",
									Image: "nginx:1.12",
									Ports: []apiv1.ContainerPort{
										{
											Name:          "http",
											Protocol:      apiv1.ProtocolTCP,
											ContainerPort: 80,
										},
									},
								},
							},
						},
					},
				},
			}

			// Create Deployment
			fmt.Println("Creating deployment...")
			result, err := deploymentsClient.Create(deployment)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())

			},
	}

	cmd.Flags().StringVarP(&queryOptions.NameSpace, "namespace", "n", "--all-namespaces", "Namespace name to query")
	cmd.Flags().StringVarP(&queryOptions.PodName, "podname", "p", "", "Pod name to query")

	return cmd
}

func int32Ptr(i int32) *int32 { return &i }