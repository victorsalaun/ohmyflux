package kube

import (
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var (
	kubernetesClient              *kubernetes.Clientset
	kubernetesApiExtensionsClient *clientset.Clientset
	kubernetesDynamicClient       dynamic.Interface
)

func GetKubernetesClient(config *rest.Config) *kubernetes.Clientset {
	if kubernetesClient == nil {
		client, err := kubernetes.NewForConfig(config)
		if err != nil {
			panic(err)
		}
		kubernetesClient = client
	}
	return kubernetesClient
}

func GetKubernetesApiExtensionsClient(config *rest.Config) *clientset.Clientset {
	if kubernetesApiExtensionsClient == nil {
		client, err := clientset.NewForConfig(config)
		if err != nil {
			panic(err)
		}
		kubernetesApiExtensionsClient = client
	}
	return kubernetesApiExtensionsClient
}

func GetKubernetesDynamicClient(config *rest.Config) dynamic.Interface {
	if kubernetesDynamicClient == nil {
		client, err := dynamic.NewForConfig(config)
		if err != nil {
			panic(err)
		}
		kubernetesDynamicClient = client
	}
	return kubernetesDynamicClient
}
