package dp

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMaxProfit(t *testing.T) {
	Convey("Testing maxProfit", t, func() {
		bad := maxProfit([]int{7, 1, 5, 3, 6, 4})
		So(bad, ShouldEqual, 5)

		bad = maxProfit([]int{7})
		So(bad, ShouldEqual, 0)
	})
}
