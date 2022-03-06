package math

import (
	"sort"
)

/** 最接近的三数之和
给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
返回这三个数的和。

假定每组输入只存在恰好一个解。

示例 1：

输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
示例 2：

输入：nums = [0,0,0], target = 1
输出：0

提示：

3 <= nums.length <= 1000
-1000 <= nums[i] <= 1000
-104 <= target <= 104
*/
func threeSumClosest(nums []int, target int) int {
	// 先排序
	// 1 遍历 i
	// 1.1 使得 start 为i+1，end为 length-1
	// 1.2 遍历 start<end
	// 1.2.1 计算 sum = nums[i] + nums[left] + nums[end]
	// 1.2.2 判断差距，target-res 与 target-sum
	// 1.2.3 如果sum差距小，则sum给res
	// 1.2.4 接着，根据sum与target比较大小改变start与end的位置
	// 1.2.4 大则减小end，小则放大start
	// 1.2.4 相等则返回
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-1; i++ {
		start, end := i+1, len(nums)-1
		for start < end {
			sum := nums[i] + nums[start] + nums[end]
			// 更接近
			if abs(target-sum) < abs(target-res) {
				res = sum
			}

			if sum > target {
				end--
			} else if sum < target {
				start++
			} else {
				// 一样则换咯
				return res
			}
		}
	}
	return res
}

func abs(x int) int {
	if x > 0 {
		return x
	}

	return -x
}
