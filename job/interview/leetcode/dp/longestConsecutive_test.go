package dp

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLongestConsecutive(t *testing.T) {
	Convey("Testing longestConsecutive", t, func() {
		x := longestConsecutive([]int{100, 4, 200, 1, 3, 2})
		So(x, ShouldEqual, 4)

		x = longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
		So(x, ShouldEqual, 9)
	})
}
