package queue

import "errors"

type node struct {
	data interface{}
	next *node
}
type Queue struct {
	head *node
	tail *node
}

func NewQueue() Queue {
	return Queue{}
}

func (q *Queue) Push(value interface{}) {
	node := &node{data: value, next: nil}
	if q.head == nil {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = q.tail.next
	}
}

func (q *Queue) Pop() error {
	if q.head == nil {
		return errors.New("queue is empty")
	}
	q.head = q.head.next
	return nil
}

func (q *Queue) Top() (interface{}, error) {
	if q.head == nil {
		return 0, errors.New("queue is empty")
	}
	value := q.head.data
	return value, nil
}

func (q *Queue) IsEmpty() bool {
	if q.head == nil {
		return true
	}
	return false
}
