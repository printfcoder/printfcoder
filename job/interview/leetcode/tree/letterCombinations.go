package tree

/** 电话号码的字母组合
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

示例 1：
输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]

示例 2：
输入：digits = ""
输出：[]

示例 3：
输入：digits = "2"
输出：["a","b","c"]

提示：
0 <= digits.length <= 4
digits[i] 是范围 ['2', '9'] 的一个数字。
*/

var dic = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	// 用深度优先即可，将digits分开，比如23
	//             2
	//         a | b       | c
	//           | 3       |
	//   d  e  f | d  e  f | d  e  f
	//-> ad ae af| bd be bf| cd ce cf
	// 再多几层都一样展开即可，最后依次遍历深度
	ret, tmp := make([]string, 0), make([]string, 0)
	if len(digits) == 0 {
		return ret
	}

	// 第一层
	ret = backtrack("", string(digits[0]))
	for i := 1; i < len(digits); i++ {
		tmp = nil
		for j := range ret {
			tmp = append(tmp, backtrack(ret[j], string(digits[i]))...)
		}

		ret = tmp
	}

	return ret
}

func backtrack(combination, digit string) []string {
	arrs := dic[digit]
	ret := make([]string, 0)
	for i := 0; i < len(arrs); i++ {
		ret = append(ret, combination+arrs[i])
	}

	return ret
}
