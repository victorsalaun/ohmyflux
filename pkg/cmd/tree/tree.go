package tree

import (
	"fmt"
	"github.com/disiqueira/gotree"
	"github.com/spf13/cobra"
	"github.com/victorsalaun/ohmyflux/internal/api"
	"github.com/victorsalaun/ohmyflux/pkg/cmdutil"
)

type TreeOptions struct {
	Namespace string
}

func NewCmdTree(f *cmdutil.Factory, runF func(options *TreeOptions) error) *cobra.Command {
	opts := &TreeOptions{}
	cmd := &cobra.Command{
		Use:    "tree",
		Hidden: false,
		RunE: func(cmd *cobra.Command, args []string) error {
			return treeRun(f, opts)
		},
	}

	cmd.Flags().StringVarP(&opts.Namespace, "namespace", "n", "", "")

	return cmd
}

func treeRun(f *cmdutil.Factory, opts *TreeOptions) error {
	rootTree := gotree.New("tree")
	err := treeNamespaceRun(f, opts, rootTree)
	if err != nil {
		return err
	}
	fmt.Println("")
	fmt.Println(rootTree.Print())
	return nil
}

func treeNamespaceRun(f *cmdutil.Factory, opts *TreeOptions, tree gotree.Tree) error {
	if opts.Namespace != "" {
		namespaceTree := gotree.New("Namespace/" + opts.Namespace)
		err := treeHelmReleaseRun(f, opts, namespaceTree, opts.Namespace)
		if err != nil {
			return err
		}
		//err = treeSecretsRun(f, opts, namespaceTree, opts.Namespace)
		//if err != nil {
		//	return err
		//}
		tree.AddTree(namespaceTree)
		return nil
	} else {
		namespaces, err := api.GetNamespaces(f.KubernetesClient)
		if err != nil {
			return err
		}
		for _, namespace := range namespaces.Items {
			namespaceTree := gotree.New("Namespace/" + namespace.Name)
			//err := treeSecretsRun(f, opts, namespaceTree, namespace.Name)
			//if err != nil {
			//	return err
			//}
			tree.AddTree(namespaceTree)
		}
		return nil
	}
}

func treeHelmReleaseRun(f *cmdutil.Factory, opts *TreeOptions, tree gotree.Tree, namespace string) error {
	helmReleases, err := api.GetHelmReleases(f.KubernetesDynamicClient, namespace)
	if err != nil {
		return err
	}
	for _, helmRelease := range helmReleases {
		helmReleaseTree := gotree.New("HelmRelease/" + helmRelease.Name)
		err := treeConfigMapsRun(f, opts, helmReleaseTree, opts.Namespace, helmRelease.Name)
		if err != nil {
			return err
		}
		err = treeDeploymentsRun(f, opts, helmReleaseTree, opts.Namespace, helmRelease.Name)
		if err != nil {
			return err
		}
		err = treeEndpointsRun(f, opts, helmReleaseTree, opts.Namespace, helmRelease.Name)
		if err != nil {
			return err
		}
		err = treeIngressesRun(f, opts, helmReleaseTree, opts.Namespace, helmRelease.Name)
		if err != nil {
			return err
		}
		err = treePodsRun(f, opts, helmReleaseTree, opts.Namespace, helmRelease.Name)
		if err != nil {
			return err
		}
		err = treeSecretsRun(f, opts, helmReleaseTree, opts.Namespace, helmRelease.Name)
		if err != nil {
			return err
		}
		err = treeServiceAccountsRun(f, opts, helmReleaseTree, opts.Namespace, helmRelease.Name)
		if err != nil {
			return err
		}
		err = treeServicesRun(f, opts, helmReleaseTree, opts.Namespace, helmRelease.Name)
		if err != nil {
			return err
		}
		tree.AddTree(helmReleaseTree)
		fmt.Print(".")
	}
	return nil
}

func treeConfigMapsRun(f *cmdutil.Factory, opts *TreeOptions, tree gotree.Tree, namespace string, label string) error {
	configMaps, err := api.GetConfigMaps(f.KubernetesClient, namespace, label)
	if err != nil {
		return err
	}
	for _, configMap := range configMaps.Items {
		tree.Add("ConfigMap/" + configMap.Name)
	}
	return nil
}

func treeDeploymentsRun(f *cmdutil.Factory, opts *TreeOptions, tree gotree.Tree, namespace string, label string) error {
	deployments, err := api.GetDeployments(f.KubernetesClient, namespace, label)
	if err != nil {
		return err
	}
	for _, deployment := range deployments.Items {
		deploymentTree := gotree.New("Deployment/" + deployment.Name)
		err := treeReplicaSetsRun(f, opts, deploymentTree, opts.Namespace, label)
		if err != nil {
			return err
		}
		tree.AddTree(deploymentTree)
	}
	return nil
}

func treeEndpointsRun(f *cmdutil.Factory, opts *TreeOptions, tree gotree.Tree, namespace string, label string) error {
	endpoints, err := api.GetEndpoints(f.KubernetesClient, namespace, label)
	if err != nil {
		return err
	}
	for _, endpoint := range endpoints.Items {
		tree.Add("Endpoint/" + endpoint.Name)
	}
	return nil
}

func treeIngressesRun(f *cmdutil.Factory, opts *TreeOptions, tree gotree.Tree, namespace string, label string) error {
	ingresses, err := api.GetIngresses(f.KubernetesClient, namespace, label)
	if err != nil {
		return err
	}
	for _, ingress := range ingresses.Items {
		tree.Add("Ingress/" + ingress.Name)
	}
	return nil
}

func treePodsRun(f *cmdutil.Factory, opts *TreeOptions, tree gotree.Tree, namespace string, label string) error {
	pods, err := api.GetPods(f.KubernetesClient, namespace, label)
	if err != nil {
		return err
	}
	for _, pod := range pods.Items {
		tree.Add("Pod/" + pod.Name)
	}
	return nil
}

func treeReplicaSetsRun(f *cmdutil.Factory, opts *TreeOptions, tree gotree.Tree, namespace string, label string) error {
	replicaSets, err := api.GetReplicaSets(f.KubernetesClient, namespace, label)
	if err != nil {
		return err
	}
	for _, replicaSet := range replicaSets.Items {
		tree.Add("ReplicaSet/" + replicaSet.Name)
	}
	return nil
}

func treeSecretsRun(f *cmdutil.Factory, opts *TreeOptions, tree gotree.Tree, namespace string, label string) error {
	secrets, err := api.GetSecrets(f.KubernetesClient, namespace, label)
	if err != nil {
		return err
	}
	for _, secret := range secrets.Items {
		tree.Add("Secret/" + secret.Name)
	}
	return nil
}

func treeServiceAccountsRun(f *cmdutil.Factory, opts *TreeOptions, tree gotree.Tree, namespace string, label string) error {
	serviceAccounts, err := api.GetServiceAccounts(f.KubernetesClient, namespace, label)
	if err != nil {
		return err
	}
	for _, serviceAccount := range serviceAccounts.Items {
		tree.Add("ServiceAccount/" + serviceAccount.Name)
	}
	return nil
}

func treeServicesRun(f *cmdutil.Factory, opts *TreeOptions, tree gotree.Tree, namespace string, label string) error {
	services, err := api.GetServices(f.KubernetesClient, namespace, label)
	if err != nil {
		return err
	}
	for _, service := range services.Items {
		tree.Add("Services/" + service.Name)
	}
	return nil
}
