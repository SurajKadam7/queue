package queue

import (
	"container/list"
)

type queue[T any] struct {
	lst *list.List
}

func New[T any]() queue[T] {
	return queue[T]{lst: list.New()}
}

func (q *queue[T]) Push(value T) {
	q.lst.PushBack(value)
}

func (q *queue[T]) Pop() T {
	var value T
	front := q.lst.Front()
	if front == nil {
		return value
	}
	ans := q.lst.Remove(front)
	return (ans).(T)
}

func (q *queue[T]) Len() int {
	return q.lst.Len()
}
