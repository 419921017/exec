package main

import "fmt"

func main() {
	// 	给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
	//  最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
	//  你可以假设除了整数 0 之外，这个整数不会以零开头。
	fmt.Println(plusOne([]int{4, 3, 2, 9}))
}

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return digits
	}
	var result []int
	addon := 0
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += addon
		addon = 0
		if i == len(digits)-1 {
			digits[i]++
		}
		if digits[i] == 10 {
			addon = 1
			digits[i] = digits[i] % 10
		}
	}
	if addon == 1 {
		result = make([]int, 1)
		result[0] = 1
		result = append(result, digits...)
	} else {
		result = digits
	}
	return result
}
