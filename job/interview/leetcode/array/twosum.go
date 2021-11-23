package main

import "sort"

// 找出input中两个数相加得target的值，返回坐标
// 不使用暴力计算，也就是O(N²)的
// 使用map来搞，O(N) Time，O(N) Space
// 省时，但是费空间
func twoSumWithMap(target int, input ...int) (x, y int) {
	// 忽略长度检查
	// 忽略边界检查
	// O(N) Space
	stored := map[int]int{}
	// O(N) time
	for i, num := range input {
		// 记录当前值的匹配值，然后去map里找，如果没有就把自己塞进去，让别人能找到
		tempNum := target - num
		if j, cached := stored[tempNum]; cached {
			// 返回已缓存的坐标与自己的坐标
			return j, i
		}

		stored[num] = i
	}
	return
}

// 使用排序来搞，O(NlogN) Time，O(1) Space
// 费时，但是省空间
func twoSumWithSorting(target int, input ...int) (x, y int) {
	// 忽略长度检查
	// 忽略边界检查

	// 排序，然后两端夹逼
	// 夹逼值大于目标值则end向前进一位
	// 小于则start后进一位
	sort.Ints(input)
	start, end := 0, len(input)-1
	for start < end {
		tempSum := input[start] + input[end]
		if tempSum == target {
			return start, end
		} else if tempSum < target {
			start++
		} else {
			end--
		}
	}

	return
}
