package array

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRemoveDuplicates(t *testing.T) {
	Convey("Testing remove duplicates", t, func() {
		l, output := removeDuplicates([]int{1, 2, 2, 4, 4, 8})
		So(l, ShouldEqual, 4)
		So(output, ShouldResemble, []int{1, 2, 4, 8})
	})
}
