package dp

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLongestPalindrome(t *testing.T) {
	Convey("Testing merge two lists", t, func() {
		x := longestPalindrome("babad")
		So(x, ShouldEqual, "aba")

		x = longestPalindrome("cbbd")
		So(x, ShouldEqual, "bb")
	})
}
