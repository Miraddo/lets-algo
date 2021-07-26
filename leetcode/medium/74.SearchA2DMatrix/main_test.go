package main

import (
	"testing"
)

type List struct {
	num    [][]int
	target int
	result bool
}

func TestSearchMatrix(t *testing.T) {
	handle := []List{
		{
			num:    [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 3,
			result: true,
		},
		{
			num:    [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 13,
			result: false,
		},
	}

	for _, x := range handle {

		res := searchMatrix(x.num, x.target)

		if res != x.result {

			t.Errorf("Search Matrix was incorrect, got: %t, want: %t.", res, x.result)
		}

	}
}
