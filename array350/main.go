package main

import (
	"fmt"
	"sort"
)

// 第350题：两个数组的交集
// 给定两个数组，编写一个函数来计算它们的交集。

func intersect1(nums1 []int, nums2 []int) []int {
	map0 := map[int]int{}
	for _, v := range nums1 {
		map0[v] += 1
	}
	k := 0

	for _, v := range nums2 {
		if map0[v] > 0 {
			map0[v] -= 1
			nums2[k] = v
			k++
		}

	}
	return nums2[0:k]
}

func intersect2(nums1 []int, nums2 []int) []int {
	x, y, z := 0, 0, 0
	// result := []int{}
	sort.Ints(nums1)
	sort.Ints(nums2)

	for x < len(nums1) && y < len(nums2) {
		if nums1[x] > nums2[y] {
			y++
		} else if nums1[x] < nums2[y] {
			x++
		} else {
			// result = append(result, nums1[x])
			nums1[z] = nums1[x]
			x++
			y++
			z++
		}
	}
	// return result[:]
	return nums1[:z]

}

func test() {
	fmt.Println(intersect1([]int{1, 2, 2, 3}, []int{2, 2}))
	fmt.Println(intersect2([]int{1, 2, 2, 3}, []int{2, 2}))

}

func main() {
	test()
}
