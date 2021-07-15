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
			nums: []int{1,2,3,4,5,6,7,8,99,100,101},
			val: 100,
			result: 9,
		},
		{
			nums: []int{1,2,3,4,5,6,7,8,99,100,101},
			val: 99,
			result: 8,
		},
		{
			nums: []int{1,2,3,4,5,6,7,8,99,100,101},
			val: 1,
			result: 0,
		},
	}



	for _, x := range listResult{
		res := binarySearch(x.nums, x.val)

		if res != x.result {
			t.Errorf("Binary Search should be %d, but got %d", x.result, res)
		}
	}
}
