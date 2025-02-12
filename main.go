package main

import (
	"Lab1_IA/modules"
	"fmt"
)

func main() {
	createGraph()
}

func createGraph() {
	dataMap := modules.ReadFile("data/funcion_de_costo.csv")
	nodes := modules.CreateNode(dataMap)
	
	recoveryTimes := modules.ReadFile("data/heuristica.csv")
	nodes = modules.AddRecoveryTime(recoveryTimes, nodes)
	
	i := 1
	for _, value := range nodes {
		j := 1
		fmt.Printf("%v. Actividad: %v\n", i, value.Activity)
		fmt.Printf("Tiempo de recuperación: %v\n", value.RecoveryTime)
		fmt.Printf("Destinos:\n\n")
		for key, value := range value.Destiny {
			fmt.Printf("\t%v. %v: %v\n\n", j, key, value)
			j++
		}
		i++
	}
}