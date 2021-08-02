package main

import (
	"sort"
)

// frequencySort  451. Sort Characters By Frequency
// based on : https://leetcode.com/problems/sort-characters-by-frequency/
func frequencySort(s string) string {
	fs := make([]int, 255)

	for i := 0; i < len(s); i++ {
		fs[s[i]]++
	}

	b := []byte(s)

	sort.Slice(b, func(i, j int) bool {
		if fs[b[i]] == fs[b[j]] {
			return b[i] > b[j]
		}
		return fs[b[i]] > fs[b[j]]
	})

	return string(b)
}
