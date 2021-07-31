package main

import "testing"

type List struct {
	nums    []int
	result bool
}

func TestCanJump(t *testing.T)  {
	handle := []List{
		{
			nums:    []int{2, 3, 1, 1, 4},
			result: true,
		},
		{
			nums:    []int{3, 2, 1, 0, 4},
			result: false,
		},
		{
			nums:    []int{1, 3, 0, 0, 2},
			result: true,
		},
		{
			nums:    []int{3, 0, 2, 0, 1, 2, 5},
			result: true,
		},
		{
			nums:    []int{4, 0, 0, 0, 1, 0, 0, 5, 6},
			result: false,
		},
	}

	for _, x := range handle {

		res := canJump(x.nums)

		if res != x.result {

			t.Errorf("Can Jump was incorrect, got: %t, want: %t.", res, x.result)
		}

	}
}
