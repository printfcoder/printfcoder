package dp

/**
最大子序和
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

子数组 是数组中的一个连续部分。

示例 1：

输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组[4,-1,2,1] 的和最大，为6 。
示例 2：

输入：nums = [1]
输出：1
示例 3：

输入：nums = [5,4,-1,7,8]
输出：23

提示：

1 <= nums.length <= 105
-104 <= nums[i] <= 104

进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。
*/
func maxSubArray(nums []int) int {
	// [-2,1,-3,4,-1,2,1,-5,4]
	// 直接依题意解
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}

		if max < nums[i] {
			max = nums[i]
		}

	}

	return max
}

func maxSubArrayDP(nums []int) int {
	// [-2,1,-3,4,-1,2,1,-5,4]
	// 使用动态规划求解
	// 转换方程：
	// 在[0,i)之间：
	// 上一个最大解dp[i-1]到下一个dp[i]依旧为增长： dp[i] = dp[i-1] + nums[i]
	// 上一个最大解dp[i-1]到下一个dp[i]为减小： dp[i] = nums[i]
	// 注意：不能通过判断nums[i]大于0来进行状态转换，
	dp, res := make([]int, len(nums)), nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}

		if res < dp[i] {
			res = dp[i]
		}
	}

	return res
}
