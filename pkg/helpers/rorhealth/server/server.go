package server

import (
	"fmt"
	"net"
	"net/http"
	"net/netip"
	"strconv"
	"sync"

	newhealth "github.com/dotse/go-health"
)

const (
	// Port is the default port for checking health over HTTP.
	Port = 9_999
)

//nolint:gochecknoglobals
var (
	httpServer *http.Server
	initMtx    sync.Mutex
)

type ServerParams struct {
	Server string
}

func ServerString(serverstring string) Option {
	return optionFunc(func(cfg *config) {
		var err error
		cfg.ipPort, err = netip.ParseAddrPort(serverstring)
		if err != nil {
			// handle strings with only port or ip
			var serverIp netip.Addr
			var serverPort uint16
			server, port, err := net.SplitHostPort(serverstring)
			if err != nil {
				// check if serverstring is an ip address
				serverIp, _ = netip.ParseAddr(serverstring)
				if serverIp.IsValid() {
					server = serverstring
				} else {
					// check if serverstring is a port
					_, err := strconv.ParseUint(serverstring, 10, 16)
					if err == nil {
						port = serverstring
					}
				}
			}
			if server == "" {
				serverIp = netip.IPv4Unspecified()
			}

			if port == "" {
				serverPort = Port
			} else {
				port64, err := strconv.ParseUint(port, 10, 16)
				if err != nil {
					fmt.Println("Error parsing port string: ", err)
				}
				serverPort = uint16(port64)

			}
			cfg.ipPort = netip.AddrPortFrom(serverIp, serverPort)
			if !cfg.ipPort.IsValid() {
				fmt.Println("Error parsing server string: ", err)
			}
		}
	})
}

type Option interface {
	apply(*config)
}

type config struct {
	ipPort netip.AddrPort
}

type optionFunc func(*config)

func (of optionFunc) apply(cfg *config) { of(cfg) }

func Start(opts ...Option) error {
	initMtx.Lock()
	defer initMtx.Unlock()
	cfg := &config{
		ipPort: netip.AddrPortFrom(netip.IPv4Unspecified(), Port),
	}

	for _, o := range opts {
		o.apply(cfg)
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
			_ = httpServer.Serve(listener)
		}()
	}

	return nil
}
