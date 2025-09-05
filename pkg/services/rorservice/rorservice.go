package rorservice

import (
	"os"

	"github.com/NorskHelsenett/ror/pkg/services/rorservice/rorserviceconfig"
	"github.com/NorskHelsenett/ror/pkg/services/rorservice/rorserviceinterface"
)

var Service rorserviceinterface.RorService

func MustInit() {
	Service = rorserviceconfig.MustInit()
}

func GetCancelChan() chan os.Signal {
	return Service.GetCancelChan()
}

func GetStopChan() chan struct{} {
	return Service.GetStopChan()
}

func Wait() {
	Service.Wait()
}

func Cancel() {
	Service.Cancel()
}

func Stop() {
	Service.Stop()
}
