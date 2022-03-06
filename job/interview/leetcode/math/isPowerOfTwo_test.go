package math

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsPowerOfTwo(t *testing.T) {
	Convey("testing isPowerOfTwo", t, func() {
		v := isPowerOfTwo(19)
		So(v, ShouldEqual, false)
		v = isPowerOfTwo(2)
		So(v, ShouldEqual, true)
		v = isPowerOfTwo(1)
		So(v, ShouldEqual, true)
		v = isPowerOfTwo(128)
		So(v, ShouldEqual, true)
		v = isPowerOfTwo(14)
		So(v, ShouldEqual, false)
	})
}
