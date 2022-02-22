package dp

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRob(t *testing.T) {
	Convey("Testing rob", t, func() {
		bad := rob([]int{1, 2, 3, 1})
		So(bad, ShouldEqual, 4)

		bad = rob([]int{2, 7, 9, 3, 1})
		So(bad, ShouldEqual, 12)

		bad = rob([]int{1, 1, 1})
		So(bad, ShouldEqual, 2)

		bad = rob([]int{2, 3, 2})
		So(bad, ShouldEqual, 4)
	})
}
