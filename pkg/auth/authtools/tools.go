package authtools

import (
	"fmt"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	UserLookupHistogram       *prometheus.HistogramVec
	ServerConnectionHistogram *prometheus.HistogramVec
	ServerReconnectCounter    *prometheus.CounterVec
)

func init() {
	ServerConnectionHistogram = promauto.NewHistogramVec(prometheus.HistogramOpts{Name: "auth_server_connection_duration_seconds", Help: "Duration of server connection in seconds"}, []string{"provider", "domain", "host", "port", "status"})
	UserLookupHistogram = promauto.NewHistogramVec(prometheus.HistogramOpts{Name: "auth_user_lookup_duration_seconds", Help: "Duration of user lookup in seconds"}, []string{"provider", "domain", "status"})
	ServerReconnectCounter = promauto.NewCounterVec(prometheus.CounterOpts{Name: "auth_server_reconnects_total", Help: "Total number of server reconnects"}, []string{"provider", "domain"})
}

func SplitUserId(userId string) (string, string, error) {
	parts := strings.Split(userId, "@")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid userId: %s", userId)
	}
	return parts[0], parts[1], nil
}
