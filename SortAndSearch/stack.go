package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	data []int
}

func (s *Stack) Push(x int) {
	s.data = append(s.data, x)
	fmt.Printf("Push %d, Stack: %v \n", x, s.data)
}

func (s *Stack) Pop() (int, error) {
	if s.isEmpty() {
		return 0, errors.New("stack is empty")
	}

	topIdx := s.size() - 1
	topValue := s.getHead()

	s.data = s.data[:topIdx]
	fmt.Printf("Pop %d, Stack now: %v \n", topValue, s.data)
	return topValue, nil
}

func (s *Stack) Peek() (int, bool) {
	if s.isEmpty() {
		return 0, false
	}
	return s.getHead(), true
}

func (s *Stack) Clear() {
	s.data = []int{}
	fmt.Println("Stack cleared")
}

func (s *Stack) Print() {
	fmt.Println("Stack:", s.data)
}

func (s *Stack) isEmpty() bool {
	return s.size() == 0
}

func (s *Stack) getHead() int {
	return s.data[s.size()-1]
}

func (s *Stack) size() int {
	return len(s.data)
}
