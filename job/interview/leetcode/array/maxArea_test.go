package array

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMaxArea(t *testing.T) {
	Convey("Testing maxArea", t, func() {
		bad := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
		So(bad, ShouldEqual, 49)
		bad = maxArea([]int{1, 1})
		So(bad, ShouldEqual, 1)
	})
}
