package math

/**
分割回文串
给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。

回文串 是正着读和反着读都一样的字符串。

示例 1：

输入：s = "aab"
输出：[["a","a","b"],["aa","b"]]
示例 2：

输入：s = "a"
输出：[["a"]]

提示：

1 <= s.length <= 16
s 仅由小写英文字母组成
*/

func partition(s string) [][]string {
	ret := make([][]string, 0)
	// 每次先拆i个，一个一个拆
	for i := 0; i < len(s); i++ {
		// 窗口开始查找
		// 每次窗口右边挪，挪完再进行下一轮
		ret = append(ret, []string{})
		for j := i; j < len(s); j++ {
			if isPalindrome(s[i:j]) {
				ret[i] = append(ret[i], s[i:j])
			}
		}
	}

	return ret
}

func isPalindrome(s string) bool {
	// 双指针
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}
