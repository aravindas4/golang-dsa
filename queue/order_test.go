package queue

import "testing"

func TestQueue(t *testing.T) {
	queue := &Queue{}
	t.Run("Test Add method", func(t *testing.T) {
		order := &Order{}
		order.New(2, 10, "Python Bundle", "Aravinda")
		queue.Add(order)
		addedOrder := (*queue)[0]
		if addedOrder != order {
			t.Error("Test Empty Queue: inputted ", nil, " expected ", order, ", recieved ", addedOrder)
		}

		order = &Order{}
		order.New(1, 10, "Go Bundle", "Aravinda")
		queue.Add(order)
		addedOrder = (*queue)[1]
		if addedOrder != order {
			t.Error("Test Add Low priority: inputted ", nil, " expected ", order, ", recieved ", addedOrder)
		}

		order = &Order{}
		order.New(3, 10, "Django Bundle", "Aravinda")
		queue.Add(order)
		addedOrder = (*queue)[0]
		if addedOrder != order {
			t.Error("Test Add Low priority: inputted ", nil, " expected ", order, ", recieved ", addedOrder)
		}
	})
}
