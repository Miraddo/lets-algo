// Binary Search Tree

package main

import "fmt"

type Node struct {
	value       int
	left, right *Node
}

type BSTree struct {
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
func (bt *BSTree) Insert(data int) {
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
func (bt *BSTree) Lookup(value int) bool {
	if bt.root == nil {
		return false
	}

	for bt.root != nil {

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

func InOrder(nd *Node) {
	if nd == nil{
		return
	}

	InOrder(nd.left)
	fmt.Println(nd.value)
	InOrder(nd.right)

}

func PreOrder(nd *Node) {
	if nd == nil{
		return
	}

	fmt.Println(nd.value)
	InOrder(nd.left)
	InOrder(nd.right)

}

func PostOrder(nd *Node) {
	if nd == nil{
		return
	}
	InOrder(nd.left)
	InOrder(nd.right)
	fmt.Println(nd.value)

}

func DeleteNode(nd *Node, value int)  {

}


func main() {
	var bt BSTree
	bt.Insert(50) // first value is our root
	bt.Insert(30)
	bt.Insert(20)
	bt.Insert(40)
	bt.Insert(70)
	bt.Insert(60)
	bt.Insert(80)


	//fmt.Println(bt.length)
	//fmt.Println(bt.Lookup(700))

	InOrder(bt.head)
	//fmt.Println("=========================")
	//PreOrder(bt.head)
	//fmt.Println("=========================")
	//PostOrder(bt.head)
	//fmt.Println("=========================")


}
