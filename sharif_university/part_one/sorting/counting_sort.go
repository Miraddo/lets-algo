package Sorting

// CountingSort Counting Sort Algorithm
func CountingSort(nums []int) []int{

	var max int
	ln := len(nums)
	j := 1
	// find max number inside array
	for i := range nums{
		if nums[i] > nums[j+1]{
			j++
			max = nums[i]
		}
	}

	// create an array with max+1 size
	count := make([]int, max+1)
	// with x value we add 1 to index x inside count array
	for _, x := range nums{
		count[x] += 1
	}

	// now we looking for each key and each key has one or more value
	j = 1
	for i:=1; i< len(count) - 1; i++{

		count[j+1] = count[i] + count[j+1]
		j++
	}

	// then we create another array with the same size as nums
	result := make([]int, ln)
	for _, x :=range nums{

		result[count[x]-1] = x
		count[x] = count[x] - 1

	}

	return result
}
