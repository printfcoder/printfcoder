package bit

/** 计算各个位数不同的数字个数
给定一个非负整数 n，计算各位数字都不同的数字 x 的个数，其中 0 ≤ x < 10n 。

示例:

输入: 2
输出: 91
解释: 答案应为除去 11,22,33,44,55,66,77,88,99 外，在 [0,100) 区间内的所有数字。

https://leetcode-cn.com/problems/count-numbers-with-unique-digits/
*/
func countNumbersWithUniqueDigits(n int) int {
	// 每一位都不同
	// 1:  10，0，1，2，3，4，5，6，7，8，9
	// 2:  9*9+dp[0]， 除11,22,33,44,55,66,77,88,99
	// 3:  9*9*8+dp[1], 除11,22,33,44,55,66,77,88,99, 111,222,333,444,555,666,777,888,999, 112,113....998
	// 本质就是排列组合，每个数字只能用一次，n代表数字的数量
	// 公式为：dp[i] = dp[i-1] + 9*9*8*7...
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}

	dp := make([]int, n)
	dp[0] = 10
	dp[1] = 91

	for i := 2; i < n; i++ {
		s := 9
		for j := 0; j < i; j++ {
			s *= 9 - j
		}

		dp[i] = dp[i-1] + s
	}

	return dp[n-1]
}
