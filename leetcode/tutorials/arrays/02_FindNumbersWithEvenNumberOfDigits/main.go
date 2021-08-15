package main

import (
	"fmt"
)

// Find Numbers With Even Number Of Digits
func findNumbers(nums []int) int {
	// is our main result
	var result int
	for _, x := range nums {

		// set result into ln variable
		// ln :=  len(strconv.Itoa(x)) // another way to count length of number
		ln := returnLength(x)

		// if ln is even so ln % 2 should be 0 otherwise is not even
		if ln%2 == 0 {
			result++
		}
	}

	return result
}

// return the length of numbers
func returnLength(x int) (res int) {
	for {
		x = x / 10
		if x != 0 {
			res++
		} else {
			res++
			return
		}
	}
}

func main() {
	fmt.Println(findNumbers([]int{22, 555, 1, 2, 0, 444, 66, 8888}))
}
