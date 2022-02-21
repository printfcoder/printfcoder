package array

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRemoveElementWithTwoPoints(t *testing.T) {
	Convey("Testing remove duplicates", t, func() {
		output := removeElementWithTwoPoints([]int{1, 2, 2, 4, 4, 8}, 2)
		So(len(output), ShouldEqual, 4)
		So(output, ShouldResemble, []int{1, 4, 4, 8})
	})
}

func TestRemoveElementWithHeadTailPoints(t *testing.T) {
	Convey("Testing remove duplicates", t, func() {
		output := removeElementWithHeadTailPoints([]int{1, 2, 2, 4, 4, 8}, 2)
		So(len(output), ShouldEqual, 4)
		So(output, ShouldResemble, []int{1, 8, 4, 4})
	})
}
