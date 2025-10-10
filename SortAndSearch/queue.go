package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	data []int
}

func (q *Queue) Enqueue(x int) {
	q.data = append(q.data, x)
	fmt.Printf("Enqueue %d, Queue: %v \n", x, q.data)
}

func (q *Queue) Dequeue() (int, error) {
	if q.isEmpty() {
		return 0, errors.New("Queue is empty")
	}
	val := q.data[0]
	q.data = q.data[1:]
	fmt.Printf("Dequeue %d, Queue now: %v \n", val, q.data)
	return val, nil
}

func (q *Queue) Front() {
	fmt.Print("Queue front: ", q.data[0])
}

func (q *Queue) Back() {
	fmt.Printf("Queue back: %d \n", q.data[q.size()-1])
}

func (q *Queue) Clear() {
	q.data = []int{}
	fmt.Printf("Queue cleared \n")
}

func (q *Queue) Print() {
	fmt.Printf("Queue: %d \n", q.data)
}

func (q *Queue) isEmpty() bool {
	return q.size() == 0
}

func (q *Queue) size() int {
	return len(q.data)
}
