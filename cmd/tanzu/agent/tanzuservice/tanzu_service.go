package tanzuservice

import (
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/tanzu/agent/auth"
	"github.com/NorskHelsenett/ror/cmd/tanzu/agent/controllers"
	"github.com/NorskHelsenett/ror/cmd/tanzu/agent/rorclient"
	"github.com/NorskHelsenett/ror/cmd/tanzu/agent/tanzuservice/schemas"
	"os"
	"time"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
	"github.com/NorskHelsenett/ror/pkg/rorresources/rordefs"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/spf13/viper"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
)

var (
	stopCh chan struct{}
)

func init() {
}

func Run(sigs chan os.Signal, stop chan struct{}) {
	loginEveryMinute := viper.GetInt(configconsts.TANZU_AGENT_LOGIN_EVERY_MINUTE)
	go refreshLoginEveryMinute(time.Duration(loginEveryMinute) * time.Minute)
}

func refreshLoginEveryMinute(duration time.Duration) {
	count := 0
	for {
		tanzuAccess := viper.GetBool(configconsts.TANZU_AGENT_TANZU_ACCESS)
		if tanzuAccess {
			err := auth.Login()
			if err != nil {
				panic(err)
			}
		}

		k8sconfig, err := auth.GetK8sConfig(tanzuAccess)
		if err != nil {
			panic(err)
		}

		if stopCh == nil {
			stopCh = make(chan struct{})
		}

		rlog.Debug("Setting up clients and watchers")
		err = setupClientsAndWatchers(k8sconfig, stopCh)
		if err != nil {
			rlog.Error("Failed to setup clients and watchers", err)
			panic(err)
		}
		count++

		rlog.Info("refreshLoginEveryMinute runs ", rlog.String("count", fmt.Sprintf("%d", count)))

		time.Sleep(duration)
	}
}

func setupClientsAndWatchers(k8sConfig *rest.Config, stop chan struct{}) error {
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(k8sConfig)
	if err != nil {
		rlog.Error("failed to get discovery client", err)
		return err
	}

	dynamicClient, err := dynamic.NewForConfig(k8sConfig)
	if err != nil {
		rlog.Error("failed to get dynamic client", err)
		return err
	}

	unfilteredNamespaces, err := rorclient.GetWorkspaces()
	if err != nil {
		rlog.Error("failed to get namespaces from ror", err)
		return err
	}

	datacenterUrl := viper.GetString(configconsts.TANZU_AGENT_DATACENTER_URL)
	namespaces := make([]apicontracts.Workspace, 0)
	for _, n := range unfilteredNamespaces {
		if n.Datacenter.APIEndpoint == datacenterUrl {
			namespaces = append(namespaces, n)
		}
	}

	if len(namespaces) == 0 {
		rlog.Error("Zero namespaces to watch ... ", nil)
	}

	return setupWatchers(discoveryClient, dynamicClient, stop, namespaces)
}

func setupWatchers(discoveryClient *discovery.DiscoveryClient, dynamicClient dynamic.Interface, stop chan struct{}, namespaces []apicontracts.Workspace) error {
	allResoursesDefinitions := schemas.InitSchema()

	namespacedResourceDefinitions := make([]schema.GroupVersionResource, 0)
	resourceDefinitions := make([]schema.GroupVersionResource, 0)

	for _, resourceDefinition := range allResoursesDefinitions {
		resource, err := findResourceDefinition(resourceDefinition)
		if err != nil {
			continue
		}

		if resource.Namespaced {
			namespacedResourceDefinitions = append(namespacedResourceDefinitions, resourceDefinition)
		} else {
			resourceDefinitions = append(resourceDefinitions, resourceDefinition)
		}
	}

	shouldReturn, returnValue := addSchemasToNamespaceListener(resourceDefinitions, discoveryClient, dynamicClient, apicontracts.Workspace{}, stop)
	if shouldReturn {
		return returnValue
	}

	for _, namespace := range namespaces {
		shouldReturn, returnValue = addSchemasToNamespaceListener(namespacedResourceDefinitions, discoveryClient, dynamicClient, namespace, stop)
		if shouldReturn {
			return returnValue
		}
	}

	return nil
}

func addSchemasToNamespaceListener(resourceDefinitions []schema.GroupVersionResource,
	discoveryClient *discovery.DiscoveryClient,
	dynamicClient dynamic.Interface,
	namespace apicontracts.Workspace,
	stop chan struct{}) (bool, error) {
	for _, definition := range resourceDefinitions {
		rlog.Debug("Adding listener ", rlog.String("namespace", namespace.Name), rlog.Any("definition", definition))
		check, err := discovery.IsResourceEnabled(discoveryClient, definition)
		if err != nil {
			rlog.Error("Could not query resources from cluster", err)
			return true, err
		}
		if check {
			controller := controllers.NewDynamicController(dynamicClient, definition, namespace)

			go func() {
				defer close(stop)
				controller.Run(stop)
				stop <- struct{}{}
			}()
		} else {
			errmsg := fmt.Sprintf("Could not register resource %s", definition.Resource)
			rlog.Info(errmsg)
			return true, err
		}
	}
	return false, nil
}

func findResourceDefinition(gvp schema.GroupVersionResource) (rordefs.ApiResource, error) {

	for _, resource := range rordefs.Resourcedefs {
		if resource.GroupVersionKind().Group == gvp.Group && resource.GroupVersionKind().Version == gvp.Version && resource.Plural == gvp.Resource {
			return resource, nil
		}
	}

	return rordefs.ApiResource{}, fmt.Errorf("could not find resource definition for %s", gvp.String())
}
