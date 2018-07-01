package cmd

import (
	"github.com/spf13/cobra"
	"io"
	"github.com/cloud-jeong/go-startup/cmd/kcem/app/cmd/options"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/kubernetes"
)

// NewCmdCreate returns "kcem create" command.
func NewCmdCreate(out io.Writer) *cobra.Command {

	queryOptions := &options.Status{}

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Run this command to create k8s Deployment",
		Run: func(cmd *cobra.Command, args []string) {

			config, err := clientcmd.BuildConfigFromFlags("", "/Users/cloud/.kube/acorn_kube_config")
			if err != nil {
				panic(err.Error())
			}

			clientset, err := kubernetes.NewForConfig(config)
			if err != nil {
				panic(err.Error())
			}

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

