package modules

func GBFS(start Node, goal Node, graph map[string]Node) Fifo {
	queue := PriorityQueue{}
	visited := make(map[string]bool)
	followedPath := Fifo{}

	for key := range graph {
		visited[key] = false
	}

	queue = AddPriority(queue, start)
	visited[start.Activity] = true
	followedPath = AddFifo(followedPath, start.Activity)

	for !EmptyPriority(queue) {
		fixedPQ, currentNode := PopPriority(queue)
		queue = fixedPQ

		if currentNode.Activity == goal.Activity {
			return followedPath
		}

		for key := range graph[currentNode.Activity].Destiny {
			if !visited[key] {
				queue = AddPriority(queue, graph[key])
				visited[key] = true
			}
		}
		followedPath = AddFifo(followedPath, TopPriority(queue).Activity)
	}

	return followedPath
}
