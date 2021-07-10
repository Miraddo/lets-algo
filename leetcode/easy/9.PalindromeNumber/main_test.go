package main

import (
	"testing"
)

type List struct {
	num int
	result bool
}

func TestIsPalindrome(t *testing.T) {
	handle := []List{
		{
			num: 121,
			result: true,
		},
		{
			num: -121,
			result: false,
		},
		{
			num: 10,
			result: false,
		},
		{
			num: -101,
			result: false,
		},
	}

	for _, x := range handle {

		res := isPalindrome(x.num)

		if res != x.result  {

			t.Errorf("IsPalindrome was incorrect, got: %t, want: %t.", res, x.result)
		}

	}
}