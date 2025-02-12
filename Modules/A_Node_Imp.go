package modules

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Activity     string
	RecoveryTime int
	Destiny      map[string]int
}

func ReadFile(fileDirection string) map[string][][2]string {
	activities := make(map[string][][2]string)
	file, err := os.Open(fileDirection)

	if err != nil {
		log.Fatal(err)
	}

	csvReader := csv.NewReader(file)
	csvReader.Comma = ';'

	for {
		data, err := csvReader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		if len(data) == 3 {
			data[0] = strings.TrimPrefix(data[0], "\ufeff")
			activities[data[0]] = append(activities[data[0]], [2]string{data[1], data[2]})

		} else if len(data) == 2 {
			data[0] = strings.TrimPrefix(data[0], "\ufeff")
			activities[data[0]] = append(activities[data[0]], [2]string{data[1]})

		} else {
			log.Fatal("Error: The file has an invalid format")
		}
	}

	defer file.Close()
	return activities
}

func CreateNode(activities map[string][][2]string) []Node {
	createdNodes := make([]Node, 0)
	for key, value := range activities {
		node := Node{}
		node.Destiny = make(map[string]int)
		node.Activity = key
		for _, list := range value {
			for elementIndex, element := range list {
				if elementIndex%2 == 0 {
					node.Destiny[element], _ = strconv.Atoi(list[1])
				}
			}
		}
		createdNodes = append(createdNodes, node)
	}

	createdNodes = append(createdNodes, CreateStretchingNode())

	return createdNodes
}

func CreateStretchingNode() (node Node) {
	node.Activity = "Stretching"
	node.RecoveryTime = 0
	node.Destiny = make(map[string]int)

	return
}

func AddRecoveryTime(recoveryTimes map[string][][2]string, nodes []Node) []Node {
	fullNodes := make([]Node, 0)
	for key, value := range recoveryTimes {
		for _, node := range nodes {
			if key == node.Activity {
				node.RecoveryTime, _ = strconv.Atoi(value[0][0])
				fullNodes = append(fullNodes, node)
			}
		}
	}

	return fullNodes
}