package array

// 给你一个数组 points ，其中 points[i] = [xi, yi] 表示 X-Y 平面上的一个点。求最多有多少个点在同一条直线上。

func maxPoints(points [][]int) int {
	// (a,b),(c, d),(e,f)在一条直线的公式：
	// 以(c, d)为原点，(a,b)，(e,f)到它的斜率相同：
	// (c-a)/(d-b)=(e-c)/(f-d)
	// 推导为：(c-a)(f-d)=(d-b)(e-c)
	// 两两组合，斜率为key，找其它组，找到一个斜率相同的则key+1
	// 斜率有精度问题，使用长除法保存精度

	return 0
}
