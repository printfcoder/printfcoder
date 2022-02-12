package main

/**
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

注意：若s 和 t中每个字符出现的次数都相同，则称s 和 t互为字母异位词。

示例1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false


提示:

1 <= s.length, t.length <= 5 * 104
s 和 t仅包含小写字母


进阶:如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
*/

// 使用一个字符串数组表来表示每个字符出现的次数
// s塞进去，t的用来抵消
// 最终为0即可
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	comp := make([]int, 26)

	for _, sr := range s {
		comp[sr-'a']++
	}

	for _, tr := range t {
		comp[tr-'a']--
	}

	for i := 0; i < 26; i++ {
		if comp[i] != 0 {
			return false
		}
	}
	return true
}
