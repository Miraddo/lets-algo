package main


type TreeNode struct {
	Val       int
	Left, Right *TreeNode
}

type BSTree struct {
	head   *TreeNode
	root   *TreeNode
	length int
}

// NewTreeNode create new node
func NewTreeNode(n int) *TreeNode {
	var newTreeNode TreeNode
	newTreeNode.Val = n
	newTreeNode.Left = nil
	newTreeNode.Right = nil
	return &newTreeNode
}

// Insert data to node
func (bt *BSTree) Insert(data int) {
	if bt.head == nil {
		bt.head = NewTreeNode(data)
		bt.root = bt.head
		bt.length++
		return
	}

	for bt.head != nil {

		if data >= bt.head.Val {
			if bt.head.Right == nil {
				bt.head.Right, bt.head = NewTreeNode(data), bt.root
				bt.length++
				return
			} else {
				bt.head = bt.head.Right
			}
		} else {
			if bt.head.Left == nil {
				bt.head.Left, bt.head = NewTreeNode(data), bt.root
				bt.length++

				return
			} else {
				bt.head = bt.head.Left
			}
		}
	}

}
