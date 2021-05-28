package main

import "fmt"

// 给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次。
func main() {
	fmt.Println(removeEle([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

func removeEle(nums []int) int {
	for i := 0; i+1 < len(nums); {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return len(nums)
}
