package binary_search

import "testing"

type List struct {
	nums   []int
	val    int
	result int
}

func TestBinarySearch(t *testing.T) {
	var listResult = []List{
		{
			nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 99, 100, 101},
			val:    100,
			result: 9,
		},
		{
			nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 99, 100, 101},
			val:    99,
			result: 8,
		},
		{
			nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 99, 100, 101},
			val:    1,
			result: 0,
		},
		{
			nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 99, 100, 101},
			val:    155,
			result: -1,
		},
	}

	// For Iteration Binary Search
	t.Run("Iteration Binary Search", func(t *testing.T) {
		for _, x := range listResult {
			result := binarySearch(x.nums, x.val)

			if result != x.result {
				t.Errorf("Iteration Binary Search should be %d, but got %d", x.result, result)
			}

		}
	})

	// For Recursive Binary Search
	t.Run("Recursive Binary Search", func(t *testing.T) {
		for _, x := range listResult {

			result := binarySearchRec(x.nums, x.val, len(x.nums)-1, 0)

			if result != x.result {
				t.Errorf("Recursive Binary Search should be %d, but got %d", x.result, result)
			}
		}
	})
}
