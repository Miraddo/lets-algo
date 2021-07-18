package stack

// Node structure
type Node struct {
	Val  int
	Next *Node
}

// Stack has jost top of node and with length
type Stack struct {
	top    *Node
	length int
}

// push add value to last index
func (ll *Stack) push(n int) {
	newStack := &Node{} // new node

	newStack.Val = n
	newStack.Next = ll.top

	ll.top = newStack
	ll.length++
}

// pop remove last item as first output
func (ll *Stack) pop() int {
	result := ll.top.Val
	if ll.top.Next == nil {
		ll.top = nil
	} else {
		ll.top.Val, ll.top.Next = ll.top.Next.Val, ll.top.Next.Next
	}

	ll.length--
	return result
}

// isEmpty to check our array is empty or not
func (ll *Stack) isEmpty() bool {
	return ll.length == 0
}


// len use to return length of our stack
func (ll *Stack) len() int {
	return ll.length
}

// peak return last input value
func (ll *Stack) peak() int {
	return ll.top.Val
}

// show all value as an interface array
func (ll *Stack) show() (in []int) {
	current := ll.top

	for current != nil {
		in = append(in, current.Val)
		current = current.Next
	}
	return
}
