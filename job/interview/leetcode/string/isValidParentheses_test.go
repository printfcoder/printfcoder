package string

import "testing"

import . "github.com/smartystreets/goconvey/convey"

func TestIsValidParentheses(t *testing.T) {
	Convey("testing IsValidParentheses", t, func() {
		v := IsValidParentheses("{}[]()")
		So(v, ShouldEqual, true)
		v = IsValidParentheses("{[()]}")
		So(v, ShouldEqual, true)
		v = IsValidParentheses("{[()]}}")
		So(v, ShouldEqual, false)
		v = IsValidParentheses("{[}]")
		So(v, ShouldEqual, false)
	})
}
