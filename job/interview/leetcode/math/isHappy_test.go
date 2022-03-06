package math

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsHappy(t *testing.T) {
	Convey("testing isHappy", t, func() {
		v := isHappy(19)
		So(v, ShouldEqual, true)

		v = isHappy(2)
		So(v, ShouldEqual, false)

		v = isHappy(1)
		So(v, ShouldEqual, true)
	})
}
