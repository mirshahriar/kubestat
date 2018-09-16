package pkg

import (
	"k8s.io/apimachinery/pkg/api/resource"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetTotalMetric ...
func (c *Client) GetTotalMetric() (*Metric, error) {
	nodeMegtrics, err := c.metricsclient.Metrics().NodeMetricses().List(meta.ListOptions{})
	if err != nil {
		return nil, err
	}

	cpu := new(resource.Quantity)
	mem := new(resource.Quantity)

	for _, nm := range nodeMegtrics.Items {
		uCPU := nm.Usage.Cpu()
		if uCPU != nil {
			cpu.Add(*uCPU)
		}
		uMem := nm.Usage.Memory()
		if uMem != nil {
			mem.Add(*uMem)
		}
	}

	return &Metric{
		cpu: *cpu,
		mem: *mem,
	}, nil
}
