package config

import "strings"

func ParseFeatures(input string) map[string]bool {
	features := make(map[string]bool)
	items := strings.Split(input, ",")
	for _, item := range items {
		feature := strings.TrimSpace(item)
		if feature != "" {
			features[feature] = true
		}
	}
	return features
}
