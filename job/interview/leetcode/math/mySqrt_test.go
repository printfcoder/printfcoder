package math

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMySqrt(t *testing.T) {
	Convey("testing mySqrt", t, func() {
		v := mySqrt(3, 3)
		So(v, ShouldEqual, 3)

		v = mySqrt(5, 3)
		So(v, ShouldEqual, 1)
	})
}
