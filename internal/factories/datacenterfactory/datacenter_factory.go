package datacenterfactory

func DatacenterToTanzu(dc string) string {
	switch dc {
	case "trd1cl02", "trd1cl2", "trd1-cl02", "trd01-cl02":
		return "trd1-cl02"
	case "trd1", "trd1-cl01", "trd01-cl01":
		return "trd1"
	case "osl1", "osl1-cl01", "osl01-cl01":
		return "osl1"
	default:
		return dc
	}
}
