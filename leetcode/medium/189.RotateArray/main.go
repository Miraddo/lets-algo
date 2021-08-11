package main

// 189. Rotate Array : https://leetcode.com/problems/rotate-array/
// rotate
// another way :
func rotate(nums []int, k int) {
	n := len(nums)

	pos(nums, 0, n-1)
	pos(nums, 0, k-1)
	pos(nums, k, n-1)

}

func pos(nums []int, s int, e int) {
	for i := 0; s < e; i++ {
		nums[s], nums[e] = nums[e], nums[s]
		s, e = s+1, e-1
	}
}


// another why
// is not efficient why, but it can be a fun why to learn how to work with memory
//func rotate(nums []int, k int) {
//	kk := len(nums) - k
//	if kk < 0{
//		kk = len(nums) - kk
//		if kk > len(nums) {
//			kk = kk % 2
//			if kk == 0 {
//				kk++
//			}
//		}
//	}
//
//	copy(nums, append(nums[kk:], nums[:kk]...))
//
//}



func main()  {
	rotate([]int{1,2,3,4,5,6,7}, 3)
	rotate([]int{1,2,3,4,5,6,7}, 2)
	rotate([]int{1,2,3,4,5,6,7}, 1)
	rotate([]int{1,2,3,4,5,6,7}, 0)
	rotate([]int{-1,-100,3,99}, 2)
}
