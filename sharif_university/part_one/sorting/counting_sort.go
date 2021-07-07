package Sorting

func CountingSort(nums []int) []int{

	var max int
	ln := len(nums)
	j := 1

	for i := range nums{
		if nums[i] > nums[j+1]{
			j++
			max = nums[i]
		}
	}

	count := make([]int, max+1)

	for _, x := range nums{
		count[x] += 1
	}

	j = 1

	for i:=1; i< len(count) - 1; i++{

		count[j+1] = count[i] + count[j+1]
		j++
	}


	result := make([]int, ln)

	for _, x :=range nums{

		result[count[x]-1] = x
		count[x] = count[x] - 1

	}

	return result
}
