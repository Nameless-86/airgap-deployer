package k8s

import (
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func NewKubeClient(customKubeconfig string) (kubernetes.Interface, error) {
	var kubeConfigPath string
	if customKubeconfig != "" {
		kubeConfigPath = customKubeconfig
	} else {
		k3sDefault := "/etc/rancher/k3s/k3s.yaml"
		if _, err := os.Stat(k3sDefault); err == nil {
			kubeConfigPath = k3sDefault
		} else {
			home, _ := os.UserHomeDir()
			kubeConfigPath = filepath.Join(home, ".kube", "config")
		}
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		return nil, fmt.Errorf("failed to build kubeconfig from path %s: %w", kubeConfigPath, err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create kubernetes clientset: %w", err)
	}

	return clientset, nil
}
