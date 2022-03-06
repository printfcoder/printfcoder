package math

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAddDigits(t *testing.T) {
	Convey("testing addDigits", t, func() {
		v := addDigits(38)
		So(v, ShouldEqual, 2)

		v = addDigits(0)
		So(v, ShouldEqual, 0)

		v = addDigits(2)
		So(v, ShouldEqual, 2)

		v = addDigits(100)
		So(v, ShouldEqual, 1)

		v = addDigits(111)
		So(v, ShouldEqual, 3)
	})
}
