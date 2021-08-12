package main

import (
	"fmt"
	"sort"
)
// findKthLargest : 215. Kth Largest Element in an Array | https://leetcode.com/problems/kth-largest-element-in-an-array/
// sort array first
// then get the index

func findKthLargest(nums []int, k int) int {
	// nums = sorts.HeapSort(nums) // leetcode 8ms
	// nums = sorts.QuickSort(nums) // leetcode 36ms

	sort.Ints(nums) // leetcode 4ms

	return nums[len(nums) - k]
}

func main()  {
	fmt.Println(findKthLargest([]int{3,2,1,5,6,4}, 2)) // return 5
	fmt.Println(findKthLargest([]int{3,2,3,1,2,4,5,5,6}, 4)) // return 4
}