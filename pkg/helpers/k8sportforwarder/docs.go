// Package k8sportforwarder e is responsible for forwarding ports from a pod to the local machine.
//
// It uses the k8s.io/client-go library to create a port forwarder that listens on a local port and forwards traffic to a pod in a Kubernetes cluster.
// The PortForwarder struct contains the necessary information to create a port forwarder, such as the Kubernetes configuration, clientset, pod, local port, and container port.
//
// Example:
//
//	package login
//
//	import (
//	  "fmt"
//	  "log"
//	  "os"
//	  "os/signal"
//	  "sync"
//	  "syscall"
//
//	  "github.com/NorskHelsenett/ror/pkg/helpers/k8sportforwarder"
//	  v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	)
//
//	func main() {
//		  //     Create a new PortForwarder instance
//		  forwarder := k8sportforwarder.NewPortForwarder()
//
//		  //     Add a pod to the port forwarder
//		  err := forwarder.AddPodByLabels(v1.LabelSelector{
//		  	MatchLabels: map[string]string{
//		  		"app.kubernetes.io/name":     "grafana",
//		  		"app.kubernetes.io/instance": "prometheus",
//		  	},
//		  }, "prometheus-operator")
//
//		  //     Handle error
//		  if err != nil {
//		  	log.Fatal(err)
//		  }
//
//		  //     Create a waitgroup to wait for the port forwarding to finish
//		  var wg sync.WaitGroup
//		  wg.Add(1)
//
//		  //     Create a channel to stop the port forwarding and a channel to signal when the port forwarding is ready
//		  stopCh := make(chan struct{}, 1)
//		  readyCh := make(chan struct{})
//
//		  //     Create a signal channel to handle SIGINT and SIGTERM signals
//		  sigs := make(chan os.Signal, 1)
//		  signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
//
//		  //     Wait for the SIGINT or SIGTERM signal
//		  go func() {
//		  	<-sigs
//		  	fmt.Println("Bye...")
//		  	close(stopCh)
//		  	wg.Done()
//		  }()
//
//		  //     Create a port forwarder instance
//		  go func() {
//		  	err := forwarder.Forward(readyCh, stopCh)
//		  	if err != nil {
//		  		panic(err)
//		  	}
//		  }()
//
//		  //     Wait for the port forwarding to be ready
//		  <-readyCh
//
//		  fmt.Println("Port forwarding is ready")
//
//		  //     Wait for the port forwarding to finish
//		  wg.Wait()
//	}
package k8sportforwarder
