package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(v string) []string {
	basket := map[string]uint{}
	for _, key := range strings.Fields(v) {
		if key == "-" {
			continue
		}
		last := string(key[len(key)-1])
		if last == "!" || last == "," || last == "." || last == "-" {
			key = key[:len(key)-1]
		}
		str := strings.ToLower(key)
		basket[str]++
	}
	keys := make([]string, 0, len(basket))
	for key := range basket {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		if basket[keys[i]] == basket[keys[j]] {
			return strings.Compare(keys[j], keys[i]) == 1
		}
		return basket[keys[i]] > basket[keys[j]]
	})
	if len(keys) >= 10 {
		return keys[:10]
	}
	return keys
}
