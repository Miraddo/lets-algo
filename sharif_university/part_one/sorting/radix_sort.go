package Sorting

// RadixSort Radix Sort Algorithm
// Range numbers 0 < x < 9999

func RadixSort(nums []int) []int {

	var list [][]int = make([][]int, 10)
	var f, s int

	for i := 0; i < len(nums); i++ {

		switch i {
		case 0:
			f, s = 10, 1 // number : (2548 % 10) / 1 = 8
		case 1:
			f, s = 100, 10 // number : (2548 % 100) / 10 = 4
		case 2:
			f, s = 1000, 100 // number : (2548 % 1000) / 100 = 5
		case 3:
			f, s = 10000, 1000 // number : (2548 % 10000) / 1000 = 2
		}

		// add number to each array inside list
		for _, x := range nums {
			key := (x % f) / s

			list[key] = append(list[key], x)
		}

		nums = []int{}
		// join arrays into nums
		for j := range list {

			for _, x := range list[j] {
				nums = append(nums, x)
			}

			list[j] = []int{}
		}

	}

	return nums
}