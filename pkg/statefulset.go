package pkg

import (
	apps "k8s.io/api/apps/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	metricsclientset "k8s.io/metrics/pkg/client/clientset/versioned"
)

// getStatefulset ...
func getStatefulset(client kubernetes.Interface) (*apps.StatefulSetList, error) {
	return client.AppsV1().StatefulSets(meta.NamespaceAll).List(meta.ListOptions{
		LabelSelector: labels.Everything().String(),
	})
}

// GetStatefulsetMetrics ...
func GetStatefulsetMetrics(config *rest.Config) (NamespaceWiseServiceMetrics, error) {
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	statefulsets, err := getStatefulset(client)
	if err != nil {
		return nil, err
	}

	// var pml metrics.PodMetricsList
	mClient, err := metricsclientset.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	metrics := make(NamespaceWiseServiceMetrics)

	for _, deploy := range statefulsets.Items {
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
