package queue

import (
	goqueue "github.com/yireyun/go-queue"
)

type Queue interface {
	Put(any)
	Get() (any, bool)
	Cap() int
	Len() int
}

type queue struct {
	queue *goqueue.EsQueue

	// queue *internal.Queue[any]
	// mu sync.Mutex
}

var _ Queue = (*queue)(nil)

func NewQueue() Queue {
	return &queue{
		queue: goqueue.NewQueue(64),
		// queue: internal.New[any](64),
	}
}

// Get implements Queue.
func (q *queue) Get() (any, bool) {
	v, ok, _ := q.queue.Get()
	return v, ok
}

// Put implements Queue.
func (q *queue) Put(val any) {
	q.queue.Put(val)
}

// Cap implements Queue.
func (q *queue) Cap() int {
	return int(q.queue.Capaciity())
}

// Len implements Queue.
func (q *queue) Len() int {
	return int(q.queue.Quantity())
}
