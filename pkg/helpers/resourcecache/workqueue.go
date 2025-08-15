package resourcecache

import (
	"fmt"
	"sync"

	"github.com/NorskHelsenett/ror/pkg/rorresources"

	"github.com/NorskHelsenett/ror/pkg/rlog"
)

type ResourceCacheWorkQueue struct {
	mu sync.RWMutex
	*rorresources.ResourceSet
	RetryCount map[string]int
}

func NewResourceCacheWorkQueue() ResourceCacheWorkQueue {
	return ResourceCacheWorkQueue{ResourceSet: rorresources.NewResourceSet(), RetryCount: make(map[string]int)}
}

func (wq *ResourceCacheWorkQueue) GetRetrycount(uid string) int {
	wq.mu.RLock()
	defer wq.mu.RUnlock()
	return wq.RetryCount[uid]
}

func (wq *ResourceCacheWorkQueue) SetRetrycount(uid string, count int) {
	wq.mu.Lock()
	defer wq.mu.Unlock()
	wq.RetryCount[uid] = count
}

func (wq *ResourceCacheWorkQueue) AddResource(add *rorresources.Resource) {
	wq.mu.Lock()
	defer wq.mu.Unlock()
	wq.Add(add)
	ResourceCache.HashList.UpdateHash(add.GetUID(), add.GetRorHash())
}

func (wq *ResourceCacheWorkQueue) AddResourceSet(add *rorresources.ResourceSet) {
	wq.mu.Lock()
	defer wq.mu.Unlock()
	for wq.Next() {
		add.Add(wq.Get())
	}
}

// TODO use reQueueResource to add resources to the work queue
func (wq *ResourceCacheWorkQueue) reQueue(wqadd *ResourceCacheWorkQueue) {
	faileduuids := ""
	var failcount int
	for wqadd.Next() {
		resource := wqadd.Get()
		retrycount := wqadd.RetryCount[resource.GetUID()] + 1
		if retrycount > 10 {
			if faileduuids != "" {
				faileduuids = fmt.Sprintf("%s, %s", faileduuids, resource.GetUID())
			} else {
				faileduuids = resource.GetUID()
			}
			failcount++
			wq.DeleteByUid(resource.GetUID())
		} else {
			wq.AddResource(resource)
			wq.SetRetrycount(resource.GetUID(), retrycount)
		}
	}
	if faileduuids != "" {
		rlog.Error("retry limit reached", fmt.Errorf("%d resources has been retried 10 times, giving up", failcount), rlog.String("uuids", faileduuids))
	}

}

func (wq *ResourceCacheWorkQueue) reQueueResource(resource *rorresources.Resource, retrycount int) {
	retrycount++
	if retrycount > 10 {
		rlog.Error("retry limit reached", fmt.Errorf("resource has been retried 10 times, giving up"), rlog.String("uuids", resource.GetUID()))
		wq.DeleteByUid(resource.GetUID())
	} else {
		wq.AddResource(resource)
		wq.SetRetrycount(resource.GetUID(), retrycount)
	}
}

func (wq *ResourceCacheWorkQueue) NeedToRun() bool {
	return wq.ItemCount() > 0
}
func (wq *ResourceCacheWorkQueue) ItemCount() int {
	wq.mu.RLock()
	defer wq.mu.RUnlock()
	return wq.Len()
}

func (wq *ResourceCacheWorkQueue) GetQuedResourceByUid(uid string) *rorresources.Resource {
	wq.mu.RLock()
	defer wq.mu.RUnlock()
	if res := wq.GetByUid(uid); res != nil {
		return res
	}
	return nil
}

func (wq *ResourceCacheWorkQueue) DeleteByUid(uid string) {
	wq.mu.Lock()
	defer wq.mu.Unlock()
	if uid == "" {
		return
	}
	wq.ResourceSet.DeleteByUid(uid)
	wq.RetryCount[uid] = 0
}

func (wq *ResourceCacheWorkQueue) ConsumeWorkQeue() *ResourceCacheWorkQueue {
	wq.mu.Lock()
	returnQueue := wq.DeepCopy()
	wq.mu.Unlock()
	for _, res := range returnQueue.Resources {
		wq.DeleteByUid(res.GetUID())
	}
	return returnQueue
}

func (wq *ResourceCacheWorkQueue) DeepCopy() *ResourceCacheWorkQueue {
	returnQueue := &ResourceCacheWorkQueue{
		ResourceSet: rorresources.NewResourceSet(),
		RetryCount:  make(map[string]int),
	}
	for wq.Next() {
		r := wq.Get()
		returnQueue.Add(r)
	}
	returnQueue.RetryCount = wq.RetryCount
	return returnQueue
}
