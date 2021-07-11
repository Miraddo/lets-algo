package Sorting

import (
	"math"
)

//BucketSort Bucket Sort Algorithm base on
func BucketSort(nums []int) []int {

	ln := len(nums)
	var result []int
	var storeList [][]int = make([][]int, ln)

	// put numbers into buckets
	for _, x := range nums {

		j := int(math.Floor(float64(x / 10)))

		storeList[j] = append(storeList[j], x)
	}

	// sort each array with bubble sort algorithm or other algorithm that we have
	for _, x := range storeList {
		if len(x) > 1 {
			list := BubbleSort(x)
			// save result
			result = append(result, list...)
		}
	}

	return result
}
