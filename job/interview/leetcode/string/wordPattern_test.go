package string

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWordPattern(t *testing.T) {
	Convey("testing wordPattern", t, func() {
		v := wordPattern("abba", "dog cat cat dog")
		So(v, ShouldEqual, true)
		v = wordPattern("abba", "dog cat cat fish")
		So(v, ShouldEqual, false)
		v = wordPattern("aaaa", "dog cat cat dog")
		So(v, ShouldEqual, false)

		v = wordPattern("aaa", "aa aa aa aa")
		So(v, ShouldEqual, false)
	})
}
