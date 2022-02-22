package array

func generateYanHuiTriangle(numRows int) [][]int {
	// 公式
	// a[i][j] = a[i-1][j-1]+a[i-1][j-1]
	if numRows < 1 {
		return [][]int{}
	}

	ret := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		ret[i] = make([]int, i+1)
		ret[i][0] = 1
		ret[i][i] = 1
		for j := 1; j < i; j++ {
			ret[i][j] = ret[i-1][j-1] + ret[i-1][j]
		}
	}

	return ret
}
