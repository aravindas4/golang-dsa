package stack

import (
	"reflect"
	"testing"
)

func TestStack(t *testing.T) {
	stack := &Stack{}
	stack.New()

	t.Run("Test initially", func(t *testing.T) {
		var expected, received []int
		received = stack.ListItems()
		if nil != received {
			t.Error("Test Failed: inputted ", nil, " expected ", expected, ", recieved ", received)
		}

		stack.Push(&Element{1})
		stack.Push(&Element{2})
		expected = []int{1, 2}
		received = stack.ListItems()
		if !reflect.DeepEqual(expected, received) {
			t.Error("Test Failed: inputted ", nil, " expected ", expected, ", recieved ", received)
		}

		popped := stack.Pop()
		received_element := &Element{2}
		if popped.GetValue() != received_element.GetValue() {
			t.Error("Test Failed: inputted ", nil, " expected ", received_element.GetValue(), ", recieved ", 1)
		}
	})
}
