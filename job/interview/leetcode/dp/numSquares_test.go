package dp

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNumSquares(t *testing.T) {
	Convey("Testing numSquares", t, func() {
		x := numSquares(1)
		So(x, ShouldEqual, 1)

		x = numSquares(2)
		So(x, ShouldEqual, 2)

		x = numSquares(4)
		So(x, ShouldEqual, 1)

		x = numSquares(13)
		So(x, ShouldEqual, 2)

		x = numSquares(12)
		So(x, ShouldEqual, 3)
	})
}
