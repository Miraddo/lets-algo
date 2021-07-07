package Sorting

func InsertionSort(nums []int) []int {
	// we should start array from index 1 not 0
	for i := 1; i < len(nums) ; i++{
		// create a variable and put i inside it
		j := i

		// another array to checking each item with each other
		for j > 0 {
			// remmber j is equal with i so for default is start from index 1 for checking j with prev item we should minus it with 1 like j-1
			if nums[j-1] > nums[j]{
				// replace values
				nums[j-1] , nums[j] = nums[j], nums[j-1]
			}
			j--
		}
	}

	return nums
}
