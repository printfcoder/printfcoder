package array

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMissingNumber(t *testing.T) {
	Convey("Testing missingNumber", t, func() {
		output := missingNumber([]int{1, 0, 3})
		So(output, ShouldResemble, 2)

		output = missingNumberXOR([]int{1, 0, 3})
		So(output, ShouldResemble, 2)
	})
}
