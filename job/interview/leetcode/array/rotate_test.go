package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRotate(t *testing.T) {
	Convey("Testing two sum", t, func() {
		in := []int{1, 2, 3, 4, 5, 6, 7}
		rotateWithReverse(in, 3)
		So(in, ShouldResemble, []int{5, 6, 7, 1, 2, 3, 4})
		in = []int{1, 2, 3, 4, 5, 6, 7}
		rotate(in, 3)
		So(in, ShouldResemble, []int{5, 6, 7, 1, 2, 3, 4})
	})
}
