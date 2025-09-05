package rorserviceinterface

import "os"

type RorService interface {
	GetCancelChan() chan os.Signal
	GetStopChan() chan struct{}
	Wait()

	Cancel()
	Stop()
}
