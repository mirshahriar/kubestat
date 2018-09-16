package pkg

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"k8s.io/apimachinery/pkg/api/resource"
)

// Show ...
func Show(resourceName string, sm NamespaceWiseServiceMetrics) {
	table := tablewriter.NewWriter(os.Stdout)

	fmt.Printf("========== %s ==========\n", resourceName)

	table.SetCenterSeparator("+")
	table.SetColumnSeparator("|")
	table.SetRowSeparator("-")

	table.SetHeader([]string{"#", "Namespace", "Name", "CPU", "Memory", "Pods", "Containers"})

	//table.SetHeaderColor(tablewriter.BgRedColor)
	table.SetColumnAlignment([]int{
		tablewriter.ALIGN_CENTER,
		tablewriter.ALIGN_LEFT,
		tablewriter.ALIGN_LEFT,
		tablewriter.ALIGN_RIGHT,
		tablewriter.ALIGN_RIGHT,
		tablewriter.ALIGN_CENTER,
		tablewriter.ALIGN_CENTER,
	})

	cpu := new(resource.Quantity)
	mem := new(resource.Quantity)

	pods := 0
	containers := 0

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

			pods = pods + metric.pod
			containers = containers + metric.container

			//mem.SetScaled(mem.Value(), resource.Mega)
			row++
			table.Append([]string{
				fmt.Sprintf("%v", row),
				ns,
				svc,
				uCPU.String(),
				fmt.Sprintf("%vMi", uMem.Value()/(1024*1024)),
				fmt.Sprintf("%v", metric.pod),
				fmt.Sprintf("%v", metric.container),
			})
		}

	}

	table.SetFooter([]string{
		"", "", "",
		cpu.String(),
		fmt.Sprintf("%vMi", mem.Value()/(1024*1024)),
		fmt.Sprintf("%v", pods),
		fmt.Sprintf("%v", containers),
	})
	table.Render() // Send output

	//sMem, _ := mem.AsScale(resource.Giga)
	//fmt.Printf("CPU: %s\n", cpu.String())
	//fmt.Printf("MEM: %vMi\n", mem.Value()/(1024*1024))
}
