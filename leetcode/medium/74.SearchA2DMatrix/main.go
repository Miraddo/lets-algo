package main

// searchMatrix : Search a 2D Matrix | solved with binary search algorithm
// based on : https://leetcode.com/problems/search-a-2d-matrix/
func searchMatrix(matrix [][]int, target int) bool {
	var mid, left, right int

	for _, x := range matrix {
		if x[len(x)-1] == target {
			return true
		}

		if x[len(x)-1] > target {

			left = len(x) - 1

			for range x {
				mid = (left + right) / 2

				if x[mid] == target {
					return true
				}

				if x[mid] < target {
					right = mid + 1
				}

				if x[mid] > target {
					left = mid - 1
				}
			}

			return false
		}
	}

	return false
}
