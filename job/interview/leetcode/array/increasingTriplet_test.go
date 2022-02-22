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
	})
}
