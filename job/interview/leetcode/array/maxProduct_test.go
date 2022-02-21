package array

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMaxProduct(t *testing.T) {
	Convey("Testing maxProduct", t, func() {
		output := maxProduct([]int{2, 3, -2, 4})
		So(output, ShouldEqual, 6)

		output = maxProduct([]int{-2, 0, -1})
		So(output, ShouldEqual, 0)

		output = maxProduct([]int{0, 0, 0})
		So(output, ShouldEqual, 0)

		output = maxProduct([]int{-1, 1, -1})
		So(output, ShouldEqual, 1)

		output = maxProduct([]int{1, 1, 1})
		So(output, ShouldEqual, 1)

		output = maxProduct([]int{1, -6, -6, 1})
		So(output, ShouldEqual, 36)

		output = maxProduct([]int{1, -6, -6, 6, 1})
		So(output, ShouldEqual, 216)

		output = maxProduct([]int{1, -6, -6, 6, -6, -6, 1})
		So(output, ShouldEqual, 7776)

		output = maxProduct([]int{-4, -3})
		So(output, ShouldEqual, 12)
	})
}
