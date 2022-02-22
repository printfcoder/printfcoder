package array

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
	// 双指针法，fast在后面，负责为slow找匹配的
	if len(nums) < 3 {
		return false
	}

	// res为长度
	slow, fast, res := 0, 1, make([]int, 3)
	res[0] = nums[0]
	for slow < fast && fast < len(nums)+1 {
		if nums[slow] < nums[fast] {
			res = append(res, nums[fast])
			if len(res) == 3 {
				return true
			}
		}

		if fast < len(nums) {
			fast++
		} else {
			// fast 到头，没找到
			// 此时slow要向前，也即时res要置空，重新从下一个位置来
			res = make([]int, 3)
			slow++
			fast = slow + 1
		}
	}

	return false
}
