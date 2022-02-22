package math

/**
3的幂
给定一个整数，写一个函数来判断它是否是 3的幂次方。如果是，返回 true ；否则，返回 false 。

整数 n 是 3 的幂次方需满足：存在整数 x 使得 n == 3x

示例 1：

输入：n = 27
输出：true
示例 2：

输入：n = 0
输出：false
示例 3：

输入：n = 9
输出：true
示例 4：

输入：n = 45
输出：false

提示：

-231 <= n <= 231 - 1

进阶：你能不使用循环或者递归来完成本题吗？
*/
func isPowerOfThree(n int) bool {
	// 假设是3的幂次
	// 那一定不是偶数，3的次幂一定是1，3，7，9结尾，0，2，4，5，6，8都不是
	last := n % 10
	if last == 1 || last == 3 || last == 7 || last == 9 {
		return 1162261467%n == 0
	}

	return false
}
