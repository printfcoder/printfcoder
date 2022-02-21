package array

// FindAnyDuplicatedItemFromZero2Nm1 找到指定数组[0,n-1]中任一重复的数字，假设元数都小于n
// 不使用Map，避免O(n)的空间复杂度
// 不使用排序，避免T(nLog(n))的时间复杂度
// 采集数组标记法，将已知数组按索引排排坐，即数值与坐标相同
func FindAnyDuplicatedItemFromZero2Nm1(ar []int) (output int) {
	// 长度检测，忽略
	// 负数检测，忽略
	// 溢出检测，忽略

	for i := 0; i < len(ar); i++ {
		// 如果值与坐标不同，则找到它应该在的位置
		for ar[i] != i {
			if ar[i] == ar[ar[i]] {
				return ar[i]
			}

			temp := ar[ar[i]]
			ar[ar[i]] = ar[i]
			ar[i] = temp
		}
	}

	return
}

// FindAnyDuplicatedItemFromZero2Nm1WithDichotomy 找到指定数组[0,n-1]中任一重复的数字，假设元数都小于n
// 不使用Map，避免O(n)的空间复杂度
// 不使用排序，避免T(nLog(n))的时间复杂度
// 不改变原数组，使用二分法修改
func FindAnyDuplicatedItemFromZero2Nm1WithDichotomy(ar []int) (output int) {
	// 长度检测，忽略
	// 负数检测，忽略
	// 溢出检测，忽略

	return 0
}
