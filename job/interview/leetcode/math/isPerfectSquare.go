package math

/**  有效的完全平方数

给定一个 正整数 num ，编写一个函数，如果 num 是一个完全平方数，则返回 true ，否则返回 false 。
进阶：不要 使用任何内置的库函数，如  sqrt 。

示例 1：

输入：num = 16
输出：true
示例 2：

输入：num = 14
输出：false

提示：
1 <= num <= 2^31 - 1
*/
func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}

	// 二分查找，不断用中位数去命中
	left, right := 0, num
	for left <= right {
		mid := (left + right) / 2
		s := mid * mid
		if s < num {
			left = mid + 1
		} else if s > num {
			right = mid - 1
		} else {
			return true
		}
	}

	return false
}
