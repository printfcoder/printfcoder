package math

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsPerfectSquare(t *testing.T) {
	Convey("testing isPerfectSquare", t, func() {
		v := isPerfectSquare(19)
		So(v, ShouldEqual, false)
		v = isPerfectSquare(4)
		So(v, ShouldEqual, true)
		v = isPerfectSquare(16)
		So(v, ShouldEqual, true)
		v = isPerfectSquare(18)
		So(v, ShouldEqual, false)
		v = isPerfectSquare(9)
		So(v, ShouldEqual, true)
		v = isPerfectSquare(1048576)
		So(v, ShouldEqual, true)
		v = isPerfectSquare(1048575)
		So(v, ShouldEqual, false)
		v = isPerfectSquare(2147483647)
		So(v, ShouldEqual, false)
	})
}
