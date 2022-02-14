package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSingleNumber(t *testing.T) {
	Convey("Testing two sum", t, func() {
		x := singleNumber(6, 2, 4, 3, 3, 4, 2)
		So(x, ShouldEqual, 6)
	})
}
