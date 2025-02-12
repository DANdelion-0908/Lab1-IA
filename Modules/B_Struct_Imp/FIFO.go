package bstructimp

import "slices"

type Fifo struct {
	items []int
}

func EmptyFifo(queue Fifo) bool {
	if len(queue.items) == 0 {
		return true
	
	} else {
		return false
	}
}

func TopFifo(queue Fifo) int {
	return queue.items[0]
}

func PopFifo(queue Fifo) (int, Fifo) {
	removedItem := queue.items[0]
	queue.items = slices.Delete(queue.items, 0, 0)
	return removedItem, queue
}

func AddFifo(queue Fifo, item int) Fifo {
	queue.items = append(queue.items, item)
	return queue
}