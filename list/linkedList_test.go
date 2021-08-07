package list

import "testing"

func TestLinkedList(t *testing.T) {
	linkedList := &LinkedList{}
	if linkedList.HeadNode != nil {
		t.Error("Expected nil")
	}

	linkedList.AddToHead(2)
	if linkedList.HeadNode.Property != 2 {
		t.Error("Expected 2")
	}

	linkedList.AddToHead(3)
	if linkedList.HeadNode.Property != 3 {
		t.Error("Expected 3")
	}
}
