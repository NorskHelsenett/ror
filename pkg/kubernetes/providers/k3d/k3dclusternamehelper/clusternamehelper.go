package k3dclusternamehelper

import (
	"strconv"
	"strings"
)

// GetClusternameFromHostname returns the clustername from a hostname on one of the formats:
// k3d-<clustername>-agent-<number>
// k3d-<clustername>-server-<number>
// It also gracefully handles control-plane variants and load-balancer nodes.
func GetClusternameFromHostname(hostname string) string {
	if hostname == "" {
		return ""
	}

	parts := strings.Split(hostname, "-")
	if len(parts) == 0 {
		return ""
	}

	start := 0
	if parts[0] == "k3d" {
		start = 1
	}

	end := len(parts)
	if end <= start {
		return ""
	}

	if _, err := strconv.Atoi(parts[end-1]); err == nil {
		end--
	}
	if end <= start {
		return ""
	}

	var builder strings.Builder
	for i := start; i < end; i++ {
		part := parts[i]
		if part == "" {
			continue
		}
		switch part {
		case "agent", "server", "serverlb", "control", "plane", "controlplane":
			return builder.String()
		}
		if builder.Len() > 0 {
			builder.WriteString("-")
		}
		builder.WriteString(part)
	}

	return builder.String()
}
