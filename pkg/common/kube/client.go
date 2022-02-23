package kube

import (
	"k8s.io/client-go/kubernetes"
)

// init new incluster kube config
func NewInClusterClientSet() (*kubernetes.Clientset, error) {
	cfg, err := GetKubeClusterConfig()
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}
