package main

import "fmt"

//给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
//
// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
// 因为已经排好序，所以逐步比较临近的两个元素就好
// 比较两个临近位置就需要两个指针：a用于保留当前不重复的最后位置，b用于向后滑动确认与上一个不重复，不重复时a，b同时滑动，否则b继续滑，a保留原地，让b去找下一个不重复的替换它的位置
func removeDuplicates(input []int) (l int, output []int) {
	// 省略边界与长度检测
	n := len(input)
	if n == 0 {
		return 0, input
	}
	// 从第二个开始判断
	a := 1
	for b := 1; b < n; b++ {
		// 让b去找下一个不重复的替换a的位置
		if input[b] != input[b-1] {
			input[a] = input[b]
			fmt.Printf("%d,%d non-duplicated：%v \n", a, b, input)
			a++
		} else {
			fmt.Printf("%d,%d duplicated：%v \n", a, b, input)
		}
	}
	return a, input[:a]
}
