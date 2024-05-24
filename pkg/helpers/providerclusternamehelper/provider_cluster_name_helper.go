package providerclusternamehelper

import (
	"strconv"
	"strings"
)

func GetKindClustername(hostname string) string {
	return getClusterNameOfArray(hostname)
}

func GetK3dClustername(hostname string) string {
	return getClusterNameOfArray(hostname)
}

func getClusterNameOfArray(hostname string) string {
	hostnameArray := strings.Split(hostname, "-")
	lastblock := hostnameArray[len(hostnameArray)-1]
	var clusterName string
	var length int
	if _, err := strconv.Atoi(lastblock); err == nil {
		length = len(hostnameArray) - 2

	} else {
		length = len(hostnameArray) - 1
	}

	var separator string
	for i := 0; i < length; i++ {
		if hostnameArray[i] == "control" || hostnameArray[i] == "plane" {
			break
		}
		if len(clusterName) > 0 {
			separator = "-"
		}
		clusterName = clusterName + separator + hostnameArray[i]
	}
	return clusterName
}
