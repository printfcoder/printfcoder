package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGenerateYanHuiTriangle(t *testing.T) {
	Convey("Testing remove duplicates", t, func() {
	 output := generateYanHuiTriangle(5)
		So(l, ShouldEqual, 4)
		So(output, ShouldResemble, []int{1, 2, 4, 8})
	})
}
