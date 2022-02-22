package math

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTrailingZeroes(t *testing.T) {
	Convey("testing trailingZeroes", t, func() {
		v := trailingZeroes(3)
		So(v, ShouldEqual, 0)

		v = trailingZeroes(5)
		So(v, ShouldEqual, 1)

		v = trailingZeroes(0)
		So(v, ShouldEqual, 0)

		v = trailingZeroes(30)
		So(v, ShouldEqual, 7)
	})
}
