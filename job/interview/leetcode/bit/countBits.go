package bit

/** 比特位计数
给你一个整数 n ，对于 0 <= i <= n 中的每个 i ，计算其二进制表示中 1 的个数 ，返回一个长度为 n + 1 的数组 ans 作为答案。

示例 1：

输入：n = 2
输出：[0,1,1]
解释：
0 --> 0
1 --> 1
2 --> 10
示例 2：

输入：n = 5
输出：[0,1,1,2,1,2]
解释：
0 --> 0
1 --> 1
2 --> 10
3 --> 11
4 --> 100
5 --> 101

提示：

0 <= n <= 105

*/
func countBits(n int) []int {
	// i是递增的，每次n/2
	/**
	0 --> 0
	1 --> 1
	2 --> 10
	3 --> 11
	4 --> 100
	5 --> 101
	6 --> 110
	*/

	// 奇数比前面的偶数多1
	// 偶数与/2的一样多
	// 所以我们从0开始
	ret := make([]int, n+1)
	ret[0] = 0
	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			ret[i] = ret[i-1] + 1
		} else {
			ret[i] = ret[i/2]
		}
	}

	return ret
}
