package main

import (
"testing"
)

type List struct {
	roman string
	result int
}

func TestRomanToInterger(t *testing.T) {
	handle := []List{
		{
			roman: "III",
			result: 3,
		},
		{
			roman: "IV",
			result: 4,
		},
		{
			roman: "LVIII",
			result: 58,
		},
		{
			roman: "MCMXCIV",
			result: 1994,
		},

	}

	for _, x := range handle {

		res := RomanToInterger(x.roman)

		if res != x.result  {

			t.Errorf("RomanToInterger was incorrect, got: %d, want: %d.", res, x.result)
		}

	}
}