package binary_search

// binarySearch is an algorithm for searching inside an sorted array and find the index of value then return it
func binarySearch(nums []int, val int) (mid int ) {

	// up is equal with length of array nums
	// down is 0
	// mid is equal with (up + down) / 2
	var up, down int = len(nums) - 1, 0
	//
	for range nums {
		mid = (up + down) / 2

		if nums[mid] == val {
			return
		}

		if nums[mid] < val {
			// set down if nums[mid] value is smaller then val
			down = mid + 1
		}

		if nums[mid] > val {
			// set up if nums[mid] value is greater then val
			up = mid - 1
		}

	}

	return
}