package pkg

import (
	apps "k8s.io/api/apps/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	metricsclientset "k8s.io/metrics/pkg/client/clientset/versioned"
)

// getDeplotment ...
func getDeplotment(client kubernetes.Interface) (*apps.DeploymentList, error) {
	return client.AppsV1().Deployments(meta.NamespaceAll).List(meta.ListOptions{
		LabelSelector: labels.Everything().String(),
	})
}

// GetDeploymetMetrics ...
func GetDeploymetMetrics(config *rest.Config) (NamespaceWiseServiceMetrics, error) {
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	deployments, err := getDeplotment(client)
	if err != nil {
		return nil, err
	}

	// var pml metrics.PodMetricsList
	mClient, err := metricsclientset.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	metrics := make(NamespaceWiseServiceMetrics)

	for _, deploy := range deployments.Items {
		selector, err := meta.LabelSelectorAsSelector(deploy.Spec.Selector)
		if err != nil {
			return nil, err
		}
		podMetricsList, err := mClient.Metrics().PodMetricses(deploy.Namespace).List(meta.ListOptions{
			LabelSelector: selector.String(),
		})
		if err != nil {
			return nil, err
		}

		metric := processServiceWiseMetrics(podMetricsList)
		if _, found := metrics[deploy.Namespace]; !found {
			metrics[deploy.Namespace] = make(map[string]*Metric)
		}

		metrics[deploy.Namespace][deploy.Name] = metric

	}

	return metrics, nil
}
