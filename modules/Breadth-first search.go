package modules

func BFS(start Node, goal Node, graph map[string]Node) Fifo {
	visited := make(map[string]bool)
	queue := Fifo{}
	followedPath := Fifo{}

	for key := range graph {
		visited[key] = false
	}

	queue = AddFifo(queue, start.Activity)
	visited[start.Activity] = true

	for !EmptyFifo(queue) {
		currentNode, fixedQueue := PopFifo(queue)
		queue = fixedQueue

		if currentNode == goal.Activity {
			followedPath = AddFifo(followedPath, currentNode)
			return followedPath
		}

		for key := range graph[currentNode].Destiny {
			if !visited[key] {
				queue = AddFifo(queue, graph[key].Activity)
				visited[key] = true
			}
		}
		followedPath = AddFifo(followedPath, currentNode)
	}
	
	return followedPath
}
