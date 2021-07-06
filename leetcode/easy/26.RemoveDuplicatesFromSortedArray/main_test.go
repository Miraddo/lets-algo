package main

import (
	"testing"
)

type List struct {
	nums []int
	result int
}

func TestRemoveDuplicates(t *testing.T) {
	handle := []List{
		{
			nums: []int{0,0,1,1,1,2,2,3,3,4},
			result: 5,
		},
		{
			nums: []int{1,1,2},
			result: 2,
		},
	}

	for _, x := range handle {

		res := RemoveDuplicates(x.nums)

		if res != x.result  {

			t.Errorf("Reverse was incorrect, got: %d, want: %d.", res, x.result)
		}

	}
}