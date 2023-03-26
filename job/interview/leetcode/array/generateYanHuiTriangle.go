package array

import "fmt"

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

func f1(n int) [][]int {

	if n <= 0 {
		return [][]int{}
	}

	var nums = make([]int, n)
	nums[0] = 1
	for i := 0; i < n; i++ {
		var nums2 = make([]int, n)
		nums2[0] = 1

		fmt.Printf("%d ", nums[0])
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", nums[j]+nums[j-1])
			nums2[j] = nums[j] + nums[j-1]
		}
		nums = nums2
		fmt.Println("")
	}

	return [][]int{}
}
