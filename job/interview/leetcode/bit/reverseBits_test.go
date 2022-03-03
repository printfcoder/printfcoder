package bit

import (
	"fmt"
	"testing"
)

func TestReverseBits(t *testing.T) {
	x := reverseBits(uint32(33223))
	fmt.Println(x)
}
