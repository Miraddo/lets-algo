package main

// AddTwoNumbers based on : https://leetcode.com/problems/add-two-numbers (Search : Linked List)
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sum int
	newNode := &ListNode{}
	result := newNode
	for l1 != nil || l2 != nil || sum > 0 {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		newNode.Next = &ListNode{Val: sum % 10}
		sum = sum / 10

		newNode = newNode.Next
	}

	return result.Next
}

// AddTwoNumbersWithoutLinkedList Add Two Numbers with arrays not using linked list
func AddTwoNumbersWithoutLinkedList(l1 []int, l2 []int) []int {

	// first i extend arrays to make both of them in same length
	if len(l1) > len(l2) {
		num := len(l1) - len(l2)

		for i := 0; i < num; i++ {
			l2 = append(l2, 0)
		}
	} else {
		num := len(l2) - len(l1)

		for i := 0; i < num; i++ {
			l1 = append(l1, 0)
		}
	}
	var sum, f, s int
	var result []int
	// sum numbers and append to result array
	for i := 0; i < len(l1); i++ {
		sum = l1[i] + l2[i]

		s = (sum + f) % 10
		f = (sum + f) / 10

		result = append(result, s)

	}
	// check if we still have number inside f so we should added it into result
	if f != 0 {
		result = append(result, f)
	}

	return result
}
