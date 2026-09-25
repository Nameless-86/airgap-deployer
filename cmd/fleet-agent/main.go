package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Nameless-86/airgap-deployer/pkg/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	client, err := k8s.NewKubeClient("")
	if err != nil {
		log.Fatalf("Error connecting to cluster: %v", err)
	}

	pods, err := client.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Failed to list pods: %v", err)
	}

	fmt.Printf("Successfully connected, pods running: %d \n", len(pods.Items))
}
