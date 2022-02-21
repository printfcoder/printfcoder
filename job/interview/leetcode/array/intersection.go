package array

//给定两个数组 nums1 和 nums2 ，返回 它们的交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。
//
//
//
// 示例 1：
//
//
//输入：nums1 = [1,2,2,1], nums2 = [2,2]
//输出：[2]
//
//
// 示例 2：
//
//
//输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出：[9,4]
//解释：[4,9] 也是可通过的
//
//
//
//
// 提示：
//
//
// 1 <= nums1.length, nums2.length <= 1000
// 0 <= nums1[i], nums2[i] <= 1000

// 将较少的数组的值存入map，另一个来命中，没有命中的删掉这个元素
func intersect(nums1 []int, nums2 []int) (res []int) {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}

	m := make(map[int]int)
	for _, num := range nums1 {
		// 因为要返回最大重复数量，所以要记录次数
		m[num]++
	}

	for _, num := range nums2 {
		if _, ok := m[num]; ok && m[num] > 0 {
			// 命中一次后减一次
			m[num]--
			res = append(res, num)
		}
	}

	return
}
