package queue

import "testing"

func TestClassicQueue(t *testing.T) {
	queue := &ClassicQueue{}

	t.Run("Test ListItems", func(t *testing.T) {
		received := queue.ListItems()
		if received != nil {
			t.Error("Test ListItems: inputted ", nil, " expected ", nil, ", recieved ", received)
		}

		newQueue := &ClassicQueue{}
		newQueue.Enqueue(4)
		newQueue.Enqueue(3)

		//
		expected := []int{4, 3}
		received = queue.ListItems()
		if received != nil {
			t.Error("Test ListItems: inputted ", nil, " expected ", expected, ", recieved ", received)
		}

		newQueue.Enqueue(5)
		newQueue.Dequeue()
		expected = []int{5, 4}
		received = queue.ListItems()
		if received != nil {
			t.Error("Test ListItems: inputted ", nil, " expected ", expected, ", recieved ", received)
		}
	})
}
