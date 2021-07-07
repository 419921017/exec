package main

import "fmt"

func main() {
	arr := []int{5, 3, 6, 2, 10}
	fmt.Println(selectionSort(arr))
}

func selectionSort(arr []int) []int {
	var newArr []int
	for range arr {
		smallest := findSmallest(arr)
		newArr = append(newArr, arr[smallest])
		arr = append(arr[0:smallest], arr[smallest+1:]...)
	}
	return newArr
}

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallestIndex = i
		}
	}
	return smallestIndex
}
