package dp

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSupperEgg(t *testing.T) {
	Convey("Testing supper egg", t, func() {
		x := superEggDrop(2, 6)
		So(x, ShouldEqual, 3)

		x = superEggDrop(3, 14)
		So(x, ShouldEqual, 4)
	})
}
