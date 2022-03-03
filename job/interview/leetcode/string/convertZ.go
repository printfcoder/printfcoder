package string

import "strings"

/**
Z 字形变换
将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行Z 字形排列。

比如输入字符串为 "PAYPALISHIRING"行数为 3 时，排列如下：

P   A   H   N
A P L S I I G
Y   I   R
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);


示例 1：

输入：s = "PAYPALISHIRING", numRows = 3
输出："PAHNAPLSIIGYIR"
示例 2：
输入：s = "PAYPALISHIRING", numRows = 4
输出："PINALSIGYAHRPI"
解释：
P     I    N
A   L S  I G
Y A   H R
P     I
示例 3：

输入：s = "A", numRows = 1
输出："A"

提示：
1 <= s.length <= 1000
s 由英文字母（小写和大写）、',' 和 '.' 组成
1 <= numRows <= 1000
*/
func convertZ(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	// 定义numRows个数组
	// 每个数组就是就是每行的数据
	// 依次遍历s，将字符串从下到上，再从上到上填入上面的三个数组中即可
	// 关键是记录翻转，每次循环numRows次数时，倒过来
	// i表示数组顺序，flag表示当前是正向还是反向
	// true表示正向，i为[0, numRows-1]
	// false表示正向，i为[numRows-1, 0]
	ret := make([]string, numRows)
	i, flag := 0, true
	for _, c := range s {
		ret[i] += string(c)

		// 正的
		// todo flag 可以优化成用标量表示，即与i配合算出坐标，而不是单纯是方向
		if flag {
			// 到头了
			if i == numRows-1 {
				flag = false
				i--
				continue
			}
			i++
		}
		// 反的
		if !flag {
			// 到头了
			if i == 0 {
				flag = true
				i++
				continue
			}
			i--
		}
	}

	return strings.Join(ret, "")
}
