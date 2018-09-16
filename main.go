package main

import (
	"fmt"

	"github.com/aerokite/kubestat/pkg"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func main() {

	// uses the current context in kubeconfig
	config, err := pkg.BuildConfigFromFlags("")
	if err != nil {
		panic(err)
	}

	//fmt.Println("===== Deployment =====")
	pml, err := pkg.GetDeploymetMetrics(config)
	if err != nil {
		panic(err)
	}
	pkg.Show("Deployment", pml)

	fmt.Println()
	fmt.Println()

	// fmt.Println("===== Daemonset =====")
	pml, err = pkg.GetDaemonsetMetrics(config)
	if err != nil {
		panic(err)
	}
	pkg.Show("Daemonset", pml)

	fmt.Println()
	fmt.Println()

	// fmt.Println("===== StatefulSet =====")
	pml, err = pkg.GetStatefulsetMetrics(config)
	if err != nil {
		panic(err)
	}
	pkg.Show("StatefulSet", pml)
}
