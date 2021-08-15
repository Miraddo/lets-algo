package main

import (
	"fmt"
)

func sortedSquares(nums []int) []int {
	// create another variable with same length of nums
	result := make([]int, len(nums))

	// create end and start variable to select last and first index of array
	// using two pointer

	end, start := len(nums)-1, 0

	// if end is 0 it means array has no value
	if len(nums) == 0 {
		return nums
	}

	// count from last index to the first inside array
	for i := len(nums) - 1; i >= 0; i-- {
		// check last value and first value of array
		fmt.Println(nums[end]*nums[end], nums[start]*nums[start])
		if nums[end]*nums[end] > nums[start]*nums[start] {
			result[i] = nums[end] * nums[end]
			end--
		} else {
			result[i] = nums[start] * nums[start]
			start++
		}
	}

	return result
}

func main() {
	// fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
	fmt.Println(sortedSquares([]int{-1}))
}
