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
	left, right, winWidth, maxLen, maxStart := 0, 0, 1, 0, 0
	for i, _ := range s {
		// 每次遍历都移动一位
		left = i - 1
		right = i + 1

		// left 找一样的
		for left >= 0 && s[i] == s[left] {
			winWidth++
			left--
		}

		// right 找一样的
		for right < len(s) && s[i] == s[right] {
			winWidth++
			right++
		}

		// 找一样left和right两边一样的
		for left >= 0 && right < len(s) && s[left] == s[right] {
			winWidth += 2
			right++
			left--
		}

		if winWidth > maxLen {
			maxLen = winWidth
			maxStart = left
		}

		winWidth = 1
	}

	return s[maxStart+1 : maxStart+maxLen+1]
}
