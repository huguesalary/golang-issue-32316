package main

import (
	_ "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/minikube/pkg/minikube/config"
	_ "k8s.io/minikube/pkg/minikube/constants"

	_ "k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	_ "k8s.io/client-go/rest"
)

func main() {

}
