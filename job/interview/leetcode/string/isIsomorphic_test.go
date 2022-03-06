package string

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsIsomorphic(t *testing.T) {
	Convey("testing isIsomorphic", t, func() {
		v := isIsomorphic("egg", "add")
		So(v, ShouldEqual, true)
		v = isIsomorphic("foo", "bar")
		So(v, ShouldEqual, false)
		v = isIsomorphic("paper", "title")
		So(v, ShouldEqual, true)
		v = isIsomorphic("badc", "baba")
		So(v, ShouldEqual, false)
	})
}
