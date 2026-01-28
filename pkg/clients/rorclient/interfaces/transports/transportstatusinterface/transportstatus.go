package transportstatusinterface

type RorTransportStatus interface {
	IsEstablished() bool
	GetApiVersion() string
	GetLibVersion() string
}
