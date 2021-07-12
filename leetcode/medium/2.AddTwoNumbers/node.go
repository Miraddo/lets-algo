package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	length int
	head   *ListNode
}

func (ll *LinkedList) Add(n *ListNode) {
	var snd *ListNode
	snd, ll.head = ll.head, n
	ll.head.Next = snd
	ll.length++
}

func (ll *LinkedList) Gen(nums []int) {
	for _, x := range nums {
		ll.Add(&ListNode{Val: x})
	}
}

func (ll *LinkedList) Length() int {
	return ll.length
}

func (ll *LinkedList) Reverse() {
	current := ll.head
	var prevent *ListNode
	for current != nil {
		tmp := current.Next
		current.Next = prevent
		prevent = current
		current = tmp
	}
	ll.head = prevent
}

func (ll *LinkedList) Show() {

	current := ll.head
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}

func (ll *LinkedList) Delete() {

	ll.head = nil
	ll.length = 0

}
