package string

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWordBreak(t *testing.T) {
	Convey("testing wordBreak", t, func() {
		v := wordBreak("leetcode", []string{"leet", "code"})
		So(v, ShouldEqual, true)
		v = wordBreak("applepenapple", []string{"apple", "pen"})
		So(v, ShouldEqual, true)
		v = wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})
		So(v, ShouldEqual, false)
		v = wordBreak("aaaaaaa", []string{"aaaa", "aa"})
		So(v, ShouldEqual, false)
		v = wordBreak("cars", []string{"car", "ca", "rs"})
		So(v, ShouldEqual, true)
		v = wordBreak("catskicatcats", []string{"cats", "cat", "dog", "ski"})
		So(v, ShouldEqual, true)
	})
}
