package main

import "fmt"

func setZeroes(matrix [][]int) [][]int {

	var arr []int

	for i, x := range matrix{
		for j := range x{
			if x[j] == 0{
				arr = append(arr, i)
			}
		}
	}

	for _, x := range matrix{
		for j := range x{
			if x[j] == 0{

				for _, d := range matrix{
					d[j] = 0
				}
			}
		}
	}

	for _, d := range arr{

		for  k := range matrix[d]{
			matrix[d][k] = 0
		}
	}

	fmt.Println(arr)
	return matrix
}

func main()  {
	fmt.Println(setZeroes([][]int{{1,1,1}, {1,0,1}, {1,1,1}}))
	fmt.Println(setZeroes([][]int{{0,1,2,0}, {3,4,5,2}, {1,3,1,5}}))
}
