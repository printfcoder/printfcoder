package dp

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLengthOfLIS(t *testing.T) {
	Convey("Testing lengthOfLIS", t, func() {
		x := lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})
		So(x, ShouldEqual, 4)

		x = lengthOfLIS([]int{0, 1, 0, 3, 2, 3})
		So(x, ShouldEqual, 4)
		x = lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7})
		So(x, ShouldEqual, 1)
	})
}
