package math

import "sort"

/**  三数之和
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。
示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
示例 2：

输入：nums = []
输出：[]
示例 3：

输入：nums = [0]
输出：[]
提示：

0 <= nums.length <= 3000
-105 <= nums[i] <= 105
*/
func threeSum(nums []int) [][]int {
	ret := make([][]int, 0)

	if len(nums) < 3 {
		return ret
	}

	// 先排序
	// 1. i 表示当前位置，l，r为基于i开始遍历的后续元素的两端坐标
	// 1.1 开始遍历
	// 1.1.1 num[i]>0时，跳出，后续不可能==0了
	// 1.1.2 遍历i后元素
	// 1.1.3 跳过num[i-1]=num[i]的情况
	// 1.1.3 num[i] + num[l] + num[r]
	// 1.1.3.1 if == 0，进行 l, r 下一次遍历，且要路过l,r临近的相等值
	// 1.1.3.2 if > 0，说明r大了，r--
	// 1.1.3.2 if < 0，说明l小了，l++
	// 1.1.4 i++
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return ret
		}
		// 上一个和当前值一样，则跳过
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for l, r := i+1, len(nums)-1; r > l; {
			if nums[i]+nums[l]+nums[r] == 0 {
				ret = append(ret, []int{nums[i], nums[l], nums[r]})
				// 避开相同的值
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				for l < r && nums[l] == nums[l+1] {
					l++
				}

				r--
				l++
			} else if nums[i]+nums[l]+nums[r] > 0 {
				r--
			} else {
				l++
			}
		}
	}

	return ret
}
