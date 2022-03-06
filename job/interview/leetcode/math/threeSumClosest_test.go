package math

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestThreeSumClosest(t *testing.T) {
	Convey("testing threeSumClosest", t, func() {
		v := threeSumClosest([]int{-1, 2, 1, -4}, 1)
		So(v, ShouldResemble, 2)

		v = threeSumClosest([]int{0, 0, 0}, 1)
		So(v, ShouldResemble, 0)
	})
}
