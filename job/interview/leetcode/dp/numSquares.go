package dp

/**
完全平方数
给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。

完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。

示例1：

输入：n = 12
输出：3
解释：12 = 4 + 4 + 4
示例 2：

输入：n = 13
输出：2
解释：13 = 4 + 9
*/

func numSquares(n int) int {
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		// 随便来一个大于j*j的就行
		minn := i
		for j := 1; j*j <= i; j++ {
			minn = min(minn, f[i-j*j])
		}
		// 加上f[j*j]本身，就是最后一个
		f[i] = minn + 1
	}
	return f[n]
}