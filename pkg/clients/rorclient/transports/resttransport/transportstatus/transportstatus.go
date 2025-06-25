package transportstatus

type RorTransportStatus interface {
	IsEstablished() bool
	GetApiVersion() string
	GetLibVersion() string
}
