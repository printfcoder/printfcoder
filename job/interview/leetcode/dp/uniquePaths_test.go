package dp

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUniquePaths(t *testing.T) {
	Convey("Testing uniquePaths", t, func() {
		x := uniquePaths(7, 3)
		So(x, ShouldEqual, 28)

		x = uniquePaths(3, 3)
		So(x, ShouldEqual, 6)
	})
}
