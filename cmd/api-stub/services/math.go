package services

import (
	"math"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
	"github.com/NorskHelsenett/ror/pkg/models/providers"
)

func CpuPercentage(cpuCount int64, cpuConsumed int64) int64 {
	result := int64(0)
	if cpuCount <= 0 {
		return result
	}

	result = (cpuConsumed / 10) / cpuCount
	return result
}

func MemoryPercentage(memoryTotal int64, memoryConsumed int64) int64 {
	result := int64(0)
	if memoryTotal <= 0 {
		return result
	}

	result = (memoryConsumed * 100) / memoryTotal
	return result
}

func FindMachineClass(memoryAllocatedBytes int64, cpuAllocated int64, provider providers.ProviderType, prices []apicontracts.Price) (int64, string) {
	pricesCount := len(prices)
	if pricesCount == 0 {
		return -1, "Unknown"
	}
	memorygb := math.Round(float64(memoryAllocatedBytes) / float64(1050000000))
	var price apicontracts.Price
	for i := 0; i < pricesCount; i++ {
		p := prices[i]
		if p.Cpu == int(cpuAllocated) && p.Memory == int64(memorygb) && p.Provider == provider {
			price = p
			break
		}
	}

	if price.Price == 0 {
		return -1, price.MachineClass
	}

	return int64(price.Price), price.MachineClass
}
