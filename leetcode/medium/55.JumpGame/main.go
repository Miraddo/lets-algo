package main

// canJump : 55.JumpGame. Jump Game
// based on = https://leetcode.com/problems/jump-game/
func canJump(nums []int) bool {
	var res int

	for i := range nums {
		x := i + nums[i]

		if res < i {
			return false
		}

		if res < x {
			res = x
		}
	}

	return true
}
