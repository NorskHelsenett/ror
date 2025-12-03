package server

import (
	"fmt"
	"net"
	"net/http"
	"net/netip"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/NorskHelsenett/ror/pkg/config/rorconfig"
	"github.com/NorskHelsenett/ror/pkg/rlog"
	newhealth "github.com/dotse/go-health"
)

const (
	defaultPort = "9999"
	defaultIp   = "0.0.0.0"
)

//nolint:gochecknoglobals
var (
	httpServer *http.Server
	initMtx    sync.Mutex
)

type ServerParams struct {
	Server string
}

func ServerString(serverstring string) optionFunc {
	return optionFunc(func(cfg *config) {
		var err error
		serverstring = parseServerString(serverstring)
		cfg.ipPort, err = netip.ParseAddrPort(serverstring)
		if err != nil {
			rlog.Error("error parsing server string", err, rlog.String("serverstring", serverstring))
		}
	})
}

func parseServerString(serverstring string) string {
	if serverstring == "" {
		return fmt.Sprintf("%s:%s", defaultIp, defaultPort)
	}
	splits := strings.Split(serverstring, ":")
	if len(splits) == 2 {
		if splits[0] == "" {
			// only port
			splits[0] = defaultIp
		}
		if splits[1] == "" {
			// only ip
			splits[1] = defaultPort
		}
		// ip and port
		return strings.Join(splits, ":")
	}
	if len(splits) == 1 {
		_, err := strconv.ParseUint(splits[0], 10, 16)
		if err == nil {
			// only port
			return fmt.Sprintf("%s:%s", defaultIp, splits[0])
		}
		// only ip
		return fmt.Sprintf("%s:%s", splits[0], defaultPort)
	}
	// invalid
	rlog.Error("Invalid server string format", nil, rlog.String("serverstring", serverstring))

	return getDefaultServerString()
}

func getDefaultServerString() string {
	return fmt.Sprintf("%s:%s", defaultIp, defaultPort)
}

func getDefaultAddrPort() netip.AddrPort {
	addrPort, _ := netip.ParseAddrPort(getDefaultServerString())
	return addrPort
}

type config struct {
	ipPort netip.AddrPort
}

type optionFunc func(*config)

func Start(opts ...optionFunc) error {
	initMtx.Lock()
	defer initMtx.Unlock()
	cfg := &config{
		ipPort: getDefaultAddrPort(),
	}

	for _, o := range opts {
		o(cfg)
	}

	if httpServer == nil {
		listener, err := net.Listen("tcp", cfg.ipPort.String())
		if err != nil {
			return err
		}

		httpServer = &http.Server{
			Addr:              cfg.ipPort.String(),
			Handler:           http.HandlerFunc(newhealth.HandleHTTP),
			ReadHeaderTimeout: 0,
		}
		go func() {
			rlog.Info("Starting health server", rlog.Any("endpoint", cfg.ipPort.String()))
			err := httpServer.Serve(listener)
			if err != nil {
				rlog.Error("Failed to start health server", err)
			}
		}()
	}
	return nil
}

func MustStart(opts ...optionFunc) {
	if err := Start(opts...); err != nil {
		rlog.Error("Failed to start health server", err)
		os.Exit(1)
	}
}
func StartWithDefaults(opts ...optionFunc) error {
	opts = append(opts, ServerString(getHealthEndpoint()))
	return Start(opts...)
}

func MustStartWithDefaults(opts ...optionFunc) {
	if err := StartWithDefaults(opts...); err != nil {
		rlog.Error("Failed to start health server with defaults", err)
		os.Exit(1)
	}
}

func getHealthEndpoint() string {
	return fmt.Sprintf("%s:%s", rorconfig.GetString(rorconfig.HTTP_HEALTH_HOST), rorconfig.GetString(rorconfig.HTTP_HEALTH_PORT))
}
