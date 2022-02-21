package array

import "fmt"

func main() {
	arr := [][]int{
		{3, 5, 8, 10},
		{4, 6, 10, 11},
		{7, 11, 12, 13},
		{9, 13, 14, 15},
	}

	x, y := FindAnyOneInSorted2DArray(arr, 13)
	fmt.Printf("普通值13[2,3]：x = %d, y = %d\n", x, y)

	x, y = FindAnyOneInSorted2DArray(arr, 10)
	fmt.Printf("普通值10[0,3]：x = %d, y = %d\n", x, y)

	x, y = FindAnyOneInSorted2DArray(arr, 3)
	fmt.Printf("边界值3[0，0]：x = %d, y = %d\n", x, y)

	x, y = FindAnyOneInSorted2DArray(arr, 15)
	fmt.Printf("边界值15[3，3]：x = %d, y = %d\n", x, y)

	x, y = FindAnyOneInSorted2DArray(arr, 9)
	fmt.Printf("边界值9[3,0]：x = %d, y = %d\n", x, y)

	x, y = FindAnyOneInSorted2DArray(arr, 100)
	fmt.Printf("不存在：x = %d, y = %d\n", x, y)
}

func FindAnyOneInSorted2DArray(arr [][]int, beLookUp int) (x, y int) {
	// 空检测 忽略
	// 负检测 忽略

	// 一维长度
	xLen := len(arr)
	// 二维长度
	yLen := len(arr[0])

	xPoint := 0
	yPoint := yLen - 1

	// 通过右上角判断
	// 小于右上角说明最后一列不存在，y--
	// 大于右上角说明第一行不存在，x++
	for yPoint >= 0 && xPoint <= xLen-1 {
		// 等于，找到了
		if arr[xPoint][yPoint] == beLookUp {
			return xPoint, yPoint
		}

		// 大于，忽略当前列
		if arr[xPoint][yPoint] > beLookUp {
			yPoint--
			continue
		} else {
			xPoint++
			continue
		}
	}

	return -1, -1
}
