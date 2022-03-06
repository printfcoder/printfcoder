package string

import "strings"

/** 单词规律
给定一种规律 pattern 和一个字符串 s ，判断 s 是否遵循相同的规律。

这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。

示例1:
输入: pattern = "abba", str = "dog cat cat dog"
输出: true
示例 2:

输入:pattern = "abba", str = "dog cat cat fish"
输出: false
示例 3:

输入: pattern = "aaaa", str = "dog cat cat dog"
输出: false

提示:
1 <= pattern.length <= 300
pattern 只包含小写英文字母
1 <= s.length <= 3000
s 只包含小写英文字母和 ' '
s 不包含 任何前导或尾随对空格
s 中每个单词都被 单个空格 分隔
*/
func wordPattern(pattern string, s string) bool {
	// 用于存储pattern字节对s元素的关系
	mP2S := make(map[uint8]string)
	// 用于检查映射关系是否正确，因为有可能pattern中的两个对应一个s中的值，这是不行的，毕竟要双向绑定
	mS2P := make(map[string]uint8)
	sArray := strings.Split(s, " ")
	if len(pattern) != len(sArray) {
		return false
	}

	for i := range pattern {
		// 映射表中存在该关系，则值一致
		if v, ok := mP2S[pattern[i]]; ok {
			// 值要一致
			if v != sArray[i] {
				return false
			}
		}

		// 没有关系，但是该值已经有其它pattern使用了
		if v, ok := mS2P[sArray[i]]; ok {
			if v != pattern[i] {
				return false
			}
		}
		// 没有关系，保留关系
		mP2S[pattern[i]] = sArray[i]
		mS2P[sArray[i]] = pattern[i]
	}

	return true
}
