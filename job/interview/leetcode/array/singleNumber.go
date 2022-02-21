package array

/**
  采用位运算的异或法，用0去依次异或，出现两次因为bit位都相同，会被异或为0抵消掉
*/
func singleNumber(inputs ...int) int {
	output := 0
	for i := 0; i < len(inputs); i++ {
		output ^= inputs[i]
	}

	return output
}
