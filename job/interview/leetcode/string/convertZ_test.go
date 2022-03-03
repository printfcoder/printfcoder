package string

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConvertZ(t *testing.T) {
	Convey("testing convertZ", t, func() {
		v := convertZ("PAYPALISHIRING", 3)
		So(v, ShouldEqual, "PAHNAPLSIIGYIR")

		v = convertZ("PAYPALISHIRING", 4)
		So(v, ShouldEqual, "PINALSIGYAHRPI")

		v = convertZ("PA", 4)
		So(v, ShouldEqual, "PA")

		v = convertZ("PA", 1)
		So(v, ShouldEqual, "PA")
	})
}
