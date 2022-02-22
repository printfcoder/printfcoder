package array

// 给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
// 你可以 假设数组是非空的，并且给定的数组总是存在众数。

// 采用抵消法，从头开始遍历，前一个与后面的逐个对比，遇到一样的总数加1，不一样的总数减一，直到抵消完，剩下的就是目标数，因为大于一半的数肯定只有一个
// 当值一样时，说明二者相同，那么，count+1
// 否则count-1，代表不同
// count为0时说明抵挡
func majorityElement(input []int) int {
	output, count := input[0], 0
	for i := 0; i < len(input); i++ {
		if count == 0 {
			output, count = input[i], 1
		} else {
			if output == input[i] {
				count++
			} else {
				count--
			}
		}
	}

	return output
}
