package generator

import (
	"autoapi/graph/model"
)

func MapExpectationsToResources(performanceRequirements *model.PerformanceRequirements) *model.ResourceConfig {
	config := &model.ResourceConfig{
		MaxMemory:   "1Gi",
		MinMemory:   "100Mi",
		MaxCPU:      "1",
		MinCPU:      "100m",
		MinReplicas: 1,
	}

	if performanceRequirements.APIUsageFrequency == "HIGH" {
		config.MaxCPU = "2"
	}

	if performanceRequirements.RequestVolume == "LARGE" {
		config.MaxMemory = "2Gi"
		config.MinReplicas = 3
	}

	if performanceRequirements.HighAvailability {
		config.MinReplicas = 3
	}

	if performanceRequirements.BatchLoad {
		config.MaxMemory = "4Gi"
		config.MaxCPU = "4"
	}

	return config
}
