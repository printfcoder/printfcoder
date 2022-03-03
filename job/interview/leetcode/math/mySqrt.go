package math

import "math"

func mySqrt(x int, precision float32) float32 {
	xf := float32(x)
	low, high, res := float32(0), xf, float32(-1)
	for {
		mid := low + (high-low)/2
		if math.Abs(float64((float32(x)/mid)-mid)) < float64(precision) {
			res = mid
			break
		}
		if mid*mid <= high {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return res
}
