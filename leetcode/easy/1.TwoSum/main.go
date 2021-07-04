package main


// this problem is based : https://leetcode.com/problems/two-sum/

func TwoSum(nums []int, target int) []int {

	sMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		sMap[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		cmt := target - nums[i]

		if (sMap[cmt] != 0) && (sMap[cmt] != i) {
			return []int{i, sMap[cmt]}
		}
	}

	return []int{}
}


func main()  {

}


