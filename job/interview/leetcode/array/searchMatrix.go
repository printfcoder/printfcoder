package main

func searchMatrix(arr [][]int, target int) bool {
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
		if arr[xPoint][yPoint] == target {
			return true
		}

		// 大于，忽略当前列
		if arr[xPoint][yPoint] > target {
			yPoint--
		} else {
			xPoint++
			continue
		}
	}

	return false
}
