package bit

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAddBinary(t *testing.T) {
	Convey("Testing addBinary", t, func() {
		output := addBinary("11", "1")
		So(output, ShouldEqual, "100")
		output = addBinary("1010", "1011")
		So(output, ShouldEqual, "10101")
	})
}
