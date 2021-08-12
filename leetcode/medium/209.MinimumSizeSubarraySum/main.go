package main

import (
	"math"
)

// 209. Minimum Size Subarray Sum | Two Pointer | https://leetcode.com/problems/minimum-size-subarray-sum/
// minSubArrayLen approach one
func minSubArrayLen(target int, nums []int) int {
	result := float64(math.MaxUint64)
	left, sum := 0 , 0

	for i := range nums{
		sum += nums[i]
		for sum >= target{
			result = math.Min(result, float64(i + 1 - left))
			sum -= nums[left]
			left++
		}
	}

	if result != math.NaN(){
		if int(result) < 0 {
			return 0
		}
		return int(result)
	}

	return 0
}


// minSubArrayLen2 better and cleaner way to handle this problem
func minSubArrayLen2(target int, nums []int) int {
	left, right, result := 0 , 0, 0.0

	temp := nums[0]

	for right < len(nums){
		if temp >= target{
			if result == 0{
				result = float64(right - left + 1)
			}else {
				result= math.Min(result, float64(right - left + 1))
			}

			temp -= nums[left]
			left++
		}else{
			right++
			if right < len(nums){
				temp += nums[right]
			}
		}
	}

	return int(result)
}

