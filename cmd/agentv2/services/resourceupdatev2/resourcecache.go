package resourceupdatev2

import (
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/agentv2/clients"
	"github.com/NorskHelsenett/ror/cmd/agentv2/services/authservice"
	"time"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	"github.com/NorskHelsenett/ror/pkg/apicontracts/v2/apicontractsv2resources"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/go-co-op/gocron"
)

var ResourceCache resourcecache

type resourcecache struct {
	HashList       *apicontractsv2resources.HashList
	WorkQueue      ResourceCacheWorkQueue
	cleanupRunning bool
	scheduler      *gocron.Scheduler
}

func (rc *resourcecache) Init() error {
	var err error
	rc.HashList, err = InitHashList()
	if err != nil {
		return err
	}
	rc.WorkQueue = NewResourceCacheWorkQueue()
	rc.scheduler = gocron.NewScheduler(time.Local)
	rc.addWorkQeueScheduler(10)
	rc.scheduler.StartAsync()
	rc.startCleanup()
	return nil
}

func (rc *resourcecache) CleanupRunning() bool {
	return rc.cleanupRunning
}
func (rc *resourcecache) MarkActive(uid string) {
	rc.HashList.MarkActive(uid)
}

func (rc *resourcecache) addWorkQeueScheduler(seconds int) {
	_, err := rc.scheduler.Every(seconds).Second().Tag("workQeuerunner").Do(rc.runWorkQeueScheduler)
	if err != nil {
		rlog.Error("error starting workQeueScheduler", err)
	}

}
func (rc *resourcecache) runWorkQeueScheduler() {
	if rc.WorkQueue.NeedToRun() {
		rlog.Warn("resourceQueue has non zero lenght", rlog.Int("resource Queue length", rc.WorkQueue.ItemCount()))
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
	inactive := rc.HashList.GetInactiveUid()
	if len(inactive) == 0 {
		return
	}
	for _, uid := range inactive {
		rlog.Info(fmt.Sprintf("Removing resource %s", uid))

		// rorres := rorkubernetes.NewResourceFromDynamicClient(input)
		// err := rorres.SetRorMeta(rortypes.ResourceRorMeta{
		// 	Version:  "v2",
		// 	Ownerref: clients.RorConfig.CreateOwnerref(),
		// 	Action:   action,
		// })

		resource := apiresourcecontracts.ResourceUpdateModel{
			Owner:  authservice.CreateOwnerref(),
			Uid:    uid,
			Action: apiresourcecontracts.K8sActionDelete,
		}
		_ = sendResourceUpdateToRor(&resource)
	}
	rlog.Info(fmt.Sprintf("resource cleanup done, %d resources removed", len(inactive)))
}

// RunWorkQueue Will run from the scheduler if the resource-Queue is non zero length.
// Resources in the Queue wil be reQeued using the sendResourceUpdateToRor function.
func (rc *resourcecache) RunWorkQeue() {
	if !rc.WorkQueue.NeedToRun() {
		return
	}
	cacheworkqueue := rc.WorkQueue.ConsumeWorkQeue()
	rorclient := clients.RorConfig.GetRorClient()
	status, err := rorclient.ResourceV2().Update(cacheworkqueue.ResourceSet)
	if err != nil {
		rlog.Error("error sending resources update to ror, added to retryQeue", err)
		rc.WorkQueue.reQueue(cacheworkqueue)
		return
	}

	failed := status.GetFailedResources()
	if len(failed) > 0 {
		for failuuid, result := range failed {
			rlog.Error("error sending resource update to ror, added to retryQeue", fmt.Errorf("uid: %s, failed with status: %d message: %s", failuuid, result.Status, result.Message))
			rc.WorkQueue.reQueueResource(cacheworkqueue.GetByUid(failuuid), cacheworkqueue.GetRetrycount(failuuid))
		}
	}
}
