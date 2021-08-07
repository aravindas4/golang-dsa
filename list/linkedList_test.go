package list

import "testing"

func TestLinkedList(t *testing.T) {
	linkedList := &LinkedList{}
	if linkedList.HeadNode != nil {
		t.Error("Test Failed: {} inputted, {} expected, recieved: {}", nil, nil, linkedList.HeadNode)
	}

	linkedList.AddToHead(2)
	if linkedList.HeadNode.Property != 2 {
		t.Error("Test Failed: {} inputted, {} expected, recieved: {}", 2, 2, linkedList.HeadNode.Property)
	}

	linkedList.AddToHead(3)
	if linkedList.HeadNode.Property != 3 {
		t.Error("Test Failed: {} inputted, {} expected, recieved: {}", 3, 3, linkedList.HeadNode.Property)
	}
}
