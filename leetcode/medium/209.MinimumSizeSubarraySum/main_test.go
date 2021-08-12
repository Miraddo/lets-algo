package main

import (
	"encoding/json"
	"log"
	"testing"

	_ "embed"
)

type Result struct {
	nums []int
	target int
	expected int
}


//go:embed testarr.txt
var b []byte

func TestMinSubArrayLen(t *testing.T)  {

	var ints []int
	err := json.Unmarshal([]byte(b), &ints)
	if err != nil {
		log.Fatal(err)
	}


	check := []Result{

		{
			nums: ints,
			target: 120331635,
			expected: 2327,
		},

		{
			nums: []int{2,3,1,2,4,3},
			target: 7,
			expected: 2,
		},

		{
			nums: []int{1,4,4},
			target: 4,
			expected: 1,
		},

		{
			nums: []int{1,1,1,1,1,1,1,1},
			target: 11,
			expected: 0,
		},
		{
			nums: []int{1,2,3,4,5},
			target: 11,
			expected: 3,
		},

		{
			nums: []int{5,1,3,5,10,7,4,9,2,8},
			target: 15,
			expected: 2,
		},
	}

	for _, x := range check{

		t.Run("Approach One", func(t *testing.T) {
			result := minSubArrayLen(x.target, x.nums)

			if result != x.expected{
				t.Errorf("Length of sub array is not correct expected, %d but got %d \n %v", x.expected, result, x)
			}
		})

		t.Run("Approach Two", func(t *testing.T) {
			result := minSubArrayLen2(x.target, x.nums)

			if result != x.expected{
				t.Errorf("Length of sub array is not correct expected, %d but got %d \n %v", x.expected, result, x)
			}
		})
	}

}

