package bit

/**
  二进制求和

给你两个二进制字符串，返回它们的和（用二进制表示）。

输入为 非空 字符串且只包含数字 1 和 0。



示例 1:

输入: a = "11", b = "1"
输出: "100"
示例 2:

输入: a = "1010", b = "1011"
输出: "10101"


提示：

每个字符串仅由字符 '0' 或 '1' 组成。
1 <= a.length, b.length <= 10^4
字符串如果不是 "0" ，就都不含前导零。
*/
func addBinary(a string, b string) string {
	builder := ""
	// 进位，要放到外面，可能会有最后一个进位
	ca := 0

	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0; {
		sum := ca

		if i >= 0 {
			sum += int(a[i] - '0')
		}

		if j >= 0 {
			sum += int(b[j] - '0')
		}

		if sum%2 == 0 {
			builder = "0" + builder
		} else {
			builder = "1" + builder
		}

		ca = sum / 2
		i--
		j--
	}
	if ca == 1 {
		builder = "1" + builder
	}

	return builder
}
