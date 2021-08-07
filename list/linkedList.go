package list

// Node class
type Node struct {
	//Exported
	Property int
	//Non exported
	nextNode *Node
}

type LinkedList struct {
	//Exported
	HeadNode *Node
}

func (list *LinkedList) AddToHead(data int) {
	// Reference to node
	node := &Node{}
	node.Property = data

	headNode := list.HeadNode

	// If empty list
	if headNode != nil {
		node.nextNode = headNode
	}
	list.HeadNode = node
}

func (list *LinkedList) IterateList() []int {
	var node *Node
	var items []int
	for node = list.HeadNode; node != nil; node = node.nextNode {
		items = append(items, node.Property)
	}
	return items
}
