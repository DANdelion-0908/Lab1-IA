package modules

import "slices"

type Fifo struct {
	Items []string
}

func EmptyFifo(queue Fifo) bool {
	if len(queue.Items) == 0 {
		return true
	
	} else {
		return false
	}
}

func TopFifo(queue Fifo) string {
	return queue.Items[0]
}

func PopFifo(queue Fifo) (string, Fifo) {
	removedItem := queue.Items[0]
	queue.Items = slices.Delete(queue.Items, 0, 1)
	return removedItem, queue
}

func AddFifo(queue Fifo, item string) Fifo {
	queue.Items = append(queue.Items, item)
	return queue
}