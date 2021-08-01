package main

import (
	"encoding/json"
	"log"
	"testing"

	_ "embed"
)

//go:embed test.txt
var b []byte


type handler struct {
	height  []int
	result     int
}

func TestMaxArea(t *testing.T)  {


	var ints []int
	err := json.Unmarshal([]byte(b), &ints)
	if err != nil {
		log.Fatal(err)
	}

	listData := []handler{
		{
			height:  ints,
			result:  402471897,
		},
		{
			height:  []int{2,3,4,5,18,17,6},
			result:  17,
		},
		{
			height:  []int{1,8,6,2,5,4,8,3,7},
			result:  49,
		},
		{
			height:  []int{3,9,3,4,7,2,12,6},
			result:  45,
		},
		{
			height:  []int{4,3,2,1,4},
			result:  16,
		},
		{
			height:  []int{1,1},
			result:  1,
		},
		{
			height:  []int{1,2,1},
			result:  2,
		},
	}


	for _, x := range listData {

		res := maxArea(x.height)

		if res != x.result {
			t.Errorf("Max Area was incorrect, got: %v, want: %v.", res, x.result)
		}

	}
}