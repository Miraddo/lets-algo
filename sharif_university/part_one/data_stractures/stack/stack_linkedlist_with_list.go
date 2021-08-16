package stack

import (
	"container/list"
	"fmt"
)

// SList is our struct that point to stack with container/list.List library
type SList struct {
	stack *list.List
}

// Push add a value into our stack
func (sl *SList) Push(val interface{})  {
	sl.stack.PushFront(val)
}

// Peak is return last value that insert into our stack
func (sl *SList) Peak() (interface{}, error) {
	if sl.stack.Len() > 0 {
		element := sl.stack.Front()
		return element.Value.(interface{}), nil
	}
	return "", fmt.Errorf("stack list is empty")
}

// Pop is return last value that insert into our stack
//also it will remove it in our stack
func (sl *SList) Pop() (interface{}, error) {
	if sl.stack.Len() > 0 {
		// get last element that insert into stack
		element := sl.stack.Front()
		// remove element in stack
		sl.stack.Remove(element)
		// return element value
		return element.Value.(interface{}), nil
	}
	return "", fmt.Errorf("stack list is empty")
}

// Length return length of our stack
func (sl *SList) Length() int {
	return sl.stack.Len()
}

// Empty check our stack has value or not
func (sl *SList) Empty() bool {
	// check our stack is empty or not
	// if is 0 it means our stack is empty otherwise is not empty
	return sl.stack.Len() == 0
}