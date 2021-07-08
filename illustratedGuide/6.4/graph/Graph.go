package graph

import (
	"fmt"
)

func GraphMain() {
	graph := make(map[string][]string)
	// graph := make(map[string][]string)
	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}

	searchProcess(graph, "you")
}

func searchProcess(graph map[string][]string, name string) {
	searchQueue := []string{}
	searchQueue = append(searchQueue, graph[name]...)
	searched := make(map[string]bool)
	for len(searchQueue) > 0 {
		person := searchQueue[0]
		searchQueue = searchQueue[1:]
		if !searched[person] && personIsSeller(person) {
			fmt.Println(person + " is a mango seller")
			return
		} else {
			searchQueue = append(searchQueue, graph[person]...)
			searched[person] = true
		}
	}
}

func personIsSeller(name string) bool {
	// if len(name) == 0 {
	// 	return false
	// }
	return string(name[len(name)-1]) == "m"
}
