package main

/**
除自身以外数组的乘积
给你一个整数数组nums，返回 数组answer，其中answer[i]等于nums中除nums[i]之外其余各元素的乘积。
题目数据 保证 数组nums之中任意元素的全部前缀元素和后缀的乘积都在 32 位 整数范围内。
请不要使用除法，且在O(n) 时间复杂度内完成此题。

示例 1:

输入: nums = [1,2,3,4]
输出: [24,12,8,6]
示例 2:

输入: nums = [-1,1,0,-3,3]
输出: [0,0,9,0,0]

提示：
2 <= nums.length <= 105
-30 <= nums[i] <= 30
保证 数组nums之中任意元素的全部前缀元素和后缀的乘积都在 32 位 整数范围内

进阶：你可以在 O(1)的额外空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）
*/

func productExceptSelf(nums []int) []int {
	// 两次遍历，先统一计算左边的值，再计算右边的值
	answer := make([]int, len(nums))
	answer[0] = 1
	for i := 1; i < len(nums); i++ {
		answer[i] = nums[i-1] * answer[i-1]
	}

	// 从右开始
	R := 1
	for i := len(nums) - 1; i > 0; i-- {
		answer[i] = answer[i] * R
		R *= nums[i]
	}

	return answer
}
