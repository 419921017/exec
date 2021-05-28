package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum1(nums, target))
}

func twoSum1(nums []int, target int) []int {
	for i, v := range nums {
		for k := i + 1; k < len(nums); k++ {
			if target-v == nums[k] {
				return []int{i, k}
			}
		}
	}
	return []int{}
}

// 空间换时间, 通过hashmap
func twoSum2(nums []int, target int) []int {
	result := []int{}
	m := make(map[int]int)
	for i, v := range nums {
		if value, exist := m[target-v]; exist {
			result = append(result, value)
			result = append(result, i)
		}
		m[v] = i
	}
	return result
}
