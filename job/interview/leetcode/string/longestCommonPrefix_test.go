package string

import "testing"

import . "github.com/smartystreets/goconvey/convey"

func TestLongestCommonPrefixCrosswise(t *testing.T) {
	Convey("testing length of longest substring", t, func() {
		v := longestCommonPrefixCrosswise("pwwkew", "pwe", "pw")
		So(v, ShouldEqual, "pw")

		v = longestCommonPrefixCrosswise("pwwkew", "vpwe", "pw")
		So(v, ShouldEqual, "")
	})
}
