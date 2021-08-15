package queue

// Order class
type Order struct {
	priority     int
	quantity     int
	product      string
	customerName string
}

// Queue Array of orders
type Queue []*Order

// New method for initialization
func (order *Order) New(priority int, quantity int, product string, customerName string) {
	order.priority = priority
	order.quantity = quantity
	order.product = product
	order.customerName = customerName
}

// Add order to queue
func (queue *Queue) Add(order *Order) {
	//Add directly to queue if empty
	if len(*queue) == 0 {
		*queue = append(*queue, order)
	} else {
		var appended bool
		appended = false
		var index int
		var addedOrder *Order
		for index, addedOrder = range *queue {
			if order.priority > addedOrder.priority {
				*queue = append(
					(*queue)[:index], append(
						Queue{order}, (*queue)[index:]...,
					)...,
				)
				appended = true
				break
			}
		}
		if !appended {
			*queue = append(*queue, order)
		}
	}
}
