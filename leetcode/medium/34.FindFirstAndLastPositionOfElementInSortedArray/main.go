package main

import (
	"fmt"
)
// 34. Find First and Last Position of Element in Sorted Array
// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

// this way is better it should be O(log n)
func searchRange(nums []int, target int) []int {

	first := binarySearch(nums, target, true)
	last := binarySearch(nums, target, false)

	return []int{first, last}
}

func binarySearch(nums []int, target int, pn bool) int {
	left, right := 0, len(nums) -1
	res := -1
	for left <= right{
		mid := (left + right) / 2

		if target > nums[mid]{
			left = mid + 1
		}else if target < nums[mid]{
			right = mid -1
		}else{
			res = mid
			if pn {
				right = mid -1
			}else {
				left = mid + 1
			}
		}
	}
	return res
}


// is take O(n) is not what we want
//func searchRange(nums []int, target int) []int {
//	var result []int
//	var lastIndex int
//	firstIndex := LinearSearch(nums, target)
//	if firstIndex == -1{
//		return []int{-1, -1}
//	}
//
//	lastIndex = findLastIndex(nums[firstIndex:], target)
//
//	result = append(result, firstIndex)
//	result = append(result, firstIndex + lastIndex)
//
//	return  result
//}
//
//func findLastIndex(last []int, target int) int {
//	var res int
//	for i, x := range last{
//		if x == target{
//			res = i
//		}
//	}
//	return res
//}
//
//func LinearSearch(nums []int, target int) int {
//	for i, x := range nums {
//		if x == target {
//			return i
//		}
//	}
//	return -1
//}

func main()  {
	fmt.Println(searchRange([]int{5,7,7,8,8,10}, 8))
	fmt.Println(searchRange([]int{5,7,7,8,8,10}, 6))
	fmt.Println(searchRange([]int{5,7,7,8,8,10}, 10))
	fmt.Println(searchRange([]int{5,7,7,8,8,10}, 5))
	fmt.Println(searchRange([]int{5,7,7,8,8,10}, 7))
	fmt.Println(searchRange([]int{5,7,7,8,8,9,9,10}, 9))
	fmt.Println(searchRange([]int{5,7,7,8,8,9,9,9,9,10}, 9))
	fmt.Println(searchRange([]int{}, 0))
	fmt.Println(searchRange([]int{1}, 1))
	fmt.Println(searchRange([]int{2,2}, 2))
	fmt.Println(searchRange([]int{1,3}, 3))
	fmt.Println(searchRange([]int{1,3}, 1))
}
