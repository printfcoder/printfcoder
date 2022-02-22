package string

import "testing"

import . "github.com/smartystreets/goconvey/convey"

func TestIsValidParentheses(t *testing.T) {
	Convey("testing IsValidParentheses", t, func() {
		v := isValidParentheses("{}[]()")
		So(v, ShouldEqual, true)
		v = isValidParentheses("{[()]}")
		So(v, ShouldEqual, true)
		v = isValidParentheses("{[()]}}")
		So(v, ShouldEqual, false)
		v = isValidParentheses("{[}]")
		So(v, ShouldEqual, false)
	})
}
