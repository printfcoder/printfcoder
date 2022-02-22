package maps

// 给你一个字符串columnTitle ，表示 Excel 表格中的列名称。返回 该列名称对应的列序号。
//
// 例如：
//
// A -> 1
// B -> 2
// C -> 3
// ...
// Z -> 26
// AA -> 27
// AB -> 28
// ...
//
// 示例 1:
//
// 输入: columnTitle = "A"
// 输出: 1
// 示例2:
//
// 输入: columnTitle = "AB"
// 输出: 28
// 示例3:
//
// 输入: columnTitle = "ZY"
// 输出: 701
//
// 提示：
//
// 1 <= columnTitle.length <= 7
// columnTitle 仅由大写英文组成
// columnTitle 在范围 ["A", "FXSHRXW"] 内

func titleToNumber(columnTitle string) int {
	ret := 0
	l := len(columnTitle)
	// 倍数
	m := 1
	for i := l - 1; i >= 0; i-- {
		// 当前位余数
		k := columnTitle[i] - 'A' + 1
		// 当前位的进位数
		ret += int(k) * m
		m *= 26
	}

	return ret
}
