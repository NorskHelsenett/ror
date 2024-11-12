package services

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"

	"github.com/NorskHelsenett/ror/pkg/rlog"
)

var EgressIp string

func GetEgressIp() {
	internettCheck := "https://api.ipify.org/"
	nhnCheck := "ip.nhn.no"
	_, err := net.LookupIP(nhnCheck)
	var apiHost string
	if err != nil {
		apiHost = internettCheck
	} else {
		apiHost = fmt.Sprintf("http://%s", nhnCheck)
	}

	rlog.Info("Resolving ip", rlog.String("api host", apiHost))
	res, err := http.Get(apiHost) // #nosec G107 - we are not using user input
	if err != nil {
		// assuming retry but on internett
		apiHost = internettCheck
		res, err = http.Get(apiHost) // #nosec G107 - we are not using user input
		if err != nil {
			errorMsg := fmt.Sprintf("could not reach host %s", apiHost)
			rlog.Info(errorMsg)
			return
		}
	}

	body, err := io.ReadAll(res.Body)
	_ = res.Body.Close()
	if res.StatusCode > 299 {
		rlog.Info("response failed", rlog.Int("status code", res.StatusCode), rlog.ByteString("body", body))
		return
	}

	if err != nil {
		rlog.Error("could not parse body", err)
		return
	}

	EgressIp = strings.Replace(string(body), "\n", "", -1)
}
