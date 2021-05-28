// 给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。

package main

import "fmt"

func main() {
	fmt.Println(removeEle([]int{1, 2, 3, 4, 2, 4}, 2))
}

func removeEle(nums []int, val int) []int {
	for i := 0; i < len(nums); {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return nums
}
