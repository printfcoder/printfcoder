package string

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLengthOfLastWord(t *testing.T) {
	Convey("testing lengthOfLastWord", t, func() {
		v := lengthOfLastWord("pwwkew")
		So(v, ShouldEqual, 6)

		v = lengthOfLastWord("Hello World")
		So(v, ShouldEqual, 5)

		v = lengthOfLastWord("   fly me   to   the moon  ")
		So(v, ShouldEqual, 4)

		v = lengthOfLastWord("luffy is still joyboy")
		So(v, ShouldEqual, 6)
	})
}
