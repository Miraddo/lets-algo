package main

import (
	"testing"
)

type List struct {
	nums   []int
	val    int
	result int
}

func TestRemoveElement(t *testing.T) {
	handle := []List{
		{
			nums:   []int{3, 2, 2, 3},
			val:    3,
			result: 2,
		},
		{
			nums:   []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:    2,
			result: 5,
		},
	}

	for _, x := range handle {

		res := removeElement(x.nums, x.val)

		if res != x.result {

			t.Errorf("RemoveElement was incorrect, got: %d, want: %d.", res, x.result)
		}

	}
}
