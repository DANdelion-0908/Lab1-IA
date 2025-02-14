package modules

import "slices"

type PriorityQueue struct {
	Items []Node
}

func EmptyPriority(queue PriorityQueue) bool {
	return len(queue.Items) == 0
}

func TopPriority(queue PriorityQueue) Node {
	return queue.Items[0]
}

func PopPriority(queue PriorityQueue) (PriorityQueue, Node) {
	removedItem := queue.Items[0]
	queue.Items = slices.Delete(queue.Items, 0, 1)
	return queue, removedItem
}

func AddPriority(queue PriorityQueue, item Node) PriorityQueue {
	queue.Items = append(queue.Items, item)
	if len(queue.Items) > 1 {
		queue = sortPriority(queue, len(queue.Items))
	}
	return queue
}

func sortPriority(queue PriorityQueue, n int) PriorityQueue { // Implementación de Bubble Sort
	if n == 1 {
		return queue
	}

	for i := 0; i < n-1; i++ {
		if queue.Items[i].RecoveryTime > queue.Items[i+1].RecoveryTime {
			queue.Items[i], queue.Items[i+1] = queue.Items[i+1], queue.Items[i]
		}
	}

	return sortPriority(queue, n-1)
}
