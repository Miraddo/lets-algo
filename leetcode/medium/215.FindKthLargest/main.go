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
	//return mapWay(nums, k)
}

//func mapWay (nums []int, k int) int {
//	val := make(map[int]int)
//	val[math.MinInt64] = math.MaxInt64
//	var res []int
//	count := 1
//	for i:=0 ; i < k -1 ; i++{
//		for j := range nums{
//			if nums[j] >= val[count - 1]{
//				val[0] = nums[j]
//				count++
//			}
//			nums[j] = -1
//		}
//	}
//	for i, x := range val{
//		if i > 1{
//			res = append(res, x)
//		}
//	}
//
//}

func main()  {
	fmt.Println(findKthLargest([]int{3,2,1,5,6,4}, 2)) // return 5
	fmt.Println(findKthLargest([]int{3,2,3,1,2,4,5,5,6}, 4)) // return 4
}