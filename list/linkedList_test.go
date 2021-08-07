package list

import (
	"reflect"
	"testing"
)

func TestLinkedList(t *testing.T) {
	linkedList := &LinkedList{}
	if linkedList.HeadNode != nil {
		t.Error("Test Failed: {} inputted, {} expected, recieved: {}", nil, nil, linkedList.HeadNode)
	}

	// Add To Head
	linkedList.AddToHead(2)
	if linkedList.HeadNode.Property != 2 {
		t.Error("Test Failed: {} inputted, {} expected, recieved: {}", 2, 2, linkedList.HeadNode.Property)
	}

	linkedList.AddToHead(3)
	if linkedList.HeadNode.Property != 3 {
		t.Error("Test Failed: {} inputted, {} expected, recieved: {}", 3, 3, linkedList.HeadNode.Property)
	}

	// Iterate list
	t.Run("Test iterate function", func(t *testing.T) {
		expected := []int{3, 2}
		received := linkedList.IterateList()
		if !reflect.DeepEqual(received, expected) {
			t.Error("Test Failed: inputted {}, expected {}, recieved {}", nil, expected, received)
		}
	})
}
