package main

import "fmt"

func frequencySort(s string) string {
	var hm = make(map[byte]int)
	var keys []int

	for i:=0; i<len(s); i++{
		hm[s[i]]++
	}

	for k := range hm {
		keys = append(keys, hm[k])
	}

	keys = BubbleSort(keys)
	str := ""
	for _, x := range keys{
		for d := range hm{
			if x == hm[d]{
				i :=0
				for i < x{
					str += string(d)
					i++
				}
				hm[d] = -1
			}
		}
	}


	return str
}

func BubbleSort(nums []int) []int{
	for i:=0; i < len(nums)-1 ; i++{

		for j:=0; j < len(nums)-i-1 ; j++{
			if nums[j] < nums[j+1]{
				// swap in golang
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	return  nums
}


func main(){
	fmt.Println(frequencySort("tree"))
	fmt.Println(frequencySort("cccaaa"))
	fmt.Println(frequencySort("Aabb"))


}
