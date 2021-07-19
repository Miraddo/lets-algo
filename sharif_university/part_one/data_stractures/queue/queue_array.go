package queue

var ListQueue []int

func EnQueue(n int)  {
	ListQueue = append(ListQueue, n)
}

func DeQueue() int {
	data := ListQueue[0]
	ListQueue = ListQueue[1:]
	return data
}

func FrontQueue() int {
	return  ListQueue[0]
}

func BackQueue() int {
	return ListQueue[len(ListQueue)-1]
}

func LenQueue() int {
	return len(ListQueue)
}

func isEmptyQueue() bool {
	return len(ListQueue) == 0
}