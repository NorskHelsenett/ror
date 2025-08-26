package resourcecache

import (
	"context"
	"fmt"
	"time"

	"github.com/NorskHelsenett/ror/pkg/clients/rorclient"
	"github.com/NorskHelsenett/ror/pkg/helpers/resourcecache/resourcecachehashlist"
	"github.com/NorskHelsenett/ror/pkg/helpers/resourcecache/workqueue"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/rorresourceowner"
	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/NorskHelsenett/ror/pkg/rorresources"
	"github.com/NorskHelsenett/ror/pkg/rorresources/rortypes"

	"github.com/go-co-op/gocron"
	"k8s.io/apimachinery/pkg/types"
)

// ResourceCacheInterface defines the contract for resource cache operations
type ResourceCacheInterface interface {
	// CleanupRunning returns whether cleanup is currently running
	CleanupRunning() bool

	// MarkActive marks a resource as active by its UID
	MarkActive(uid string)

	// RunWorkQeue processes the work queue
	RunWorkQeue()

	// AddResource adds a single resource to the work queue
	AddResource(resource *rorresources.Resource)

	// AddResourceSet adds multiple resources to the work queue
	AddResourceSet(resources *rorresources.ResourceSet)

	DeleteResourceByUID(uid string)

	CheckUpdateNeeded(uid string, hash string) bool

	GetOwnerref() rorresourceowner.RorResourceOwnerReference
}

var ResourceCache *resourcecache

// Compile-time check to ensure resourcecache implements ResourceCacheInterface
var _ ResourceCacheInterface = (*resourcecache)(nil)

type resourcecache struct {
	hashList       *resourcecachehashlist.HashList
	rorClient      rorclient.RorClientInterface
	workQueue      workqueue.ResourceCacheWorkQueue
	cleanupRunning bool
	scheduler      *gocron.Scheduler
}

type ResourceCacheConfig struct {
	WorkQueueInterval int
	RorClient         rorclient.RorClientInterface
}

func NewResourceCache(rcConfig ResourceCacheConfig) (ResourceCacheInterface, error) {
	return newResourceCache(rcConfig)
}

func newResourceCache(rcConfig ResourceCacheConfig) (*resourcecache, error) {

	var rc resourcecache

	if rcConfig.RorClient == nil {
		return nil, fmt.Errorf("ror client is required for resource cache initialization")
	}

	rc.rorClient = rcConfig.RorClient

	if rc.rorClient.Ping() == nil {
		hashes, err := rc.rorClient.ResourceV2().GetOwnHashes(context.TODO(), rc.rorClient.GetOwnerref())
		if err != nil {
			rc.hashList = NewEmptyHashList()
			return nil, fmt.Errorf("failed to get own hashes: %w", err)
		}
		rc.hashList = hashes
	} else {
		rc.hashList = NewEmptyHashList()
	}

	rc.workQueue = workqueue.NewResourceCacheWorkQueue()
	rc.scheduler = gocron.NewScheduler(time.Local)
	rc.addWorkQeueScheduler(rcConfig.WorkQueueInterval)
	rc.scheduler.StartAsync()
	rc.startCleanup()
	return &rc, nil
}

// Deprecated: Use NewResourceCache instead
func (rc *resourcecache) Init() error {
	var rcConfig ResourceCacheConfig
	rcConfig.WorkQueueInterval = 10
	newrc, err := newResourceCache(rcConfig)
	if err != nil {
		return fmt.Errorf("error initializing resource cache: %w", err)
	}

	ResourceCache = newrc
	rlog.Info("resource cache initialized")
	return nil
}

func (rc *resourcecache) CleanupRunning() bool {
	return rc.cleanupRunning
}
func (rc *resourcecache) MarkActive(uid string) {
	rc.hashList.MarkActive(uid)
}

