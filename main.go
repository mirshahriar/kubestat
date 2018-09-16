package main

import (
	"fmt"
	"os"

	"github.com/aerokite/kubestat/pkg"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func main() {

	client, err := pkg.NewClient()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	tMetric, err := client.GetTotalMetric()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//fmt.Println("===== Deployment =====")
	pml, err := client.GetDeploymetMetrics()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	pkg.Show("Deployment", pml, tMetric)

	fmt.Println()
	fmt.Println()

	// fmt.Println("===== Daemonset =====")
	pml, err = client.GetDaemonsetMetrics()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	pkg.Show("Daemonset", pml, tMetric)

	fmt.Println()
	fmt.Println()

	// fmt.Println("===== StatefulSet =====")
	pml, err = client.GetStatefulsetMetrics()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	pkg.Show("StatefulSet", pml, tMetric)
}
