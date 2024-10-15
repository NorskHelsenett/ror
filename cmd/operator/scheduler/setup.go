package scheduler

import (
	"github.com/NorskHelsenett/ror/cmd/operator/scheduler/tasks"
	"time"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/go-co-op/gocron"
	"k8s.io/client-go/kubernetes"
)

func SetUpScheduler(k8sClient *kubernetes.Clientset) {
	scheduler := gocron.NewScheduler(time.UTC)

	scheduler.SingletonModeAll()

	_, _ = scheduler.Every(3).Minute().Tag("ConfigurationChecks").Do(tasks.CheckForConfigChanges, k8sClient)

	scheduler.StartAsync()
	scheduler.StartBlocking()

	scheduleTasks := scheduler.Jobs()
	for _, scheduleTask := range scheduleTasks {
		rlog.Error("error in scheduled task", scheduleTask.Error())
	}
}
