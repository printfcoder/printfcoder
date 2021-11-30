package main

func lengthOfLongestSubstring(s string) (sub string) {
	// 忽略长度检查

	// 滑动窗口，遍历字串，当前字串，也就是窗口右边的值，不在窗口中，则附加到窗口，长度加1
	// 如果存在，则丢掉
	window := make([]byte, 0)
	// 记录窗口起、始、余留坐标
	start, end, res := 0, 0, 0

	for start < len(s) && end < len(s) {
		if isExist(window, s[end]) {
			window = window[1:]
			start++
		} else {
			window = append(window, s[end])
			end++
			if res < len(window) {
				res = len(window)
			}
		}
	}

	return string(window[:])
}

func isExist(arr []byte, key byte) bool {
	for _, value := range arr {
		if value == key {
			return true
		}
	}
	return false
}
