package main

import (
	"fmt"
)

// lengthOfLongestSubstring : Longest Substring Without Repeating Characters
// base on : https://leetcode.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	arr := [128]int{}

	l, r, m := 0, 0, 0

	for r < len(s) {
		c := s[r]
		r++
		arr[c]++
		for arr[c] > 1 {
			d := s[l]
			l++
			arr[d]--
		}

		if r-l > m {
			m = r - l
		}
	}

	return m
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring(""))
}
