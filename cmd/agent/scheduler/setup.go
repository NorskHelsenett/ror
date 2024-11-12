package scheduler

import (
	"time"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/go-co-op/gocron"
)

func SetUpScheduler() {
	scheduler := gocron.NewScheduler(time.UTC)
	_, err := scheduler.Every(1).Minute().Tag("heartbeat").Do(HeartbeatReporting)
	if err != nil {
		rlog.Fatal("Failed to setup heartbeat schedule", err)
	}

	_, err = scheduler.Every(1).Minute().Tag("metrics").Do(MetricsReporting)
	if err != nil {
		rlog.Fatal("Failed to setup metric schedule", err)
	}
	_ = scheduler.RunByTag("heartbeat")
	scheduler.StartAsync()
}
