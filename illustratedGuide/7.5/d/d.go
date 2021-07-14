package D

func D() {

	var graph map[string]map[string]int

	graph["start"]["a"] = 6
	graph["start"]["b"] = 2

	graph["a"]["fin"] = 1

	graph["b"]["a"] = 3
	graph["b"]["fin"] = 5

	graph["fin"] = nil

	// fmt.Println(getMapKeys(graph))
	// node := findLowestCostNode(costs)
}

func findLowestCostNode() {

}

// func getMapKeys(target map[interface{}]interface{}) []interface{} {
// 	keys := make([]interface{}, len(target))
// 	for _, k := range target {
// 		keys = append(keys, k)
// 	}
// 	return keys
// }
