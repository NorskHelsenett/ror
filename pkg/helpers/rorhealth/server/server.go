package server

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"net/netip"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

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
	// gate holds the optional liveness-escalation state. It is nil unless a
	// non-zero grace is configured, preserving the default behaviour where
	// liveness never inspects dependencies.
	gate *livenessGate
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

	// livenessGrace, when > 0, enables liveness escalation: /health/live fails
	// once a critical dependency has been failing for longer than this.
	livenessGrace    time.Duration
	livenessGraceSet bool
	livenessCritical []string
}

type optionFunc func(*config)

// LivenessEscalation enables liveness escalation: after a critical dependency
// has been continuously failing for longer than grace, /health/live starts
// returning 500 so Kubernetes restarts the pod. criticalChecks names the checks
// considered critical; passing none means every registered check is critical.
// A grace <= 0 disables escalation. Configuring this via option overrides the
// HTTP_HEALTH_LIVENESS_GRACE / HTTP_HEALTH_LIVENESS_CRITICAL_CHECKS env vars.
func LivenessEscalation(grace time.Duration, criticalChecks ...string) optionFunc {
	return optionFunc(func(cfg *config) {
		cfg.livenessGrace = grace
		cfg.livenessGraceSet = true
		cfg.livenessCritical = criticalChecks
	})
}

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
		if g := newLivenessGateFromConfig(cfg); g != nil {
			gate = g
			go g.run(context.Background())
			rlog.Info("liveness escalation enabled",
				rlog.String("grace", g.grace.String()),
				rlog.String("jitter", g.jitter.String()),
				rlog.Any("criticalChecks", g.criticalNames()))
		}

		listener, err := net.Listen("tcp", cfg.ipPort.String())
		if err != nil {
			return err
		}

		httpServer = &http.Server{
			Addr:              cfg.ipPort.String(),
			Handler:           newHealthMux(),
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

// newHealthMux builds the HTTP routing for the health server with separate
// liveness and readiness semantics:
//
//   - /health/live  liveness: reports the process is alive. It never runs the
//     dependency checks, so a dependency outage (e.g. vault unreachable) does
//     not cause Kubernetes to restart the pod.
//   - /health/ready readiness: runs all registered dependency checks. A failing
//     dependency makes the pod NotReady (removed from Service endpoints) while
//     it stays alive and queryable, allowing background retries to recover.
//   - /          and  /health  retain the previous behaviour (full dependency
//     check) for backwards compatibility with existing probes and callers.
func newHealthMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health/live", livenessHandler)
	mux.HandleFunc("/health/ready", newhealth.HandleHTTP)
	// Catch-all preserves the prior behaviour where any path returned the full
	// health check, including the well-known /health path.
	mux.HandleFunc("/", newhealth.HandleHTTP)
	return mux
}

