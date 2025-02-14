package main

import (
	"Lab1_IA/modules"
)

func createGraph() (map[string]modules.Node) {
	dataMap := modules.ReadFile("data/funcion_de_costo.csv")
	nodes := modules.CreateNode(dataMap)

	recoveryTimes := modules.ReadFile("data/heuristica.csv")
	nodes = modules.AddRecoveryTime(recoveryTimes, nodes)

	return nodes
}

func UseBFS(graph map[string]modules.Node) modules.Fifo {
	result := modules.BFS(graph["Warm-up activities"], graph["Stretching"], graph)
	return result
}

func UseDFS(graph map[string]modules.Node) modules.Fifo {
	result := modules.DFS(graph["Warm-up activities"], graph["Stretching"], graph)
	return result
}

func UseUCS(graph map[string]modules.Node) (modules.Fifo) {
	result := modules.UCS(graph["Warm-up activities"], graph["Stretching"], graph)
	return result
}

func UseGBFS(graph map[string]modules.Node) modules.Fifo{
	result := modules.GBFS(graph["Warm-up activities"], graph["Stretching"], graph)
	return result
}

func UseAStar(graph map[string]modules.Node) modules.Fifo{
	result := modules.AStar(graph["Warm-up activities"], graph["Stretching"], graph)
	return result
}