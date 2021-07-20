package main

import "fmt"

type Node struct {
	value       int
	left, right *Node
}

type BTree struct {
	head	*Node
	root   *Node
	length int
}

// NewNode create new node
func NewNode(n int) *Node {
	var newNode Node
	newNode.value = n
	newNode.left = nil
	newNode.right = nil
	return &newNode
}

// Insert data to node
func (bt *BTree) Insert(data int) {
	if bt.head == nil {
		bt.head = NewNode(data)
		bt.root = bt.head
		bt.length++
		return
	}

	for bt.head != nil {

		if data >= bt.head.value{
			if bt.head.right == nil {
				bt.head.right, bt.head = NewNode(data), bt.root
				bt.length++
				return
			} else {
				bt.head = bt.head.right
			}
		}else{
			if bt.head.left == nil {
				bt.head.left, bt.head = NewNode(data), bt.root
				bt.length++

				return
			} else {
				bt.head = bt.head.left
			}
		}
	}

}

// Lookup search in in bts
func (bt *BTree) Lookup(value int) bool {
	if bt.root == nil {
		return false
	}

	for bt.root != nil {
		//fmt.Println(bt.root)
		if value == bt.root.value {
			return true
		}

		if value > bt.root.value{
			bt.root = bt.root.right
		}else{
			bt.root = bt.root.left
		}


	}


	return false
}

func main() {
	var bt BTree
	bt.Insert(5) // first value is our root
	bt.Insert(3)
	bt.Insert(2)
	bt.Insert(6)
	bt.Insert(7)


	fmt.Println(bt.length)
	fmt.Println(bt.Lookup(6))
}
