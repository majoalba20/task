package queue

import "errors"

var ErrQueueFull = errors.New("queue is full")

type TaskQueue interface {
	Enqueue(taskID uint) error
	Dequeue() <-chan uint
}

type taskQueue struct {
	ch chan uint
}

func NewTaskQueue(bufferSize int) TaskQueue {
	return &taskQueue{
		ch: make(chan uint, bufferSize),
	}
}

func (q *taskQueue) Enqueue(taskID uint) error {
	select {
	case q.ch <- taskID:
		return nil
	default:
		return ErrQueueFull
	}
}

func (q *taskQueue) Dequeue() <-chan uint {
	return q.ch
}
