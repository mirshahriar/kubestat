package pkg

import "k8s.io/apimachinery/pkg/api/resource"

// Metric ...
type Metric struct {
	cpu       *resource.Quantity
	mem       *resource.Quantity
	pod       int
	container int
}

// NamespaceWiseServiceMetrics ...
type NamespaceWiseServiceMetrics map[string](map[string]*Metric)

// ServiceToLabel ...
type ServiceToLabel map[string]string

var s2l ServiceToLabel

func init() {
	s2l = make(ServiceToLabel)
}
