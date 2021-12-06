package main

//给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
//
// 不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
//
// 元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
//
//
//
// 说明:
//
// 为什么返回数值是整数，但输出的答案是数组呢?
//
// 请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。
//
// 你可以想象内部操作如下:
//
//
//// nums 是以“引用”方式传递的。也就是说，不对实参作任何拷贝
//int len = removeElement(nums, val);
//
//// 在函数里修改输入数组对于调用者是可见的。
//// 根据你的函数返回的长度, 它会打印出数组中 该长度范围内 的所有元素。
//for (int i = 0; i < len; i++) {
//    print(nums[i]);
//}
//
//
//
//
// 示例 1：
//
//
//输入：nums = [3,2,2,3], val = 3
//输出：2, nums = [2,2]
//解释：函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。你不需要考虑数组中超出新长度后面的元素。例如，函数返回的新长度为 2 ，而
//nums = [2,2,3,3] 或 nums = [2,2,0,0]，也会被视作正确答案。
//
//
// 示例 2：
//
//
//输入：nums = [0,1,2,2,3,0,4,2], val = 2
//输出：5, nums = [0,1,4,0,3]
//解释：函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。注意这五个元素可为任意顺序。你不需要考虑数组中超出新长度后面
//的元素。
//
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 100
// 0 <= nums[i] <= 50
// 0 <= val <= 100
// 设置两个指针a,b，a，b同时在第一个元素，如果b所指的元素等于val，则b滑动到下一个元素直到不是val的位置a'，然后将a'的值给a，
// 赋值后，ab同时滑向各自的下一个元素，b继续从它当前位置继续重复动作，直到数组末尾
// O(N) Time，O(1) Space
func removeElementWithTwoPoints(nums []int, val int) []int {
	// 省略长度与边界检查
	a := 0
	l := len(nums)
	for b := 0; b < l; b++ {
		if nums[b] != val {
			nums[a] = nums[b]
			a++
		}
	}

	return nums[:a]
}

// 使用首尾指针法，设置两个指针head, tail，开始分别指向首尾
// 当head等于val时，将tail的值赋给head，同时，tail向前挪一位，继续这个过程
// 当head不等于val时，head挪向后一位，继续这个过程
// 整个过程在head的位置小于tail时进行
// O(N) Time，O(1) Space
func removeElementWithHeadTailPoints(nums []int, val int) []int {
	head, tail := 0, len(nums)
	for head < tail {
		if nums[head] == val {
			nums[head] = nums[tail-1]
			tail--
		} else {
			head++
		}
	}

	return nums[:head]
}
