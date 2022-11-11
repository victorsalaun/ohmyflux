package api

import (
	"context"
	"github.com/fluxcd/helm-controller/api/v2beta1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	v1beta "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

const LabelSelectorKey = "helm.toolkit.fluxcd.io/name="

func GetNamespaces(client *kubernetes.Clientset) (*corev1.NamespaceList, error) {
	return client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
}

func GetHelmReleases(client dynamic.Interface, namespace string) ([]*v2beta1.HelmRelease, error) {
	var helmReleases []*v2beta1.HelmRelease
	resourceId := schema.GroupVersionResource{
		Group:    "helm.toolkit.fluxcd.io",
		Version:  "v2beta1",
		Resource: "helmreleases",
	}
	objects, err := client.Resource(resourceId).Namespace(namespace).List(context.TODO(), metav1.ListOptions{})
	for _, object := range objects.Items {
		helmReleases = append(helmReleases, &v2beta1.HelmRelease{
			ObjectMeta: metav1.ObjectMeta{
				Name: object.GetName(),
			},
		})
	}
	return helmReleases, err
}

func GetConfigMaps(kubernetesClient *kubernetes.Clientset, namespace string, label string) (*corev1.ConfigMapList, error) {
	return kubernetesClient.CoreV1().ConfigMaps(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: LabelSelectorKey + label,
	})
}

func GetDeployments(kubernetesClient *kubernetes.Clientset, namespace string, label string) (*appsv1.DeploymentList, error) {
	return kubernetesClient.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: LabelSelectorKey + label,
	})
}

func GetEndpoints(kubernetesClient *kubernetes.Clientset, namespace string, label string) (*corev1.EndpointsList, error) {
	return kubernetesClient.CoreV1().Endpoints(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: LabelSelectorKey + label,
	})
}

func GetIngresses(kubernetesClient *kubernetes.Clientset, namespace string, label string) (*v1beta.IngressList, error) {
	return kubernetesClient.ExtensionsV1beta1().Ingresses(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: LabelSelectorKey + label,
	})
}

func GetPods(kubernetesClient *kubernetes.Clientset, namespace string, label string) (*corev1.PodList, error) {
	return kubernetesClient.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: LabelSelectorKey + label,
	})
}

func GetReplicaSets(kubernetesClient *kubernetes.Clientset, namespace string, label string) (*appsv1.ReplicaSetList, error) {
	return kubernetesClient.AppsV1().ReplicaSets(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: LabelSelectorKey + label,
	})
}

func GetSecrets(kubernetesClient *kubernetes.Clientset, namespace string, label string) (*corev1.SecretList, error) {
	return kubernetesClient.CoreV1().Secrets(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: LabelSelectorKey + label,
	})
}

func GetServiceAccounts(kubernetesClient *kubernetes.Clientset, namespace string, label string) (*corev1.ServiceAccountList, error) {
	return kubernetesClient.CoreV1().ServiceAccounts(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: LabelSelectorKey + label,
	})
}

func GetServices(kubernetesClient *kubernetes.Clientset, namespace string, label string) (*corev1.ServiceList, error) {
	return kubernetesClient.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: LabelSelectorKey + label,
	})
}
