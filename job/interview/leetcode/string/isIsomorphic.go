package string

/** 同构字符串

给定两个字符串 s 和 t ，判断它们是否是同构的。
如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。
每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。

示例 1:
输入：s = "egg", t = "add"
输出：true

示例 2：
输入：s = "foo", t = "bar"
输出：false

示例 3：
输入：s = "paper", t = "title"
输出：true

提示：
1 <= s.length <= 5 * 104
t.length == s.length
s 和 t 由任意有效的 ASCII 字符组成
*/

func isIsomorphic(s string, t string) bool {
	if s == t {
		return true
	}

	m := make(map[uint8]uint8)
	mUsed := make(map[uint8]uint8)
	for i := 0; i < len(s); i++ {
		if v, ok := m[s[i]]; ok && v != t[i] {
			return false
		} else {
			// 没有被用过
			if v, ok := mUsed[t[i]]; ok && v != s[i] {
				return false
			}
			m[s[i]] = t[i]
			mUsed[t[i]] = s[i]
		}
	}

	return true
}
