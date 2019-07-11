package main

import (
	"flag"
	"fmt"

	"github.com/tigera/api/pkg/client/clientset_generated/clientset"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// Create a new config based on kubeconfig file.
	var kubeconfig *string
	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// Build a clientset based on the provided kubeconfig file.
	cs, err := clientset.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// List global report types.
	list, err := cs.ProjectcalicoV3().GlobalReportTypes("").List(v1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, grt := range list.Items {
		fmt.Printf("%#v\n", grt)
	}
}
