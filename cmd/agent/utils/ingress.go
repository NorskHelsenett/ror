package utils

import (
	"context"
	"fmt"
	"strings"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	networkingV1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/utils/strings/slices"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

func GetIngressDetails(ingress *networkingV1.Ingress) (*apicontracts.Ingress, error) {
	var newIngress apicontracts.Ingress
	ingressNameSpace := ingress.Namespace
	ingressName := ingress.Name
	ingressClassName := ""

	var rules []apicontracts.IngressRule
	var health apicontracts.Health = 1

	if ingress.Spec.IngressClassName != nil {
		ingressClassName = *ingress.Spec.IngressClassName
	}

	k8sConfig := config.GetConfigOrDie()
	k8sClient, err := kubernetes.NewForConfig(k8sConfig)
	if err != nil {
		rlog.Error("error in config", err)
	}

	if ingress.Spec.Rules == nil {
		return nil, fmt.Errorf("invalid ingress - missing rules")
	}

	for ruleindex, irule := range ingress.Spec.Rules {
		rlog.Debug("rule for host", rlog.String("host", irule.Host))
		rules = append(rules, apicontracts.IngressRule{
			Hostname:    irule.Host,
			IPAddresses: nil,
			Paths:       nil,
		})

		if ingress.Status.LoadBalancer.Ingress == nil {
			rlog.Debug("Ingress has no IP-address", rlog.String("ingress", ingress.Name))
		} else {
			for _, is := range ingress.Status.LoadBalancer.Ingress {
				if is.Hostname == irule.Host {
					rules[ruleindex].IPAddresses = append(rules[ruleindex].IPAddresses, is.IP)
				}
			}
		}

		for _, irulepath := range irule.IngressRuleValue.HTTP.Paths {
			rlog.Debug("rule for path", rlog.String("path", irulepath.Path))
			rlog.Debug("", rlog.String("service", irulepath.Backend.Service.Name))
			service, err := GetIngressService(k8sClient, ingressNameSpace, irulepath.Backend.Service.Name)
			if err != nil {
				rules[ruleindex].Paths = append(rules[ruleindex].Paths, apicontracts.IngressPath{
					Path:    irulepath.Path,
					Service: apicontracts.Service{},
				})
				continue
			}
			rules[ruleindex].Paths = append(rules[ruleindex].Paths, apicontracts.IngressPath{
				Path:    irulepath.Path,
				Service: service,
			})
		}
	}

	newIngress = apicontracts.Ingress{
		UID:       string(ingress.UID),
		Health:    health,
		Name:      ingressName,
		Namespace: ingressNameSpace,
		Class:     ingressClassName,
		Rules:     rules,
	}

	richIngress, err := GetIngressHealth(newIngress)

	return richIngress, nil

}

func GetIngressHealth(thisIngress apicontracts.Ingress) (*apicontracts.Ingress, error) {

	ingressClasses := []string{"internett", "helsenett", "datacenter"}
	thisIngressClass := strings.Split(thisIngress.Class, "-")[len(strings.Split(thisIngress.Class, "-"))-1]

	if !slices.Contains(ingressClasses, thisIngressClass) {
		thisIngress.Health = 3
	}
	if len(thisIngress.Rules) < 1 {
		thisIngress.Health = 3
	} else {
		for _, rule := range thisIngress.Rules {
			if len(rule.IPAddresses) < 1 {
				thisIngress.Health = 3
			}
			if len(rule.Paths) < 1 {
				thisIngress.Health = 3
			} else {
				for _, path := range rule.Paths {
					if path.Service.Type != "NodePort" {
						thisIngress.Health = 3
					}
					if len(path.Service.Endpoints) < 0 {
						thisIngress.Health = 3
					}
				}
			}
		}
	}

	return &thisIngress, nil

}

func GetIngressService(k8sClient *kubernetes.Clientset, namespace string, serviceName string) (apicontracts.Service, error) {

	var service apicontracts.Service
	var endpoints []apicontracts.EndpointAddress
	var ports []apicontracts.ServicePort

	listOptions := metav1.ListOptions{}
	svcs, err := k8sClient.CoreV1().Services(namespace).List(context.TODO(), listOptions)
	if err != nil {
		rlog.Fatal("could not list svcs", err)
	}
	for _, svc := range svcs.Items {
		if svc.Name == serviceName {
			// if svc.Spec.Type != "NodePort" {
			// 	health = 3
			// }

			for _, port := range svc.Spec.Ports {
				ports = append(ports, apicontracts.ServicePort{
					Name:     port.Name,
					NodePort: fmt.Sprint(port.NodePort),
					Protocol: string(port.Protocol),
				})
			}

			service = apicontracts.Service{
				Name:      serviceName,
				Type:      string(svc.Spec.Type),
				Selector:  svc.Spec.Selector["app.kubernetes.io/name"],
				Ports:     ports,
				Endpoints: nil,
			}
			rlog.Debug("service added ", rlog.String("service", serviceName))
		}
	}

	if service.Name == "" {
		service = apicontracts.Service{
			Name:      serviceName,
			Type:      "",
			Selector:  "",
			Ports:     nil,
			Endpoints: nil,
		}
		rlog.Debug("Could not find Service", rlog.String("service name", serviceName))
	}

	eps, _ := k8sClient.CoreV1().Endpoints(namespace).List(context.TODO(), listOptions)
	if err != nil {
		rlog.Fatal("could not list eps", err)
	}
	for _, ep := range eps.Items {
		if ep.Name != serviceName {
			continue
		}

		if ep.Subsets == nil {
			continue
		}

		for _, epAddress := range ep.Subsets[0].Addresses {
			nodename := "None"
			if epAddress.NodeName != nil {
				nodename = *epAddress.NodeName
			}
			podname := "None"
			if epAddress.TargetRef != nil {
				podname = epAddress.TargetRef.Name
			}
			endpoints = append(endpoints, apicontracts.EndpointAddress{
				NodeName: nodename,
				PodName:  podname,
			})
			service.Endpoints = endpoints
		}

	}

	return service, nil

}
