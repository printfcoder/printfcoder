package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTwoSumWithMap(t *testing.T) {
	Convey("Testing two sum", t, func() {
		x, y := twoSumWithMap(6, 2, 4, 5)
		So(x, ShouldEqual, 0)
		So(y, ShouldEqual, 1)
		x, y = twoSumWithSorting(9, 2, 4, 5, 9)
		So(x, ShouldEqual, 1)
		So(y, ShouldEqual, 2)
	})
}
