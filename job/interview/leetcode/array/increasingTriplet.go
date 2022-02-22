package array

import "math"

/**
递增的三元子序列
给你一个整数数组nums ，判断这个数组中是否存在长度为 3 的递增子序列。

如果存在这样的三元组下标 (i, j, k)且满足 i < j < k ，使得nums[i] < nums[j] < nums[k] ，返回 true ；否则，返回 false 。

示例 1：

输入：nums = [1,2,3,4,5]
输出：true
解释：任何 i < j < k 的三元组都满足题意
示例 2：

输入：nums = [5,4,3,2,1]
输出：false
解释：不存在满足题意的三元组
示例 3：

输入：nums = [2,1,5,0,4,6]
输出：true
解释：三元组 (3, 4, 5) 满足题意，因为 nums[3] == 0 < nums[4] == 4 < nums[5] == 6

提示：

1 <= nums.length <= 5 * 105
-231 <= nums[i] <= 231 - 1


进阶：你能实现时间复杂度为 O(n) ，空间复杂度为 O(1) 的解决方案吗？
*/
func increasingTriplet(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}

	// 假设老二是最大的，开始遍历
	first, second := nums[0], math.MaxInt32
	for i := 1; i < n; i++ {
		// 现在的比老二大，说明找到老三了
		num := nums[i]
		if num > second {
			return true
		} else if num > first { // 小于老三，但是比老大要大，说明是老二
			second = num
		} else { // 比老大小或等于，说明要当老大
			first = num
		}
	}
	return false
}
