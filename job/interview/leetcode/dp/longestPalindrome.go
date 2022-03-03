package dp

//给你一个字符串 s，找到 s 中最长的回文子串。
//
//
//
// 示例 1：
//
//
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
//
//
// 示例 2：
//
//
//输入：s = "cbbd"
//输出："bb"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 仅由数字和英文字母组成

func longestPalindrome(s string) string {
	// 采用动态规划完成
	// dp[i][j]，i为第i个回文串，j为该串中间值的坐标
	// 当dp[i]
	// 状态方程：
	// 在[0,i)范围内，头尾与dp时，dp[i][j] = s[i-len(dp[i-1])-1]+dp[i-1][j]+s[i]
	// 在[0,i)，头尾不一致时， 范围内dp[i][j] = dp[i-1][j]

	dp := make([]string, len(s))
	dp[0] = s[:1]
	if s[0] == s[1] {
		dp[1] = s[:2]
	} else {
		dp[1] = dp[0]
	}
	for i := 2; i < len(s); i++ {
		//	if s[i-len(dp[i-1])-1:len()/] != s[i] {
		if true {
			dp[i] = dp[i-1]
		} else {
			dp[i] = string(s[i-len(dp[i-1])-1]) + dp[i-1] + string(s[i])
		}
	}

	return dp[len(s)-1]
}
