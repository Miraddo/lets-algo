package Sorting

// SelectionSort Selection Sort Algorithm Function
func SelectionSort(nums []int) []int {

	// count length of array
	ln := len(nums)

	// a for statment start counting inside array
	for i := range nums {
		index := i // added key to index variable
		// another for statement to find minimum number inside array be careful j should be start from i value not 0 or 1 ...
		for j := i; j < ln; j++ {
			if nums[j] < nums[index] {
				index = j // if we found smallest number we save the key inside index variable
			}
		}

		// finally we swap values with the keys
		nums[i], nums[index] = nums[index], nums[i]
	}

	return nums

}
