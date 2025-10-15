package main

import "github.com/phanthehoang2503/go_training_ground/linkedList/dsa"

func main() {
	list := dsa.New()
	list.Print()
	list.AddRandomNumber(10)
	list.Print()
}
