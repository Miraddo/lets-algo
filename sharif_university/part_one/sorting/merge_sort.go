package Sorting

// MergeSort Merge Sort Algorithm function
func MergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	// recursive function
	first := MergeSort(nums[:len(nums)/2])
	second := MergeSort(nums[len(nums)/2:])

	// result with Merg Function
	result := Merge(first, second)

	return result
}

// Merge is for compare two list of arrays
func Merge(listOne, listTwo []int) []int {
	var i, j int
	var listThree []int

	ln1 := len(listOne)
	ln2 := len(listTwo)

	// if i smaller then length listOne and j is smaller then length listTwo
	for i < ln1 && j < ln2 {

		if listOne[i] < listTwo[j] {
			listThree = append(listThree, listOne[i])
			i++
		}else{
			listThree = append(listThree, listTwo[j])
			j++
		}
	}
	// at the end of sorting we looking into array if still we had numbers so we added to our listThree array
	for ; i < ln1; i++ {
		listThree = append(listThree, listOne[i])
	}

	for ; j < ln2; j++ {
		listThree = append(listThree, listTwo[j])
	}


	return listThree
}
