package stack

var stackArray []int

// stack_push push to first index of array
func stack_push(n int) {
	stackArray = append([]int{n}, stackArray...)
}

// stack_length return length of array
func stack_length() int  {
	return len(stackArray)
}

// stack_peak return last input of array
func stack_peak() int {
	return stackArray[0]
}

// stack_empty check array is empty or not
func stack_empty() bool  {
	return len(stackArray) == 0
}

// stack_pop return last input and remove it in array
func stack_pop() int {
	pop := stackArray[0]
	stackArray = stackArray[1:]
	return pop
}