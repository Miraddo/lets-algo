package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	result, counter := 0, 0

	for _, x := range nums{

		if x == 1{
			counter++
			if counter > result {
				result = counter
			}
		}else{

			counter = 0
		}
	}

	return result
}


func main()  {
	fmt.Println(findMaxConsecutiveOnes([]int{1,1,1,1,0,1,1,0,1, 0,0,0,0,1,1,1,1,0}))
	fmt.Println(findMaxConsecutiveOnes([]int{1,1,0,1,1,1}))
}
