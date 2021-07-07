package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(sum(arr))

	fmt.Println(MaxNumber(arr))
}

func sum(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	return sum([]int{arr[0]}) + sum(arr[1:])
}

func MaxNumber(arr []int) int {
	value := arr[0]
	for _, v := range arr {
		if value < v {
			value = v
		}
	}
	return value
}
