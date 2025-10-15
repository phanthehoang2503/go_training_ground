package dsa

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Insert(number int) {
	newNode := &Node{value: number}
	if l.head == nil {
		l.head = newNode
		return
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (l *LinkedList) Print() {
	current := l.head
	for current != nil {
		fmt.Printf("%d -> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

func New() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) AddRandomNumber(n int) {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	for i := 0; i < n; i++ {
		num := rng.Intn(100)
		l.Insert(num)
	}
}
