package main

import (
	"testing"
)

type List struct {
	num int
	result int
}

func TestReverse(t *testing.T) {
	handle := []List{
		{
			num: 123,
			result: 321,
		},
		{
			num: -123,
			result: -321,
		},
		{
			num: 120,
			result: 21,
		},
		{
			num: 0,
			result: 0,
		},
	}

	for _, x := range handle {

		res := Reverse(x.num)

		if res != x.result  {

			t.Errorf("Reverse was incorrect, got: %d, want: %d.", res, x.result)
		}

	}
}