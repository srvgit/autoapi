package config

import "strings"

func ParseFeatures(input string) map[string]bool {
	features := make(map[string]bool)
	pairs := strings.Split(input, ",")
	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		if len(kv) == 2 {
			features[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1]) == "true"
		}
	}
	return features
}
