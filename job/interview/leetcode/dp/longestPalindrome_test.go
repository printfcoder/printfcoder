package dp

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLongestPalindrome(t *testing.T) {
	Convey("Testing longestPalindrome", t, func() {
		x := longestPalindrome("babad")
		So(x, ShouldEqual, "bab")

		x = longestPalindrome("cbbd")
		So(x, ShouldEqual, "bb")
	})
}
