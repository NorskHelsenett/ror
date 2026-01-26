package kindclusternamehelper

import (
	"strings"
	"unicode"
)

// getClusterNameOfArray returns the clustername from a hostname on one of the formats:
// <clustername>-control-plane
// <clustername>-control-plane2
// <clustername>-worker
// <clustername>-worker2
func GetClusternameFromHostname(hostname string) string {
	if hostname == "" {
		return ""
	}

	if cluster, ok := trimSuffixWithDigits(hostname, "-control-plane"); ok {
		return cluster
	}
	if cluster, ok := trimSuffixWithDigits(hostname, "-worker"); ok {
		return cluster
	}

	return hostname
}

func trimSuffixWithDigits(hostname, suffix string) (string, bool) {
	idx := strings.LastIndex(hostname, suffix)
	if idx == -1 {
		return "", false
	}

	rest := hostname[idx+len(suffix):]
	for _, r := range rest {
		if !unicode.IsDigit(r) {
			return "", false
		}
	}

	if idx == 0 {
		return "", true
	}

	return hostname[:idx], true
}
