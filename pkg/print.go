package pkg

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"k8s.io/apimachinery/pkg/api/resource"
)

// Show ...
func Show(sm NamespaceWiseServiceMetrics) {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"#", "Namespace", "Name", "CPU", "Memory"})
	table.SetColumnAlignment([]int{
		tablewriter.ALIGN_CENTER,
		tablewriter.ALIGN_LEFT,
		tablewriter.ALIGN_LEFT,
		tablewriter.ALIGN_RIGHT,
		tablewriter.ALIGN_RIGHT,
	})

	cpu := new(resource.Quantity)
	mem := new(resource.Quantity)

	row := 0
	for ns, val := range sm {
		for svc, metric := range val {

			uCPU := metric.cpu
			if uCPU != nil {
				cpu.Add(*uCPU)
			}
			uMem := metric.mem
			if uMem != nil {
				mem.Add(*uMem)
			}

			//mem.SetScaled(mem.Value(), resource.Mega)
			row++
			table.Append([]string{fmt.Sprintf("%v", row), ns, svc, uCPU.String(), fmt.Sprintf("%vMi", uMem.Value()/(1024*1024))})
		}

	}

	table.Render() // Send output

	//sMem, _ := mem.AsScale(resource.Giga)
	fmt.Printf("CPU: %s\n", cpu.String())
	fmt.Printf("MEM: %vMi\n", mem.Value()/(1024*1024))
}
