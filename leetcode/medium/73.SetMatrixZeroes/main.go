package main
// 73. Set Matrix Zeroes based on : https://leetcode.com/problems/set-matrix-zeroes/

// setZeroes my Code
func setZeroes(matrix [][]int)  {
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
}

// setZeroesLC LeetCode Solution
func setZeroesLC(matrix [][]int){
	isCol := false
	row := len(matrix)
	col := len(matrix[0])

	for i:=0; i < row; i++{
		if matrix[i][0] == 0{
			isCol = true
		}

		for j:=1; j < col; j++{
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	for i:=1; i < row; i++{
		for j:=1; j < col; j++{
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if matrix[0][0] == 0{
		for j:=0; j < col; j++{
			matrix[0][j] = 0
		}
	}

	if isCol{
		for i:=0; i < row; i++{
			matrix[i][0] = 0
		}
	}

}

func main()  {
	//fmt.Println(setZeroes([][]int{{1,1,1}, {1,0,1}, {1,1,1}}))
	//fmt.Println(setZeroes([][]int{{0,1,2,0}, {3,4,5,2}, {1,3,1,5}}))
}
