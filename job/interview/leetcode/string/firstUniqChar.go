package string

/**
给定一个字符串s，找到 它的第一个不重复的字符，并返回它的索引 。如果不存在，则返回 -1。

示例 1：

输入: s = "leetcode"
输出: 0
示例 2:

输入: s = "loveleetcode"
输出: 2
示例 3:

输入: s = "aabb"
输出: -1
*/

func firstUniqChar(s string) int {
	result := make([]int, 26)
	for i := 0; i < len(s); i++ {
		result[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if result[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
