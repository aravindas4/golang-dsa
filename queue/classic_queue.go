package queue

type ClassicQueue struct {
	queue []int
}

func (queue *ClassicQueue) ListItems() []int {
	return queue.queue
}

func (queue *ClassicQueue) Enqueue(item int) {
	queue.queue = append(queue.queue, item)
}

func (queue *ClassicQueue) Dequeue() int {
	item := queue.queue[0]
	queue.queue = queue.queue[1:]
	return item
}
