package main

import (
	"reflect"
	"testing"
)

type handler struct {
	firstList  []int
	secondList []int
	result     []int
}

//TestAddTwoNumbers Test Add Two Numbers
func TestAddTwoNumbers(t *testing.T) {
	firstNode, secondNode := LinkedList{}, LinkedList{}

	listData := []handler{
		{
			firstList:  []int{9, 9, 9, 9, 9, 9, 9},
			secondList: []int{9, 9, 9, 9},
			result:     []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
		{
			firstList:  []int{2, 4, 3},
			secondList: []int{5, 6, 4},
			result:     []int{7, 0, 8},
		},
	}

	for _, x := range listData {

		firstNode.Gen(x.firstList)
		secondNode.Gen(x.secondList)

		resultNode := AddTwoNumbers(firstNode.head, secondNode.head)

		current := resultNode

		result := []int{}

		for resultNode != nil {
			if current != nil {
				result = append(result, current.Val)
				current = current.Next
			} else {
				break
			}
		}

		if !reflect.DeepEqual(x.result, result) {
			t.Errorf("AddTwoNumbers was incorrect, got: %v, want: %v.", result, x.result)
		}

		firstNode.Delete()
		secondNode.Delete()

	}

}

type List struct {
	firstArr  []int
	secondArr []int
	result    []int
}

// TestAddTwoNumbersWithoutLinkedList  Test Add Two Numbers Without Linked List
func TestAddTwoNumbersWithoutLinkedList(t *testing.T) {
	handle := []List{
		{
			firstArr:  []int{9, 9, 9, 9, 9, 9, 9},
			secondArr: []int{9, 9, 9, 9},
			result:    []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
		{
			firstArr:  []int{2, 4, 3},
			secondArr: []int{5, 6, 4},
			result:    []int{7, 0, 8},
		},
	}

	for _, x := range handle {

		res := AddTwoNumbersWithoutLinkedList(x.firstArr, x.secondArr)

		if !reflect.DeepEqual(res, x.result) {

			t.Errorf("AddTwoNumbersWithoutLinkedList was incorrect, got: %v, want: %v.", res, x.result)
		}

	}
}
