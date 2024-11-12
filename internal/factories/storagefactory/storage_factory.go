package storagefactory

func GetStorageClassByDatacenter(datacenter string) string {
	switch datacenter {
	case "trd1-cl01", "trd1cl01", "trd1":
		return "trd1-w02-cl01-vsan-storage-policy"
	case "trd1-cl02", "trd1cl02":
		return "trd1-w02-vc1-fc-san"
	case "osl1-cl01", "osl1cl01", "osl1":
		return "osl1-w02-cl01-vsan-storage-policy"
	default:
		return ""
	}
}
