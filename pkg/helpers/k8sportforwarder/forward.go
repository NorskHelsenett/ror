package k8sportforwarder

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"

	rorkubernetesclient "github.com/NorskHelsenett/ror/pkg/clients/kubernetes"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/portforward"
	"k8s.io/client-go/transport/spdy"
	k8sconfig "sigs.k8s.io/controller-runtime/pkg/client/config"
)

// PortForwarder is the request to port forward a pod
type PortForwarder struct {
	// RestConfig is the kubernetes config
	RestConfig *rest.Config
	// Clientset is the kubernetes clientset
	Clientset kubernetes.Interface
	// Pod is the selected pod for this port forwarding
	Pod *v1.Pod
	// LocalPort is the local port that will be selected to expose the PodPort
	LocalPort int32
	// ContainerPort is the target port for the pod
	ContainerPort int32
	// Steams configures where to write or read input from
	Streams genericclioptions.IOStreams
}

// NewPortForwarder creates a new PortForwarder instance
func NewPortForwarder() *PortForwarder {
	restconf, err := k8sconfig.GetConfig()
	cobra.CheckErr(err)

	streams := genericclioptions.IOStreams{
		In:     os.Stdin,
		Out:    os.Stdout,
		ErrOut: os.Stderr,
	}

	return &PortForwarder{
		RestConfig: restconf,
		Clientset:  kubernetes.NewForConfigOrDie(restconf),
		Streams:    streams,
	}
}

// NewPortForwarderFromRorKubernetesClient creates a new PortForwarder instance from a RorKubernetesClient
func NewPortForwarderFromRorKubernetesClient(k8sclientset *rorkubernetesclient.K8sClientsets) (*PortForwarder, error) {

	streams := genericclioptions.IOStreams{
		In:     os.Stdin,
		Out:    os.Stdout,
		ErrOut: os.Stderr,
	}

	clientset, err := k8sclientset.GetKubernetesClientset()
	if err != nil {
		return nil, err
	}

	return &PortForwarder{
		RestConfig: k8sclientset.GetConfig(),
		Clientset:  clientset,
		Streams:    streams,
	}, nil
}

// Forward starts the port forwarding process
func (p *PortForwarder) Forward(readyChan chan struct{}, stopChan <-chan struct{}) error {

	podname, podnamespace, err := p.getPodNameNamespace()
	if err != nil {
		return err
	}

	path := fmt.Sprintf("/api/v1/namespaces/%s/pods/%s/portforward",
		podnamespace, podname)
	hostIP := strings.TrimLeft(p.RestConfig.Host, "htps:/")

	transport, upgrader, err := spdy.RoundTripperFor(p.RestConfig)
	if err != nil {
		return err
	}

	containerport, err := p.GetContainerPort()
	if err != nil {
		return err
	}

	localport, err := p.GetLocalPort()
	if err != nil {
		return err
	}

	dialer := spdy.NewDialer(upgrader, &http.Client{Transport: transport}, http.MethodPost, &url.URL{Scheme: "https", Path: path, Host: hostIP})
	fw, err := portforward.New(dialer, []string{fmt.Sprintf("%d:%d", localport, containerport)}, stopChan, readyChan, p.Streams.Out, p.Streams.ErrOut)
	if err != nil {
		return err
	}
	return fw.ForwardPorts()
}

// GetPodNameNamespace returns the pod name and namespace
func (p *PortForwarder) getPodNameNamespace() (string, string, error) {
	var err error
	if p.Pod == nil {
		return "", "", errors.New("No pod added to portforwarder")
	}
	return p.Pod.Name, p.Pod.Namespace, err
}

// SetLocalPort sets the local port for the pod
func (p *PortForwarder) SetLocalPort(local int32) {
	p.LocalPort = local

}

// GetLocalPort returns the local port for the pod
func (p *PortForwarder) GetLocalPort() (int32, error) {
	if p.LocalPort != 0 {
		return p.LocalPort, nil
	}

	err := p.getFreePort()

	if err != nil {
		return 0, errors.Wrap(err, "Getting free port")
	}
	return p.LocalPort, nil
}

// SetContainerPort sets the container port for the pod
func (p *PortForwarder) SetContainerPort(port int32) {
	p.ContainerPort = port
}

// GetContainerPort returns the container port for the pod
func (p *PortForwarder) GetContainerPort() (int32, error) {
	if p.ContainerPort != 0 {
		return p.ContainerPort, nil
	}
	if p.Pod == nil {
		return 0, errors.New("No pod added to portforwarder")
	}
	containerPort := p.Pod.Spec.Containers[0].Ports[0].ContainerPort
	fmt.Printf("No pod port specified, using default port %d\n", containerPort)
	p.ContainerPort = containerPort
	return containerPort, nil
}

// AddPodByName adds a pod by its name/namespace
// It returns an error if the pod is not found or is not running.
func (p *PortForwarder) AddPodByName(name string, namespace string) error {
	ctx := context.Background()
	pod, err := p.Clientset.CoreV1().Pods(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return errors.Wrap(err, "Getting pod in kubernetes")
	}

	if pod.Status.Phase != v1.PodRunning {
		return errors.New(fmt.Sprintf("Pod %s is not running", name))
	}

	p.Pod = pod
	return nil
}

// AddPodByServiceName adds a pod by its service name/namespace
// It returns an error if the service is not found or is not running.
func (p *PortForwarder) AddPodByServiceName(service string, namespace string) error {
	ctx := context.Background()
	svc, err := p.Clientset.CoreV1().Services(namespace).Get(ctx, service, metav1.GetOptions{})
	if err != nil {
		fmt.Println("Error getting service")
	}

	err = p.AddPodByLabels(metav1.LabelSelector{
		MatchLabels: svc.Spec.Selector,
	}, namespace)
	if err != nil {
		return errors.Wrap(err, "Adding pod by service")
	}
	return nil

}

// Add  a pod by label, returns an error if the label returns
// more or less than one pod.
// It searches for the labels specified by labels.
func (p *PortForwarder) AddPodByLabels(labels metav1.LabelSelector, namespace string) error {
	if len(labels.MatchLabels) == 0 && len(labels.MatchExpressions) == 0 {
		return errors.New("No pod labels specified")
	}

	ctx := context.Background()
	pods, err := p.Clientset.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{
		LabelSelector: metav1.FormatLabelSelector(&labels),
		FieldSelector: fields.OneTermEqualSelector("status.phase", string(v1.PodRunning)).String(),
	})

	if err != nil {
		return errors.Wrap(err, "Listing pods in kubernetes")
	}

	formatted := metav1.FormatLabelSelector(&labels)

	if len(pods.Items) == 0 {
		return errors.New(fmt.Sprintf("Could not find running pod for selector: labels \"%s\"", formatted))
	}

	p.Pod = &pods.Items[0]
	return nil
}

// getFreePort returns a free port
func (p *PortForwarder) getFreePort() error {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return err
	}

	port := listener.Addr().(*net.TCPAddr).Port
	err = listener.Close()
	if err != nil {
		return err
	}

	p.LocalPort = int32(port)
	return nil
}
