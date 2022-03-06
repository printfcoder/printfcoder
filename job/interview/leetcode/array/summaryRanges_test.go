package array

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSummaryRanges(t *testing.T) {
	Convey("Testing summaryRanges", t, func() {
		x := summaryRanges([]int{0, 1, 2, 4, 5, 7})
		So(x, ShouldResemble, []string{"0->2", "4->5", "7"})

		x = summaryRanges([]int{0, 2, 3, 4, 6, 8, 9})
		So(x, ShouldResemble, []string{"0", "2->4", "6", "8->9"})

		x = summaryRanges([]int{})
		So(x, ShouldResemble, []string{})

		x = summaryRanges([]int{0})
		So(x, ShouldResemble, []string{"0"})
	})
}
