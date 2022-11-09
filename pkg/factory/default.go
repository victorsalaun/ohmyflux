package factory

import (
	"github.com/victorsalaun/ohmyflux/pkg/cmdutil"
	"github.com/victorsalaun/ohmyflux/pkg/iostreams"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

func New(kubernetesClient *kubernetes.Clientset, kubernetesDynamicClient dynamic.Interface) *cmdutil.Factory {
	f := &cmdutil.Factory{
		KubernetesClient:        kubernetesClient,
		KubernetesDynamicClient: kubernetesDynamicClient,
		ExecutableName:          "ohmyflux",
	}

	f.IOStreams = ioStreams(f) // Depends on Config
	return f
}

func ioStreams(f *cmdutil.Factory) *iostreams.IOStreams {
	io := iostreams.System()
	return io
}
