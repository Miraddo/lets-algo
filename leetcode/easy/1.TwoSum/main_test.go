package main

import (
	"reflect"
	"testing"
)

type List struct {
	Nums   []int
	Target int
	Result []int
}

func TestTwoSum(t *testing.T) {

	handle := []List{
		{
			Nums:   []int{2, 7, 11, 15},
			Target: 9,
			Result: []int{0, 1},
		},

		{
			Nums:   []int{3, 2, 4},
			Target: 6,
			Result: []int{1, 2},
		},

		{
			Nums:   []int{3, 3},
			Target: 6,
			Result: []int{0, 1},
		},
	}

	for _, x := range handle {

		res := TwoSum(x.Nums, x.Target)

		if !reflect.DeepEqual(res, x.Result)  {

			t.Errorf("TwoSum was incorrect, got: %d, want: %d.", res, x.Result)
		}

	}

}
