package main

import "fmt"

var arr = []int{1, 0, 0, 2, 3, 0, 4, 5}

func moveZero(arr []int) {
	j := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[j] = arr[i]
			if i != j {
				arr[i] = 0
			}
			j++
		}
	}
}

func main() {
	moveZero(arr)
	fmt.Println(arr)

}
