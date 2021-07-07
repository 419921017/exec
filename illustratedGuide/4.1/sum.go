package main

import (
	"fmt"
	"xiaohao/illustratedGuide/4.1/number"
	"xiaohao/illustratedGuide/4.1/quicksort"
)

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(sum(arr))

	fmt.Println(number.MaxNumber(arr))

	arr2 := []int{10, 5, 2, 3, 8}
	fmt.Println(quicksort.Quicksort(arr2))
}

func sum(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	return sum([]int{arr[0]}) + sum(arr[1:])
}
