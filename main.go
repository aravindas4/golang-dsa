package main

import (
	"fmt"

	"github.com/aravindas4/golang-dsa/list"
)

func main() {

	// Linked list
	var linkedList list.LinkedList
	linkedList.AddToHead(2)
	linkedList.AddToHead(3)
	linkedList.AddToHead(5)
	linkedList.AddAfter(2, 19)
	fmt.Println(linkedList.IterateList())
	var doubleLinkedList list.DoubleLinkedList
	doubleLinkedList.AddToHead(3)
	doubleLinkedList.AddToHead(2)
	doubleLinkedList.AddToHead(1)
	//node := doubleLinkedList.NodeBetweenValues(3, 1)
	fmt.Println(doubleLinkedList.IterateList())
}
