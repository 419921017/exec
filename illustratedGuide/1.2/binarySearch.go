package main

import (
	"fmt"
)

func main() {
	list := []int{1, 3, 4, 5, 7, 9}

	fmt.Println(binarySearch(list, 3))
}

func binarySearch(list []int, item int) int {
	low := 0
	high := len(list) - 1
	for low <= high {
		mid := (low + high) / 2
		guess := list[mid]

		if guess == item {
			return mid
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}

	}
	return -1
}
