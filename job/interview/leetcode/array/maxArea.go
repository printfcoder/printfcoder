package array

/**
盛最多水的容器

给定一个长度为 n 的整数数组height。有n条垂线，第 i 条线的两个端点是(i, 0)和(i, height[i])。
找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为49。

示例 2：
输入：height = [1,1]
输出：1


提示：

n == height.length
2 <= n <= 105
0 <= height[i] <= 104
*/
func maxArea(height []int) int {
	// 本质是找到两个位置高度小者与两者之间的距离乘积最大值
	// 那只要从两边开始就可以了
	// 不断从两边逼近中间，找到最大值
	start, end, res := 0, len(height)-1, 0

	for start < end {
		// 前面的短
		if height[start] < height[end] {
			res = max(res, height[start]*(end-start))
			start++
		} else {
			res = max(res, height[end]*(end-start))
			end--
		}
	}

	return res
}

func max(i, j int) int {
	if i < j {
		return j
	}

	return i
}
