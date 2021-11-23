package main

// 找出input中两个数相加得target的值，返回坐标
// 不使用暴力计算，也就是O(N²)的
// 使用map来搞，O(N)
func twoSum(target int, input ...int) (x, y int) {
	// 忽略长度检查
	// 忽略边界检查
	stored := map[int]int{}
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
