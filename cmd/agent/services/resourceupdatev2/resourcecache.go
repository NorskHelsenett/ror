package resourceupdatev2

import (
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/agent/services/authservice"
	"time"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/NorskHelsenett/ror/pkg/helpers/stringhelper"

	"github.com/go-co-op/gocron"
)

var ResourceCache resourcecache

type resourcecache struct {
	HashList       hashList
	Workqueue      ResourceCacheWorkqueue
	cleanupRunning bool
	scheduler      *gocron.Scheduler
}

func (rc *resourcecache) Init() error {
	var err error
	rc.HashList, err = getResourceHashList()
	if err != nil {
		return err
	}
	rc.scheduler = gocron.NewScheduler(time.Local)
	rc.scheduler.StartAsync()
	rc.addWorkqueScheduler(10)
	rc.startCleanup()
	return nil
}

func (rc resourcecache) CleanupRunning() bool {
	return rc.cleanupRunning
}
func (rc *resourcecache) MarkActive(uid string) {
	rc.HashList.markActive(uid)
}

func (rc resourcecache) addWorkqueScheduler(seconds int) {
	_, _ = rc.scheduler.Every(seconds).Second().Tag("workquerunner").Do(rc.runWorkqueScheduler)
}
func (rc resourcecache) runWorkqueScheduler() {
	if rc.Workqueue.NeedToRun() {
		rlog.Warn("resourceQue has non zero lenght", rlog.Int("resource que length", rc.Workqueue.ItemCount()))
		rc.RunWorkQue()
	}
}

func (rc *resourcecache) startCleanup() {
	rc.cleanupRunning = true
	_, _ = rc.scheduler.Every(1).Day().At(time.Now().Add(time.Minute * 1)).Tag("resourcescleanup").Do(rc.finnishCleanup)
}

func (rc *resourcecache) finnishCleanup() {
	if !rc.cleanupRunning {
		return
	}
	rc.cleanupRunning = false
	_ = rc.scheduler.RemoveByTag("resourcescleanup")
	inactive := rc.HashList.getInactiveUid()
	if len(inactive) == 0 {
		return
	}
	for _, uid := range inactive {
		rlog.Info(fmt.Sprintf("Removing resource %s", uid))
		resource := apiresourcecontracts.ResourceUpdateModel{
			Owner:  authservice.CreateOwnerref(),
			Uid:    uid,
			Action: apiresourcecontracts.K8sActionDelete,
		}
		_ = sendResourceUpdateToRor(&resource)
	}
	rlog.Info(fmt.Sprintf("resource cleanup done, %d resources removed", len(inactive)))
}

func (rc resourcecache) PrettyPrintHashes() {
	stringhelper.PrettyprintStruct(rc.HashList)
}

// RunWorkQue Will run from the scheduler if the resource-que is non zero length.
// Resources in the que wil be requed using the sendResourceUpdateToRor function.
func (rc *resourcecache) RunWorkQue() {
	for _, resourceReturn := range rc.Workqueue {
		err := sendResourceUpdateToRor(resourceReturn.ResourceUpdate)
		if err != nil {
			rlog.Error("error re-sending resource update to ror, added to retryque", err)
			rc.Workqueue.Add(resourceReturn.ResourceUpdate)
			return
		}
		rc.Workqueue.DeleteByUid(resourceReturn.ResourceUpdate.Uid)
		rc.HashList.updateHash(resourceReturn.ResourceUpdate.Uid, resourceReturn.ResourceUpdate.Hash)
	}
}
