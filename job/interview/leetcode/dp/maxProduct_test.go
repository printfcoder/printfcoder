package dp

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMaxProduct(t *testing.T) {
	Convey("Testing maxProduct", t, func() {
		bad := maxProduct([]int{2, 3, -2, 4})
		So(bad, ShouldEqual, 6)

		bad = maxProduct([]int{2, 3, -2})
		So(bad, ShouldEqual, 6)

		bad = maxProduct([]int{2, 3, 2})
		So(bad, ShouldEqual, 12)

		bad = maxProduct([]int{2, 3})
		So(bad, ShouldEqual, 6)

		bad = maxProduct([]int{3})
		So(bad, ShouldEqual, 3)

		bad = maxProduct([]int{-1, 3})
		So(bad, ShouldEqual, 3)

		bad = maxProduct([]int{-1, 3, -2})
		So(bad, ShouldEqual, 6)
	})
}
