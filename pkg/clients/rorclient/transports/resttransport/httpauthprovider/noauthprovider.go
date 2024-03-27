package httpauthprovider

import "net/http"

type NoAuthprovider struct {
}

func NewNoAuthprovider() *NoAuthprovider {
	return &NoAuthprovider{}
}

func (a NoAuthprovider) AddAuthHeaders(req *http.Request) {

}
