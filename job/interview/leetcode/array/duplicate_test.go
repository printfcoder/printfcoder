package array

import "fmt"

func TextDuplicatedItem() {
	ar := []int{2, 1, 2, 4, 1}
	fmt.Printf("dulicated num is %d \n", FindAnyDuplicatedItemFromZero2Nm1(ar))
	FindAnyDuplicatedItemFromZero2Nm1WithDichotomy([]int{1})
}
