package options

import (
)

const defaultRootDir = "/var/lib/kubelet"

type KcemFlags struct {
	KubeConfig 	string		`json:"kubeConfig"`
	LogLevel   	string		`json:"logLevel"`
}

type Status struct {
	// NameSpace is the namespace name to query.
	NameSpace 	string 		`json:"nameSpace"`
	PodName 	string 		`json:"podName"`
}