package main

import (
	"fmt"
)

func main() {
	// phoneBook := make(map[string]int)
	// phoneBook["jenny"] = 8675309
	// phoneBook["emergency"] = 911

	// fmt.Println(phoneBook["jenny"])

	voter = make(map[string]bool)
	checkVoter("tom")
	checkVoter("mike")
	checkVoter("mike")
}

var voter map[string]bool

func checkVoter(name string) {
	_, ok := voter[name]
	// fmt.Println(v, ok)
	if ok {
		fmt.Println("kick them out")
	} else {
		voter[name] = true
		fmt.Println("let them voter")
	}
}
