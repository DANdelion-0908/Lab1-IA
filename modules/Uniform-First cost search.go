package modules

func UCS(start Node, goal Node, graph map[string]Node) Fifo {
	queue := PriorityQueue{}
	cost := make(map[string]int)
	visited := make(map[string]bool)
	parents := make(map[string]string)
	followedPath := Fifo{}

	for key := range graph {
		cost[key] = 5 << 10 * 10
		visited[key] = false
	}

	queue = AddPriority(queue, start)
	cost[start.Activity] = 0

	for !EmptyPriority(queue) {

		var currentNode Node
		queue, currentNode = PopPriority(queue)

		if currentNode.Activity == goal.Activity {
			followedPath = reconstructPath(parents, goal.Activity)
			return followedPath
		}

		visited[currentNode.Activity] = true

		for key, edgeCost := range graph[currentNode.Activity].Destiny {
			newCost := cost[currentNode.Activity] + edgeCost

			if newCost < cost[key] {
				cost[key] = newCost
				parents[key] = currentNode.Activity
				queue = AddPriority(queue, graph[key])
			}
		}
	}

	return Fifo{}
}

func reconstructPath(parents map[string]string, goal string) Fifo {
	path := Fifo{}
	current := goal

	for current != "" {
		path = AddFifo(path, current)
		current = parents[current]
	}

	path = reverseFifo(path)

	return path
}

func reverseFifo(fifo Fifo) Fifo {
	reversed := Fifo{}
	for i := len(fifo.Items) - 1; i >= 0; i-- {
		reversed = AddFifo(reversed, fifo.Items[i])
	}
	return reversed
}
