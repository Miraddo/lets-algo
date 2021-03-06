package main

// maxArea find max water in area | 11. Container With Most Water
// based on : https://leetcode.com/problems/container-with-most-water/
// Two Pointer *
func maxArea(height []int) int {
	lh, res, j := len(height)-1, 0, 0

	for i := 0; i < len(height); i++ {
		if j < i {
			if height[j] < height[lh] {
				if res < height[j]*(lh-j) {
					res = height[j] * (lh - j)
				}
				j++
			} else {
				if res < height[lh]*(lh-j) {
					res = height[lh] * (lh - j)
				}
				lh--
			}
		}
	}

	return res
}
