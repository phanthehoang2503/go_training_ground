package main

import "github.com/phanthehoang2503/go_training_ground/linkedList/dsa"

func main() {
	list := dsa.LinkedList{}
	list.Print()
	list.Insert(10)
	list.Insert(20)
	list.Insert(30)
	list.Print()
}
