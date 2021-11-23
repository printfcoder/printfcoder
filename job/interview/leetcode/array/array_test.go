package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTwoSum(t *testing.T) {
	Convey("Testing two sum", t, func() {
		x, y := twoSum(6, 2, 4, 5)
		So(x, ShouldEqual, 0)
		So(y, ShouldEqual, 1)
		x, y = twoSum(9, 2, 4, 5, 9)
		So(x, ShouldEqual, 1)
		So(y, ShouldEqual, 2)
	})
}
