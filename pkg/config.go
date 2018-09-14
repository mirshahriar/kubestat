package pkg

import (
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func BuildConfigFromFlags(contextName string) (*rest.Config, error) {
	var loader clientcmd.ClientConfigLoader
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	rules.DefaultClientConfig = &clientcmd.DefaultClientConfig
	loader = rules

	overrides := &clientcmd.ConfigOverrides{
		CurrentContext: contextName,
	}
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loader, overrides).ClientConfig()
}
