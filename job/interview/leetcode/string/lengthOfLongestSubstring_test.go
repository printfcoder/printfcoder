package string

import "testing"

import . "github.com/smartystreets/goconvey/convey"

func TestLengthOfLongestSubstring(t *testing.T) {
	Convey("testing length of longest substring", t, func() {
		v := lengthOfLongestSubstring("pwwkew")
		So(v, ShouldEqual, "kew")
	})
}
