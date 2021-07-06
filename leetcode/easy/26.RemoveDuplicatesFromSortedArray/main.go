package main


// Based on : https://leetcode.com/problems/remove-duplicates-from-sorted-array

func RemoveDuplicates(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	var i int
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

func main() {

}
