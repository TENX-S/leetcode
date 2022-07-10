package CPC

import (
	"errors"
)

type Queue[T any] struct {
	inner           []T
	head, tail, cap int
}

func NewQueue[T any](cap int) Queue[T] {
	return Queue[T]{
		inner: make([]T, cap),
		cap:   cap,
	}
}

func (q *Queue[T]) IsEmpty() bool {
	return q.head == q.tail
}

func (q *Queue[T]) IsFull() bool {
	return q.head == (q.tail+1)%q.cap
}

func (q *Queue[T]) Enqueue(x T) {
	if q.IsFull() {
		var zero T
		q.inner = append(q.inner, zero)
		q.cap = cap(q.inner)
		q.inner[q.tail] = x
		q.tail++
	}
	q.inner[q.tail] = x
	if q.tail+1 == q.cap {
		q.tail = 0
	} else {
		q.tail++
	}
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("queue is empty")
	}
	res := q.inner[q.head]
	if q.head+1 == q.cap {
		q.head = 0
	} else {
		q.head++
	}
	return res, nil
}

type Task struct {
	name string
	need int
}

type Record struct {
	name     string
	finished int
}

func HandleTasks(q int, tasks []Task) (records []Record) {
	queue := NewQueue[Task](len(tasks) + 1)
	for _, task := range tasks {
		queue.Enqueue(task)
	}
	var time int
	for !queue.IsEmpty() {
		t, err := queue.Dequeue()
		if err != nil {
			return nil
		}
		var passed int
		if t.need > q {
			t.need -= q
			queue.Enqueue(t)
			passed = q
		} else {
			passed = t.need
			records = append(records, Record{
				name:     t.name,
				finished: time + t.need,
			})
		}
		time += passed
	}
	return
}
