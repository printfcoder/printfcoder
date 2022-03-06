package string

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFindTheDifference(t *testing.T) {
	Convey("testing isIsomorphic", t, func() {
		v := findTheDifference("a", "aa")
		So(v, ShouldEqual, 'a')
	})
}
