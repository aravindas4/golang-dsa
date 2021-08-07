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
