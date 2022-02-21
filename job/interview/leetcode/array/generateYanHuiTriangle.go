package array

func generateYanHuiTriangle(numRows int) [][]int {
	// 公式
	// a[i][j] = a[i-1][j-1]+a[i-1][j-1]
	if numRows < 1 {
		return [][]int{}
	}

	ret := make([][]int, numRows)
	ret[0] = append(ret[0], 1)
	for i := 1; i < numRows; i++ {
		var cur []int
		for j := 0; j < i; j++ {
			// 注意边界
			if j == 0 {
				cur = append(cur, ret[i-1][j])
			} else {
				cur = append(cur, ret[i-1][j-1]+ret[i-1][j])
			}
		}

		ret[i] = append(ret[i], cur...)
	}

	return ret
}
