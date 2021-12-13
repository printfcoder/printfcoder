package main

//给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
//
// 请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
//
// 注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并
//的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。
//
//
//
// 示例 1：
//
//
//输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
//输出：[1,2,2,3,5,6]
//解释：需要合并 [1,2,3] 和 [2,5,6] 。
//合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。
//
//
// 示例 2：
//
//
//输入：nums1 = [1], m = 1, nums2 = [], n = 0
//输出：[1]
//解释：需要合并 [1] 和 [] 。
//合并结果是 [1] 。
//
//
// 示例 3：
//
//
//输入：nums1 = [0], m = 0, nums2 = [1], n = 1
//输出：[1]
//解释：需要合并的数组是 [] 和 [1] 。
//合并结果是 [1] 。
//注意，因为 m = 0 ，所以 nums1 中没有元素。nums1 中仅存的 0 仅仅是为了确保合并结果可以顺利存放到 nums1 中。
//
//
//
//
// 提示：
//
//
// nums1.length == m + n
// nums2.length == n
// 0 <= m, n <= 200
// 1 <= m + n <= 200
// -10⁹ <= nums1[i], nums2[j] <= 10⁹
//
//
//
//
// 进阶：你可以设计实现一个时间复杂度为 O(m + n) 的算法解决此问题吗？
// Related Topics 数组 双指针 排序 👍 1211 👎 0

// 正向双指针法，初始化一数组用于放置合并后的结果，两个指针分p1,p2别nums1和nums2开头开始遍历，每次迭代分别对比，小的append进入新数组，同时小的指针后移，继续下一步
// 直到某一个数组遍历完，然后未对比到的剩余元素直接append到新数组中
// O(m+n) time，O(m+n) space
func mergeWithTwoPointsFromHead(nums1 []int, nums2 []int) (sorted []int) {
	m, n := len(nums1), len(nums2)
	sorted = make([]int, 0, m+n)
	p1, p2 := 0, 0

	for {
		// 到尽头了
		if p1 == m {
			sorted = append(sorted, nums2[p2:]...)
			break
		}

		if p2 == n {
			sorted = append(sorted, nums1[p1:]...)
			break
		}

		if nums1[p1] < nums2[p2] {
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			p2++
		}
	}

	return
}

// 定义两个指针，p1,p2分别指向两个数组的末位，然后各自往前挪，每挪一次互比一次，大者放到最后一位
// 需要注意的是，我们要在原来的数组上进行操作，所以要保证数组的容量可以放两个数组合并后的元素
// 我们用第一个数组nums1来当，所以要先扩容
func mergeWithTowPointsFromTail(nums1 []int, nums2 []int) (sorted []int) {
	p1, p2, tail := len(nums1)-1, len(nums2)-1, len(nums1)+len(nums2)-1
	// 用nums1来当存储数组
	if cap(nums1) != len(nums1)+len(nums2) {
		nums1 = append(nums1, nums2...)
	}

	for ; p1 >= 0 || p2 >= 0; tail-- {
		var cur int
		if p1 == -1 {
			cur = nums2[p2]
			p2--
		} else if p2 == -1 {
			cur = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			cur = nums1[p1]
			p1--
		} else {
			cur = nums2[p2]
			p2--
		}
		nums1[tail] = cur
	}

	return nums1
}
