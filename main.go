package main

import (
	modules "Lab1_IA/modules"
	"fmt"
	"os"
)

func main() {
	nodes := createGraph()
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

	for {
		var path modules.Fifo
		UseUCS(nodes)
		fmt.Println("1. Breadth First Search")
		fmt.Println("2. Depth First Search")
		fmt.Println("3. Uniform Cost Search")
		fmt.Println("4. Greedy Best First Search")
		fmt.Println("5. A*")
		fmt.Println("6. Salir")

		var option int

		fmt.Println("\nIngresa el número del algoritmo a utilizar: ")
		fmt.Scan(&option)

		switch option {
		case 1:
			fmt.Println("\nBreadth First Search")
			path = UseBFS(nodes)

		case 2:
			fmt.Println("\nDepth First Search")
			path = UseDFS(nodes)

		case 3:
			fmt.Println("\nUniform Cost Search")
			path = UseUCS(nodes)

		case 4:
			fmt.Println("\nGreedy Best First Search")
			path = UseGBFS(nodes)

		case 5:
			fmt.Println("\nA*")
			path = UseAStar(nodes)

		case 6:
			fmt.Println("\nHasta la próxima")
			os.Exit(0)
		default:
			fmt.Println("\nOpción no válida")
		}

		for i := 0; i < len(path.Items); i++ {
			if i == len(path.Items)-1 {
				fmt.Printf("%v\n\n", path.Items[i])
			}
			fmt.Printf("%v -> ", path.Items[i])
		}
	}
}
