package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPlusOne(t *testing.T) {
	Convey("Testing plusOne", t, func() {
		output := plusOne([]int{0})
		So(output, ShouldResemble, []int{1})
		output = plusOne([]int{1, 2, 2, 4, 4, 8})
		So(output, ShouldResemble, []int{1, 2, 2, 4, 4, 9})
		output = plusOne(output)
		So(output, ShouldResemble, []int{1, 2, 2, 4, 5, 0})
		output = plusOne([]int{9, 9})
		So(output, ShouldResemble, []int{1, 0, 0})
	})
}
