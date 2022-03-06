package bit

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCountNumbersWithUniqueDigits(t *testing.T) {
	Convey("Testing countNumbersWithUniqueDigits", t, func() {
		output := countNumbersWithUniqueDigits(2)
		So(output, ShouldEqual, 91)
		output = countNumbersWithUniqueDigits(3)
		So(output, ShouldEqual, 739)
		output = countNumbersWithUniqueDigits(4)
		So(output, ShouldEqual, 5275)
		output = countNumbersWithUniqueDigits(0)
		So(output, ShouldEqual, 1)
	})
}
