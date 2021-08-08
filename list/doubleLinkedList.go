package list

import "fmt"

//Node class
type Node2 struct {
	Property     int
	nextNode     *Node2
	previousNode *Node2
}

//Double Linked List
type DoubleLinkedList struct {
	HeadNode *Node2
}

func (list *DoubleLinkedList) AddToHead(property int) {
	node := &Node2{
		Property: property,
	}

	headNode := list.HeadNode
	if headNode != nil {
		node.nextNode = headNode
		headNode.previousNode = node
	}
	list.HeadNode = node
	fmt.Println(node.nextNode)
}

func (list *DoubleLinkedList) IterateList() []int {
	var node *Node2
	var items []int
	for node = list.HeadNode; node != nil; node = node.nextNode {
		items = append(items, node.Property)
	}
	return items
}

func (list *DoubleLinkedList) NodeBetweenValues(firstProperty int, secondProperty int) *Node2 {
	var node, nodeWith *Node2

	//Iterate
	for node = list.HeadNode; node != nil; node = node.nextNode {
		if node.nextNode != nil && node.previousNode != nil {
			if node.previousNode.Property == firstProperty && node.nextNode.Property == secondProperty {
				nodeWith = node
				break
			}
		}
	}
	return nodeWith
}

func (list *DoubleLinkedList) NodeWithValue(property int) *Node2 {
	var node, nodeWith *Node2

	//Iterate
	for node = list.HeadNode; node != nil; node = node.nextNode {
		if node.Property == property {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

func (list *DoubleLinkedList) AddAfter(nodeProperty int, property int) {
	var node *Node2
	newNode := &Node2{
		Property: property,
	}

	//Get the node
	node = list.NodeWithValue(nodeProperty)

	if node != nil {
		previousNode := node.nextNode
		node.nextNode = newNode
		newNode.nextNode = previousNode
		newNode.previousNode = node
	}
}

func (list *DoubleLinkedList) LastNode() *Node2 {
	var node, lastNode *Node2
	for node = list.HeadNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
			break
		}
	}
	return lastNode
}

func (list *DoubleLinkedList) AddToEnd(property int) {
	var node *Node2
	newNode := &Node2{
		Property: property,
	}
	node = list.LastNode()

	if node != nil {
		node.nextNode = newNode
		newNode.previousNode = node
	}
}
