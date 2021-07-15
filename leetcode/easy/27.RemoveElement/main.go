package main

// RemoveElement based on : https://leetcode.com/problems/remove-element
func removeElement(nums []int, val int) int {
	var i int
	for _ , x := range nums{
		if x != val{
			nums[i] = x
			i++
		}
	}

	return i
}
