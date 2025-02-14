package modules

import "slices"

type Lifo struct {
	Items []string
}

func EmptyLifo(stack Lifo) bool {
	if len(stack.Items) == 0 {
		return true

	} else {
		return false
	}
}

func TopLifo(stack Lifo) string {
	return stack.Items[len(stack.Items)-1]
}

func PopLifo(stack Lifo) (string, Lifo) {
	removedItem := stack.Items[len(stack.Items)-1]
	stack.Items = slices.Delete(stack.Items, len(stack.Items)-1, len(stack.Items))
	return removedItem, stack
}

func AddLifo(stack Lifo, item string) Lifo {
	stack.Items = append(stack.Items, item)
	return stack
}