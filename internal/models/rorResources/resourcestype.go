package rorResources

import (
	"crypto/md5" // #nosec G501 - MD5 is used for hash calculation only
	"encoding/json"
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	jsonpatch "github.com/evanphx/json-patch/v5"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type rorResource struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Uid        string `json:"uid"`
	Hash       string `json:"hash"`
	Resource   any    `json:"resource"`
}

type rorResourceJson []byte

func NewFromUnstructured(input *unstructured.Unstructured) (rorResource, error) {
	returnResource := rorResource{
		ApiVersion: input.GetAPIVersion(),
		Kind:       input.GetKind(),
		Uid:        string(input.GetUID()),
	}
	bytes, err := input.MarshalJSON()
	jsonData := rorResourceJson(bytes)
	if err != nil {
		rlog.Error("Could not unmarshal resource", err)
		return returnResource, err
	}

	err = jsonData.removeUnnecessaryData()
	if err != nil {
		return returnResource, err
	}
	returnResource.Hash, err = jsonData.calculateHash()
	if err != nil {
		return returnResource, err
	}
	err = jsonData.getResource(&returnResource)
	if err != nil {
		return returnResource, err
	}
	return returnResource, nil
}

func (r rorResource) NewResourceUpdateModel(owner apiresourcecontracts.ResourceOwnerReference, action apiresourcecontracts.ResourceAction) *apiresourcecontracts.ResourceUpdateModel {
	return &apiresourcecontracts.ResourceUpdateModel{
		Owner:      owner,
		ApiVersion: r.ApiVersion,
		Kind:       r.Kind,
		Uid:        r.Uid,
		Action:     action,
		Hash:       r.Hash,
		Resource:   r.Resource,
	}
}

func prepareResourcePayload[D any](input []byte) (D, error) {
	var outStruct D
	err := json.Unmarshal(input, &outStruct)
	if err != nil {
		rlog.Error("error unmarshaling json", err)
		return outStruct, err
	}
	return outStruct, nil
}

func (rj rorResourceJson) calculateHash() (string, error) {
	bytes := []byte(rj)
	patch := []byte(`{"metadata":{"resourceVersion":null,"creationTimestamp":null,"generation":null}}`)
	input, err := jsonpatch.MergePatch(bytes, patch)
	if err != nil {
		rlog.Error("error patching json", err)
		return "", err
	}
	resourceHash := fmt.Sprintf("%x", md5.Sum(input)) // #nosec G401 - MD5 is used for hash calculation only
	return resourceHash, nil

}
func (rj *rorResourceJson) removeUnnecessaryData() error {
	bytes := []byte(*rj)
	patch := []byte(`{"metadata":{"annotations":{"kubectl.kubernetes.io/last-applied-configuration":null}}}`)
	bytes, err := jsonpatch.MergePatch(bytes, patch)
	if err != nil {
		rlog.Error("error patching json", err)
		return err
	}
	*rj = rorResourceJson(bytes)

	return nil
}
