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
	ans := q.lst.Remove(q.lst.Front())
	return (ans).(T)
}
