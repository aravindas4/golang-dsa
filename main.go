package main

import (
	"fmt"

	"github.com/aravindas4/golang-dsa/list"
)

func main() {

	// Linked list
	var linkedList list.LinkedList
	linkedList.AddToHead(1)
	linkedList.AddToHead(2)
	head := linkedList.HeadNode
	fmt.Println(head.Property)
	fmt.Println(linkedList.IterateList())
}
