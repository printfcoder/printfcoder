package math

/**
给你一个整数 n ，请你判断 n 是否为 丑数 。如果是，返回 true ；否则，返回 false 。
丑数 就是只包含质因数2、3、5的正整数。

示例 1：

输入：n = 6
输出：true
解释：6 = 2 × 3
示例 2：

输入：n = 8
输出：true
解释：8 = 2 × 2 × 2
示例 3：

输入：n = 14
输出：false
解释：14 不是丑数，因为它包含了另外一个质因数7 。
示例 4：

输入：n = 1
输出：true
解释：1 通常被视为丑数。

提示：
-231 <= n <= 231 - 1
*/
func isUgly(n int) bool {
	if n <= 0 {
		return false
	}

	factors := []int{2, 3, 5}
	// n = 2^a*3^b*5^c
	// 反复2，3，5去除n，余为0就说明对应的因子是在其中的
	// 直到所有余不为0时，说明除不尽了，此时n剩余量应该为1
	for _, f := range factors {
		for n%f == 0 { // 第二个for是除尽当前因子
			n = n / f
		}
	}

	return n == 1
}
