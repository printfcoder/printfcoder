package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReverseString(t *testing.T) {
	Convey("Testing reverse int", t, func() {
		x := reverseInt(2345)
		So(x, ShouldEqual, 5432)

		x = reverseInt(-2345)
		So(x, ShouldEqual, -5432)

		x = reverseInt(2340)
		So(x, ShouldEqual, 432)
	})
}

func TestReverseInt(t *testing.T) {
	Convey("Testing reverse int", t, func() {
		x := reverseInt(2345)
		So(x, ShouldEqual, 5432)

		x = reverseInt(-2345)
		So(x, ShouldEqual, -5432)

		x = reverseInt(2340)
		So(x, ShouldEqual, 432)
	})
}
