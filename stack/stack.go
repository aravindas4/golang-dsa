package stack

import "strconv"

//Element class
type Element struct {
	elementValue int
}

func (element *Element) String() string {
	return strconv.Itoa(element.elementValue)
}

func (element *Element) GetValue() int {
	return element.elementValue
}

type Stack struct {
	elements []*Element
	// elementCount int
}

func (stack *Stack) New() {
	stack.elements = make([]*Element, 0)
}

func (stack *Stack) ListItems() []int {
	var items []int
	for _, item := range stack.elements {
		items = append(items, item.elementValue)
	}
	return items
}

//Push to end of stack
func (stack *Stack) Push(element *Element) {
	stack.elements = append(stack.elements, element)
	// stack.elementCount = len(stack.elements)
}

//Pop from end of stack
func (stack *Stack) Pop() *Element {
	length := len(stack.elements)
	if length == 0 {
		return nil
	}

	var element *Element
	element, stack.elements = stack.elements[length-1], stack.elements[:length-1]
	return element
}
