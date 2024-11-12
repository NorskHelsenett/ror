package configuration

import (
	"context"
	"encoding/base64"
	resourcesservice "github.com/NorskHelsenett/ror/cmd/api/services/resourcesService"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"

	"github.com/NorskHelsenett/ror/pkg/rlog"
)

// NewResourceLoader creates a ConfigLoaderInterface from a resource
//
// Parameters:
//
//	ownerref: the owner reference (apiresourcecontracts.ResourceOwnerReference)
//	resourceName: the name of the resource (string)
//
// The resource must have its data represented in spec.data but can be base64 encoded if spec.b64enc is true
func NewResourceLoader(ownerref apiresourcecontracts.ResourceOwnerReference, resourceName string) ConfigLoaderInterface {

	if resourceName == "" {
		rlog.Error("resource name is empty", nil)
		return nil
	}

	resources, err := resourcesservice.GetConfigurations(context.TODO(), ownerref)
	if err != nil {
		rlog.Error("failed to get resources", err)
		return nil
	}

	if len(resources.Configurations) == 0 {
		rlog.Error("no resources found", nil)
		return nil
	}

	var data string
	resource := resources.GetByName(resourceName)
	parser := ParserType(resource.Spec.Type)

	if resource.Spec.B64enc {
		data = DecodeB64(resource.Spec.Data)
	} else {
		data = resource.Spec.Data
	}

	var parserinterface ConfigParserInterface
	switch parser {
	case ParserTypeJson:
		parserinterface = NewJsonParser([]byte(data))
	case ParserTypeYaml:
		parserinterface = NewYamlParser([]byte(data))
	default:
		return nil
	}
	ret := Loader{
		Parser: parserinterface,
		Secret: false,
	}

	return ret
}

func DecodeB64(data string) string {
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		rlog.Error("failed to decode b64", err)
		return ""
	}
	return string(decoded)
}
