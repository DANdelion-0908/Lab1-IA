package modules

func AStar(start Node, goal Node, graph map[string]Node) Fifo {
	queue := PriorityQueue{}
	visited := make(map[string]bool)
	gScore := make(map[string]int)
	heuristic := make(map[string]int)
	followedPath := Fifo{}

	for key, value := range graph {
		gScore[key] = 1 << 10 * 10
		heuristic[key] = value.RecoveryTime
	}
	gScore[start.Activity] = 0

	queue = AddPriority(queue, start)

	for !EmptyPriority(queue) {
		queue, currentNode := PopPriority(queue)

		followedPath = AddFifo(followedPath, currentNode.Activity)

		if currentNode.Activity == goal.Activity {
			return followedPath
		}

		for neighbor, cost := range graph[currentNode.Activity].Destiny {
			tentativeG := gScore[currentNode.Activity] + cost

			if tentativeG < gScore[neighbor] {
				gScore[neighbor] = tentativeG
				tempNode := graph[neighbor]

				tempNode.RecoveryTime = tentativeG + heuristic[neighbor] // f(n) = g(n) + h(n)
				graph[neighbor] = tempNode
				queue = AddPriority(queue, graph[neighbor])
				visited[neighbor] = true
			}
		}
	}
	return followedPath
}
