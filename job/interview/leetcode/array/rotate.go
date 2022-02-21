package array

/**
给你一个数组，将数组中的元素向右轮转 k个位置，其中k是非负数。

示例 1:

输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]
示例2:

输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右轮转 1 步: [99,-1,-100,3]
向右轮转 2 步: [3,99,-1,-100]


提示：

1 <= nums.length <= 105
-231 <= nums[i] <= 231 - 1
0 <= k <= 105

*/
func rotate(nums []int, k int) {
	l := len(nums)
	newNums := make([]int, l)
	for i := 0; i < l; i++ {
		newNums[(i+k)%l] = nums[i]
	}

	copy(nums, newNums)
}

func reverse(nums []int) {
	l := len(nums)
	for i := 0; i < l/2; i++ {
		// 当前位置与对称位置互换
		nums[i], nums[l-i-1] = nums[l-i-1], nums[i]
	}
}

func rotateWithReverse(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	// 翻转k次，说明后k%l的部分顺序倒过来放到前面，而前部分顺序不动放到后面
	// 故先翻转后k%l部分，此时已经在前k%l位
	reverse(nums[:k])
	// 翻转后k%l位，此时已经在前k%l位
	reverse(nums[k:])
}
