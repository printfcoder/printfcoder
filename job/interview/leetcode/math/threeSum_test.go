package math

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestThreeSum(t *testing.T) {
	Convey("testing threeSum", t, func() {
		v := threeSum([]int{-1, 0, 1, 2, -1, -4})
		So(v, ShouldResemble, [][]int{{-1, -1, 2}, {-1, 0, 1}})

		v = threeSum([]int{})
		So(v, ShouldResemble, [][]int{})

		v = threeSum([]int{0})
		So(v, ShouldResemble, [][]int{})
	})
}
