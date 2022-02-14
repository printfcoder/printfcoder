package main

//给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
//
// 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
//
// 你可以假设除了整数 0 之外，这个整数不会以零开头。
//
//
//
// 示例 1：
//
//
//输入：digits = [1,2,3]
//输出：[1,2,4]
//解释：输入数组表示数字 123。
//
//
// 示例 2：
//
//
//输入：digits = [4,3,2,1]
//输出：[4,3,2,2]
//解释：输入数组表示数字 4321。
//
//
// 示例 3：
//
//
//输入：digits = [0]
//输出：[1]
//
//
//
//
// 提示：
//
//
// 1 <= digits.length <= 100
// 0 <= digits[i] <= 9

// 从尾至头模拟十进制加法即可，注意可能要进位
func plusOne(digits []int) []int {
	// 加数
	addend := 1

	for last := len(digits) - 1; last >= 0; last-- {
		// 需要进位
		if digits[last]+addend > 9 {
			digits[last] = (digits[last] + addend) % 10
			addend = 1
		} else if addend == 1 {
			// 小于10的直接加上即可
			digits[last] += 1
			addend = 0
		}
	}
	// 第一位要特别判断一下，因为可能第一位加上进位后超过了9
	if addend == 1 && digits[0] == 0 {
		digits = append([]int{1}, digits...)
	}

	return digits
}
