package array

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIncreasingTriplet(t *testing.T) {
	Convey("Testing increasingTriplet", t, func() {
		output := increasingTriplet([]int{1, 2, 3, 4, 5})
		So(output, ShouldEqual, true)
		output = increasingTriplet([]int{5, 4, 3, 2, 1})
		So(output, ShouldEqual, false)
		output = increasingTriplet([]int{5, 4})
		So(output, ShouldEqual, false)
		output = increasingTriplet([]int{1, 4, 3, 2, 5})
		So(output, ShouldEqual, true)
		output = increasingTriplet([]int{3, 4, 3, 2, 5})
		So(output, ShouldEqual, true)
		output = increasingTriplet([]int{0, 4, 2, 1, 0, -1, -3})
		So(output, ShouldEqual, false)
		output = increasingTriplet([]int{20, 100, 10, 12, 5, 13})
		So(output, ShouldEqual, true)
		output = increasingTriplet([]int{1, 5, 0, 4, 1, 3})
		So(output, ShouldEqual, true)
	})
}