func (rc *resourcecache) addWorkQeueScheduler(seconds int) {
	_, err := rc.scheduler.Every(seconds).Second().Tag("workQeuerunner").Do(rc.runWorkQeueScheduler)
	if err != nil {
		rlog.Error("error starting workQeueScheduler", err)
	}

}
func (rc *resourcecache) runWorkQeueScheduler() {
	if rc.workQueue.NeedToRun() {
		rlog.Info("resourceQueue has non zero length", rlog.Int("resource Queue length", rc.workQueue.ItemCount()))
		rc.RunWorkQeue()
	}
}

func (rc *resourcecache) startCleanup() {
	rc.cleanupRunning = true
	_, err := rc.scheduler.Every(1).Day().At(time.Now().Add(time.Minute * 1)).Tag("resourcescleanup").Do(rc.finnishCleanup)
	if err != nil {
		rlog.Error("error starting cleanup", err)
	}
}

func (rc *resourcecache) finnishCleanup() {
	if !rc.cleanupRunning {
		return
	}
	rc.cleanupRunning = false
	_ = rc.scheduler.RemoveByTag("resourcescleanup")
	inactive := rc.hashList.GetInactiveUid()
	if len(inactive) != 0 {
		for _, uid := range inactive {
			if uid == "" {
				rlog.Warn("resource uid is empty")
				continue
			}
			rc.DeleteResourceByUID(uid)
		}
	}

	rlog.Info(fmt.Sprintf("resource cleanup done, %d resources removed", len(inactive)))
}

// RunWorkQueue Will run from the scheduler if the resource-Queue is non zero length.
// Resources in the Queue wil be reQeued using the sendResourceUpdateToRor function.
func (rc *resourcecache) RunWorkQeue() {
	if !rc.workQueue.NeedToRun() {
		return
	}
	cacheworkqueue := rc.workQueue.ConsumeWorkQeue()

	status, err := rc.rorClient.ResourceV2().Update(context.Background(), cacheworkqueue.ResourceSet)
	if err != nil {
		rlog.Error("error sending resources update to ror, added to retryQeue", err)
		rc.workQueue.ReQueue(cacheworkqueue)
		return
	}

	failed := status.GetFailedResources()
	if len(failed) > 0 {
		for failuuid, result := range failed {
			rlog.Error("error sending resource update to ror, added to retryQeue", fmt.Errorf("uid: %s, failed with status: %d message: %s", failuuid, result.Status, result.Message))
			rc.workQueue.ReQueueResource(cacheworkqueue.GetByUid(failuuid), cacheworkqueue.GetRetrycount(failuuid))
		}
	}
}

// AddResource adds a single resource to the work queue
func (rc *resourcecache) AddResource(resource *rorresources.Resource) {
	rc.workQueue.AddResource(resource)
	rc.hashList.UpdateHash(resource.GetUID(), resource.GetRorHash())
}

// AddResourceSet adds multiple resources to the work queue
func (rc *resourcecache) AddResourceSet(resources *rorresources.ResourceSet) {
	for resources.Next() {
		rc.AddResource(resources.Get())
	}
}

func (rc *resourcecache) CheckUpdateNeeded(uid string, hash string) bool {
	return rc.hashList.CheckUpdateNeeded(uid, hash)
}

func (rc *resourcecache) GetOwnerref() rorresourceowner.RorResourceOwnerReference {
	return rc.rorClient.GetOwnerref()
}

func (rc *resourcecache) DeleteResourceByUID(uid string) {
	if uid == "" {
		return
	}
	resource := rorresources.NewRorResource("Unknown", "unknown.ror.internal/v1")
	resource.Metadata.SetUID(types.UID(uid))
	resource.SetRorMeta(rortypes.ResourceRorMeta{
		Version:  "v2",
		Ownerref: rc.rorClient.GetOwnerref(),
		Action:   rortypes.K8sActionDelete,
	})
	rc.workQueue.AddResource(resource)
	rc.hashList.DeleteByUid(uid)
}
