package math

/**
给定一个整数 n ，返回 n! 结果中尾随零的数量。

提示n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1

示例 1：

输入：n = 3
输出：0
解释：3! = 6 ，不含尾随 0
示例 2：

输入：n = 5
输出：1
解释：5! = 120 ，有一个尾随 0
示例 3：

输入：n = 0
输出：0

提示：

0 <= n <= 104

进阶：你可以设计并实现对数时间复杂度的算法来解决此问题吗？
*/
func trailingZeroes(n int) int {
	var factorialValue int64
	factorialValue = 1
	for i := n; i > 0; i-- {
		factorialValue *= int64(i)
	}

	l := 0
	for {
		if factorialValue%10 == 0 {
			l++
			factorialValue = factorialValue / 10
		} else {
			break
		}
	}

	return l
}
