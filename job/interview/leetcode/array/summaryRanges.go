package array

import "strconv"

/**
228. 汇总区间
给定一个  无重复元素 的 有序 整数数组 nums 。

返回 恰好覆盖数组中所有数字 的 最小有序 区间范围列表 。也就是说，nums 的每个元素都恰好被某个区间范围所覆盖，并且不存在属于某个范围但不属于 nums 的数字 x 。

列表中的每个区间范围 [a,b] 应该按如下格式输出：

"a->b" ，如果 a != b
"a" ，如果 a == b

示例 1：

输入：nums = [0,1,2,4,5,7]
输出：["0->2","4->5","7"]
解释：区间范围是：
[0,2] --> "0->2"
[4,5] --> "4->5"
[7,7] --> "7"

示例 2：
输入：nums = [0,2,3,4,6,8,9]
输出：["0","2->4","6","8->9"]
解释：区间范围是：
[0,0] --> "0"
[2,4] --> "2->4"
[6,6] --> "6"
[8,9] --> "8->9"

提示：

0 <= nums.length <= 20
-231 <= nums[i] <= 231 - 1
nums 中的所有值都 互不相同
nums 按升序排列
*/
func summaryRanges(nums []int) []string {
	slow, ret := 0, make([]string, 0)
	for i, l := 0, len(nums); i < l; {
		// 记录当前值
		slow = i
		// 找到递增结束点
		for i += 1; i < l && nums[i-1]+1 == nums[i]; i++ {
		}

		cur := strconv.Itoa(nums[slow])
		// 此时i-1就是最后一位
		// 如果位点小于最后一位，说明不是一个数了
		if nums[slow] < nums[i-1] {
			cur += "->" + strconv.Itoa(nums[i-1])
		}

		// 等于的不需要处理

		ret = append(ret, cur)
	}

	return ret
}
