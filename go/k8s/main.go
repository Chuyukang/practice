/*
Copyright 2016 The Kubernetes Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	metricsv "k8s.io/metrics/pkg/client/clientset/versioned"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

// tutorial link: https://github.com/kubernetes/client-go/tree/master/examples/in-cluster-client-configuration

func main() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	var metricCS *metricsv.Clientset
	metricCS, err = metricsv.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	for {
		// get pods in all the namespaces by omitting namespace
		// Or specify namespace to get pods in particular namespace
		pods, err := clientset.CoreV1().Pods("openfaas").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		// Examples for error handling:
		// - Use helper functions e.g. errors.IsNotFound()
		// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
		_, err = clientset.CoreV1().Pods("default").Get(context.TODO(), "example-xxxxx", metav1.GetOptions{})
		if errors.IsNotFound(err) {
			fmt.Printf("Pod example-xxxxx not found in default namespace\n")
		} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
			fmt.Printf("Error getting pod %v\n", statusError.ErrStatus.Message)
		} else if err != nil {
			panic(err.Error())
		} else {
			fmt.Printf("Found example-xxxxx pod in default namespace\n")
		}

		// get available pods name by specifying function name
		namespace := "openfaas-fn"
		functionName := "show-backend"
		// label selector: faas_function, faas_function=<functionName>
		selector := labels.NewSelector()
		legalFunctionReq, _ := labels.NewRequirement("faas_function", selection.Exists, []string{})
		functionNameReq, _ := labels.NewRequirement("faas_function", selection.Equals, []string{functionName})
		selector = selector.Add(*legalFunctionReq)
		selector = selector.Add(*functionNameReq)
		listOptions := metav1.ListOptions{LabelSelector: selector.String()}
		var backends *v1.PodList
		backends, err = clientset.CoreV1().Pods(namespace).List(context.TODO(), listOptions)
		if err != nil {
			fmt.Println("Pod List error!")
		}
		for _, item := range backends.Items {
			podName := item.Name
			fmt.Printf("%s %s %s\n", podName, item.Status.Phase, item.Status.PodIP)
			// get metric info
			podMetric, _ := metricCS.MetricsV1beta1().PodMetricses(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
			// sum container metrics
			podMemory := &resource.Quantity{}
			podCPU := &resource.Quantity{}
			containersMetricsList := podMetric.Containers
			for _, containerMetrics := range containersMetricsList {
				podMemory.Add(containerMetrics.Usage[v1.ResourceMemory])
				podCPU.Add(containerMetrics.Usage[v1.ResourceCPU])
			}
			// show metric info
			fmt.Printf("%s %s %s\n", podName, podMemory, podCPU)
		}

		time.Sleep(10 * time.Second)
	}
}
