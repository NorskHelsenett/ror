package scheduler

import (
	"time"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/go-co-op/gocron"
)

func SetUpScheduler() {
	scheduler := gocron.NewScheduler(time.UTC)
	_, err := scheduler.Every(1).Minute().Tag("metrics").Do(MetricsReporting)
	if err != nil {
		rlog.Fatal("Could not setup scheduler for metrics", err)
		return
	}
	//scheduler.StartAsync()
}
