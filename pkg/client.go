package pkg

import (
	"fmt"

	"k8s.io/client-go/kubernetes"
	metricsclientset "k8s.io/metrics/pkg/client/clientset/versioned"
)

// NewClient ...
func NewClient() (*Client, error) {
	config, err := buildConfigFromFlags("")
	if err != nil {
		return nil, fmt.Errorf("Failed to build config. Error: %v", err)
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("Failed to create client. Error: %v", err)
	}

	// var pml metrics.PodMetricsList
	mClient, err := metricsclientset.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("Failed to create client. Error: %v", err)
	}

	return &Client{
		metricsclient:    mClient,
		kubernetesclient: client,
	}, nil
}
