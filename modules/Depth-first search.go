package modules

func DFS(start Node, goal Node, graph map[string]Node) Fifo {
	followedPath := Fifo{}
	stack := Lifo{}
	visited := make(map[string]bool)

	for key := range graph {
		visited[key] = false
	}

	stack = AddLifo(stack, start.Activity)
	visited[start.Activity] = true

	for !EmptyLifo(stack) {
		currentNode, fixedStack := PopLifo(stack)
		stack = fixedStack

		if currentNode == goal.Activity {
			visited[currentNode] = true
			followedPath = AddFifo(followedPath, currentNode)
			return followedPath
		}

		for key := range graph[currentNode].Destiny {
			if !visited[key] {
				stack = AddLifo(stack, key)
			}
		}
		followedPath = AddFifo(followedPath, currentNode)
	}
	return followedPath
}
