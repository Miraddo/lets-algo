package Sorting


// BubbleSort Bubble Sort Algorithm
func BubbleSort(nums []int) []int{
	for i:=0; i < len(nums)-1 ; i++{

		for j:=0; j < len(nums)-i-1 ; j++{
			if nums[j] > nums[j+1]{
				// swap in golang
				nums[j+1], nums[j] = nums[j], nums[j+1]
			}
		}
	}

	return  nums
}

