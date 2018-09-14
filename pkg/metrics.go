package pkg

import (
	"k8s.io/apimachinery/pkg/api/resource"
	v1beta1 "k8s.io/metrics/pkg/apis/metrics/v1beta1"
)

var m = new(resource.Quantity)

// ProcessServiceWiseMetrics ...
func processServiceWiseMetrics(pml *v1beta1.PodMetricsList) *Metric {
	cpu := new(resource.Quantity)
	mem := new(resource.Quantity)

	for _, pm := range pml.Items {
		for _, c := range pm.Containers {
			uCPU := c.Usage.Cpu()
			if uCPU != nil {
				cpu.Add(*uCPU)
			}
			uMem := c.Usage.Memory()
			if uMem != nil {
				mem.Add(*uMem)
				m.Add(*uMem)
			}
		}
	}

	return &Metric{
		cpu: cpu,
		mem: mem,
	}
}
