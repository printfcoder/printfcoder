package main

import "testing"

import . "github.com/smartystreets/goconvey/convey"

func TestIsAnagram(t *testing.T) {
	Convey("testing isAnagram", t, func() {
		v := isAnagram("leetcode", "leetcode")
		So(v, ShouldEqual, true)

		v = isAnagram("leetcode", "codeleet")
		So(v, ShouldEqual, true)

		v = isAnagram("leetcode", "codeleet2")
		So(v, ShouldEqual, false)

		v = isAnagram("leetcode", "codeleet")
		So(v, ShouldEqual, true)

		v = isAnagram("anagram", "nagaram")
		So(v, ShouldEqual, true)

		v = isAnagram("aa", "bb")
		So(v, ShouldEqual, false)
	})
}
