package main

import (
	"linklist"
)

func main() {
	head := linklist.NewFromList([]int{1, 2, 3, 4, 5})
	head.PrintList()
	linklist.ReorderList(head)
}
