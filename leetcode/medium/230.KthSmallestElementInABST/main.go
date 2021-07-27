// based on : https://leetcode.com/problems/kth-smallest-element-in-a-bst/submissions/
// 230. Kth Smallest Element in a BST

package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var arr []int

func kthSmallest(root *TreeNode, k int) int {
	l := ln(root)
	arr = []int{}
	return l[k-1]
}


func ln (root *TreeNode) []int{
	if root == nil{
		return arr
	}

	ln(root.Left)

	arr = append(arr, root.Val)

	ln(root.Right)

	return arr
}

func main()  {
	//var bt BSTree
	//bt.Insert(4) // first value is our root
	//bt.Insert(3)
	//bt.Insert(2)
	//bt.Insert(6)
	//bt.Insert(8)
	//bt.Insert(9)
	//
	//
	//fmt.Println(kthSmallest(bt.root, 1))
}