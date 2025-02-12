package bstructimp

import "slices"

type Lifo struct {
	items []int
}

func EmptyLifo(stack Lifo) bool {
	if len(stack.items) == 0 {
		return true
	
	} else {
		return false
	}
}

func TopLifo(stack Lifo) int {
	return stack.items[len(stack.items)-1]
}

func PopLifo(stack Lifo) (int, Lifo) {
	removedItem := stack.items[len(stack.items)-1]
	stack.items = slices.Delete(stack.items, len(stack.items)-1, len(stack.items)-1)
	return removedItem, stack
}

func AddLifo(stack Lifo, item int) Lifo {
	stack.items = append(stack.items, item)
	return stack
}