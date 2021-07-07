package quicksort

func Quicksort(arr []int) []int {
	if len(arr) < 2 {
		// 基线条件:
		// 为空或者只包含一个元素的数组是有序的
		return arr
	}
	pivot := arr[0]
	var less, greater []int
	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			less = append(less, arr[i])
		} else {
			greater = append(greater, arr[i])
		}
	}
	return append(append(Quicksort(less), pivot), Quicksort(greater)...)
}
