package main

import "testing"

import . "github.com/smartystreets/goconvey/convey"

func TestFirstUniqChar(t *testing.T) {
	Convey("testing firstUniqChar", t, func() {
		v := firstUniqChar("leetcode")
		So(v, ShouldEqual, 0)

		v = firstUniqChar("loveleetcode")
		So(v, ShouldEqual, 2)

		v = firstUniqChar("aabb")
		So(v, ShouldEqual, -1)

		v = firstUniqChar("bd")
		So(v, ShouldEqual, 0)

		v = firstUniqChar("dddccdbba")
		So(v, ShouldEqual, 8)

		v = firstUniqChar("a")
		So(v, ShouldEqual, 0)

		v = firstUniqChar("dda")
		So(v, ShouldEqual, 2)

		v = firstUniqChar("dd")
		So(v, ShouldEqual, -1)

		v = firstUniqChar("dad")
		So(v, ShouldEqual, 1)
	})
}
