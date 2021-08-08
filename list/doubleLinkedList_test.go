package list

import (
	"reflect"
	"testing"
)

func TestDoubleLinkedList(t *testing.T) {
	doubleList := &DoubleLinkedList{}

	t.Run("Test AddToHead method", func(t *testing.T) {
		input1 := 1
		expected := 1
		doubleList.AddToHead(1)
		received := doubleList.HeadNode.Property
		if expected != received {
			t.Error("Test Failed: inputted ", input1, " expected ", nil, ", recieved ", received)
		}
	})

	t.Run("Test with node with values", func(t *testing.T) {
		input1 := 3
		input2 := 1
		inputs := []int{input1, input2}
		received := doubleList.NodeBetweenValues(inputs[0], inputs[1])
		if received != nil {
			t.Error("Test Failed: inputted ", inputs, " expected ", nil, ", recieved ", received)
		}

		doubleList.AddToHead(2)
		doubleList.AddToHead(3)
		expected := 2
		received2 := doubleList.NodeBetweenValues(inputs[0], inputs[1])
		if received2 == nil {
			t.Error("Test Failed: inputted ", inputs, " expected ", expected, ", recieved ", received2)
		}
	})

	t.Run("Test IterateList", func(t *testing.T) {
		expected := []int{3, 2, 1}
		received := doubleList.IterateList()
		if !reflect.DeepEqual(received, expected) {
			t.Error("Test Failed: inputted ", nil, " expected ", expected, ", recieved ", received)
		}

	})

	t.Run("Test NodeWithValue", func(t *testing.T) {
		input := 2
		expected := 2
		received := doubleList.NodeWithValue(input).Property
		if expected != received {
			t.Error("Test Failed: inputted ", input, " expected ", expected, ", recieved ", received)
		}

		newDoubleList := &DoubleLinkedList{}
		input = 2
		received2 := newDoubleList.NodeWithValue(input)
		if received2 != nil {
			t.Error("Test Failed: inputted ", input, " expected ", expected, ", recieved ", received2)
		}
	})

	t.Run("Test AddAfter", func(t *testing.T) {
		input := 10
		expected := []int{3, 2, 10, 1}
		doubleList.AddAfter(2, input)
		received := doubleList.IterateList()
		if !reflect.DeepEqual(expected, received) {
			t.Error("Test Failed: inputted ", input, " expected ", expected, ", recieved ", received)
		}
	})

	t.Run("Test LastNode", func(t *testing.T) {
		expected := 1
		received := doubleList.LastNode().Property
		if expected != received {
			t.Error("Test Failed: inputted ", nil, " expected ", expected, ", recieved ", received)
		}

		newDoubleList := &DoubleLinkedList{}
		received2 := newDoubleList.LastNode()
		if received2 != nil {
			t.Error("Test Failed: inputted ", nil, " expected ", nil, ", recieved ", received2)
		}
	})

	t.Run("Test AddToEnd", func(t *testing.T) {
		input := 11
		expected := []int{3, 2, 10, 1, 11}
		doubleList.AddToEnd(input)
		received := doubleList.IterateList()
		if !reflect.DeepEqual(expected, received) {
			t.Error("Test Failed: inputted ", input, " expected ", expected, ", recieved ", received)
		}
	})
}
