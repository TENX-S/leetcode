package CPC

import (
	"fmt"
	"testing"
)

func TestHandleTasks(t *testing.T) {
	tasks := []Task{
		{
			name: "p1",
			need: 150,
		},
		{
			name: "p2",
			need: 80,
		},
		{
			name: "p3",
			need: 200,
		},
		{
			name: "p4",
			need: 350,
		},
		{
			name: "p5",
			need: 20,
		},
	}
	t.Log(HandleTasks(100, tasks))
}

func TestQueue(t *testing.T) {
	q := NewQueue[int](2)
	for i := 1; i <= 9; i++ {
		q.Enqueue(i)
	}
	for i := 1; i <= 8; i++ {
		_, _ = q.Dequeue()
	}
	fmt.Println()
}
