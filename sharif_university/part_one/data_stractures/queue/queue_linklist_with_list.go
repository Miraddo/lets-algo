package queue

import (
	"container/list"
	"fmt"
)

type LQueue struct {
	queue *list.List
}

func (lq *LQueue) Enqueue(value interface{}) {
	lq.queue.PushBack(value)
}

func (lq *LQueue) Dequeue() error {

	if lq.queue.Len() > 0 {
		element := lq.queue.Front()
		lq.queue.Remove(element)

		return nil
	}

	return fmt.Errorf("pop error")
}

func (lq *LQueue) Front() (interface{}, error)  {
	if lq.queue.Len() > 0 {
		if val, ok := lq.queue.Front().Value.(interface{}); ok{
			return val, nil
		}
	}

	return "", fmt.Errorf("error queue is empty")
}

func (lq *LQueue) Back() (interface{}, error)  {
	if lq.queue.Len() > 0 {
		if val, ok := lq.queue.Back().Value.(interface{}); ok{
			return val, nil
		}
	}

	return "", fmt.Errorf("error queue is empty")
}

func (lq *LQueue) Len() int {
	return lq.queue.Len()
}

func (lq *LQueue) Empty() bool {
	return lq.queue.Len() == 0
}

