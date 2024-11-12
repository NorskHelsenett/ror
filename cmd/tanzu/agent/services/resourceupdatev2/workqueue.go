package resourceupdatev2

import (
	"time"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
)

type ResourceCacheWorkqueueObject struct {
	SubmittedTime  time.Time
	RetryCount     int
	ResourceUpdate *apiresourcecontracts.ResourceUpdateModel
}

type ResourceCacheWorkqueue []ResourceCacheWorkqueueObject

func (wq ResourceCacheWorkqueue) NeedToRun() bool {
	return len(wq) > 0
}
func (wq ResourceCacheWorkqueue) ItemCount() int {
	return len(wq)
}

func (m ResourceCacheWorkqueue) GetByUid(uid string) (ResourceCacheWorkqueueObject, int) {
	for i, resourceUpdate := range m {
		if resourceUpdate.ResourceUpdate.Uid == uid {
			return resourceUpdate, i
		}
	}
	var emptyResponse ResourceCacheWorkqueueObject
	return emptyResponse, 0
}
func (m *ResourceCacheWorkqueue) Add(resourceUpdate *apiresourcecontracts.ResourceUpdateModel) {

	resourceUpdateObject := ResourceCacheWorkqueueObject{
		SubmittedTime:  time.Now(),
		RetryCount:     0,
		ResourceUpdate: resourceUpdate,
	}

	cacheResource, currentId := m.GetByUid(resourceUpdate.Uid)
	if currentId > 0 {
		resourceUpdateObject.RetryCount = cacheResource.RetryCount + 1
		(*m)[currentId] = resourceUpdateObject
	} else {
		*m = append(*m, resourceUpdateObject)
	}
}

func (m *ResourceCacheWorkqueue) DeleteByUid(uid string) {
	if uid == "" {
		return
	}
	var newCache []ResourceCacheWorkqueueObject

	_, id := m.GetByUid(uid)
	for i, resourceUpdateCacheObject := range *m {
		if i != id {
			newCache = append(newCache, resourceUpdateCacheObject)
		}
	}
	*m = newCache
}
