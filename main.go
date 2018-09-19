package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aerokite/kubestat/pkg"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

var tsvEnabled bool

func init() {
	flag.BoolVar(&tsvEnabled, "tsv", false, "Display tab separated value")
	flag.Parse()
}

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
	if tsvEnabled {
		pkg.ShowTSV("Deployment", pml, tMetric)
	} else {
		pkg.Show("Deployment", pml, tMetric)
		fmt.Println()
		fmt.Println()
	}

	// fmt.Println("===== Daemonset =====")
	pml, err = client.GetDaemonsetMetrics()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if tsvEnabled {
		pkg.ShowTSV("DaemonSet", pml, tMetric)
	} else {
		pkg.Show("DaemonSet", pml, tMetric)
		fmt.Println()
		fmt.Println()
	}

	// fmt.Println("===== StatefulSet =====")
	pml, err = client.GetStatefulsetMetrics()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if tsvEnabled {
		pkg.ShowTSV("StatefulSet", pml, tMetric)
	} else {
		pkg.Show("StatefulSet", pml, tMetric)
	}
}
