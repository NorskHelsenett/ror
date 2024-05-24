package clusternamehelper

import "strconv"

func GetClusterNameOfArray(hostnameArray []string) string {
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