// livenessHandler reports that the process is alive. By default it does not
// inspect any registered dependency checks: liveness must only fail when the
// process itself is unhealthy, never because an external dependency is down.
//
// When liveness escalation is configured (see LivenessEscalation /
// HTTP_HEALTH_LIVENESS_GRACE), it additionally fails once a critical dependency
// has been continuously unhealthy for longer than the configured grace, turning
// a chronically wedged dependency (e.g. a mongodb connection that never
// recovers) into a pod restart while still tolerating short outages.
func livenessHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet, http.MethodHead:
		if g := gate; g != nil {
			if name, dur, tripped := g.tripped(); tripped {
				rlog.Warn("liveness failing: critical dependency unhealthy beyond grace",
					rlog.String("check", name),
					rlog.String("failingFor", dur.String()),
					rlog.String("grace", g.grace.String()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
		w.WriteHeader(http.StatusOK)
	case http.MethodOptions:
		w.WriteHeader(http.StatusNoContent)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
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

// livenessSampleInterval is how often the gate re-evaluates dependency health.
const livenessSampleInterval = 15 * time.Second

// livenessGate tracks how long each critical dependency check has been
// continuously failing so livenessHandler can escalate to a pod restart once a
// failure persists beyond the configured grace.
type livenessGate struct {
	grace    time.Duration
	jitter   time.Duration
	critical map[string]struct{} // empty set means "all checks are critical"

	mu           sync.Mutex
	failingSince map[string]time.Time
}

// newLivenessGateFromConfig builds a gate from the explicit option, falling back
// to the HTTP_HEALTH_LIVENESS_* env vars. It returns nil (escalation disabled)
// when the resolved grace is not positive.
func newLivenessGateFromConfig(cfg *config) *livenessGate {
	grace := cfg.livenessGrace
	critical := cfg.livenessCritical
	if !cfg.livenessGraceSet {
		grace = parseGraceDuration(rorconfig.GetString(rorconfig.HTTP_HEALTH_LIVENESS_GRACE))
		critical = parseCriticalChecks(rorconfig.GetString(rorconfig.HTTP_HEALTH_LIVENESS_CRITICAL_CHECKS))
	}
	if grace <= 0 {
		return nil
	}

	set := make(map[string]struct{}, len(critical))
	for _, c := range critical {
		if c = strings.TrimSpace(c); c != "" {
			set[c] = struct{}{}
		}
	}

	return &livenessGate{
		grace:        grace,
		jitter:       livenessJitter(grace),
		critical:     set,
		failingSince: make(map[string]time.Time),
	}
}

func (g *livenessGate) criticalNames() []string {
	if len(g.critical) == 0 {
		return []string{"*"}
	}
	names := make([]string, 0, len(g.critical))
	for name := range g.critical {
		names = append(names, name)
	}
	return names
}

func (g *livenessGate) isCritical(name string) bool {
	if len(g.critical) == 0 {
		return true
	}
	_, ok := g.critical[name]
	return ok
}

// run periodically samples dependency health until the context is cancelled.
func (g *livenessGate) run(ctx context.Context) {
	g.sample(ctx)

	ticker := time.NewTicker(livenessSampleInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			g.sample(ctx)
		}
	}
}

// sample runs all registered checks and updates the per-check failing-since
// timestamps for the critical checks.
func (g *livenessGate) sample(ctx context.Context) {
	sampleCtx, cancel := context.WithTimeout(ctx, livenessSampleInterval)
	defer cancel()

	resp, err := newhealth.CheckNow(sampleCtx)
	if err != nil {
		// An inability to evaluate is not treated as a dependency failure:
		// liveness must not restart the pod on a transient evaluation error.
		return
	}

	now := time.Now()
	g.mu.Lock()
	defer g.mu.Unlock()

	for name, checks := range resp.Checks {
		if !g.isCritical(name) {
			continue
		}
		failing := false
		for _, c := range checks {
			if c.Status == newhealth.StatusFail {
				failing = true
				break
			}
		}
		if failing {
			if _, ok := g.failingSince[name]; !ok {
				g.failingSince[name] = now
			}
		} else {
			delete(g.failingSince, name)
		}
	}

	// Drop checks that are no longer registered so a deregistered dependency
	// cannot keep liveness tripped forever.
	for name := range g.failingSince {
		if _, ok := resp.Checks[name]; !ok {
			delete(g.failingSince, name)
		}
	}
}

// tripped reports whether any critical dependency has been failing longer than
// the grace (plus per-process jitter), returning the offending check name and
// how long it has been failing.
func (g *livenessGate) tripped() (string, time.Duration, bool) {
	g.mu.Lock()
	defer g.mu.Unlock()

	threshold := g.grace + g.jitter
	now := time.Now()
	for name, since := range g.failingSince {
		if d := now.Sub(since); d >= threshold {
			return name, d, true
		}
	}
	return "", 0, false
}

// parseGraceDuration accepts a Go duration ("5m") or a bare integer number of
// seconds. It returns 0 (escalation disabled) for empty or invalid input.
func parseGraceDuration(s string) time.Duration {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	if d, err := time.ParseDuration(s); err == nil {
		return d
	}
	if n, err := strconv.Atoi(s); err == nil {
		return time.Duration(n) * time.Second
	}
	rlog.Warn("invalid HTTP_HEALTH_LIVENESS_GRACE, disabling liveness escalation",
		rlog.String("value", s))
	return 0
}

func parseCriticalChecks(s string) []string {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}
	return strings.Split(s, ",")
}

// livenessJitter returns a small random offset (up to 10% of grace, capped at
// 30s) so a shared dependency outage does not restart every pod simultaneously.
func livenessJitter(grace time.Duration) time.Duration {
	maxJitter := grace / 10
	if maxJitter <= 0 {
		return 0
	}
	if maxJitter > 30*time.Second {
		maxJitter = 30 * time.Second
	}
	return time.Duration(rand.Int63n(int64(maxJitter) + 1))
}
