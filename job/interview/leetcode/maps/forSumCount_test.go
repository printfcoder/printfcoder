package maps

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFourSumCount(t *testing.T) {
	Convey("Testing fourSumCount", t, func() {
		output := fourSumCount(
			[]int{5, 5, 5, 5, 1},
			[]int{5, 5, 5, 5, 1},
			[]int{5, 4, 3, 2, 1},
			[]int{5, 4, 3, 2, 1})

		So(output, ShouldEqual, 0)

		output = fourSumCount(
			[]int{5, 5},
			[]int{5, 5},
			[]int{-5, 4},
			[]int{-5, 4})

		So(output, ShouldEqual, 4)
	})
}
