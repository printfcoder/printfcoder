package string

// 以第一个串作为前缀，与下一个对比，找出共同串，然后共同串再和第三个比
// 以此类推，直到比完或者共同串已经为空，剩下的共同串就是最终的串
// 比较时，用较短来比即可
func longestCommonPrefixCrosswise(args ...string) string {
	if len(args) == 0 {
		return ""
	}

	prefix := args[0]
	for i := 1; i < len(args); i++ {
		prefix = lcp(prefix, args[i])
		if prefix == "" {
			break
		}
	}

	return prefix
}

func lcp(str1, str2 string) string {
	leng := min(len(str1), len(str2))
	idx := 0

	for idx < leng && str1[idx] == str2[idx] {
		idx++
	}

	// 返回任一个串的子串都行
	return str1[:idx]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
