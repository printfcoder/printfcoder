package string

import "testing"

import . "github.com/smartystreets/goconvey/convey"

func TestNumDifferentIntegers(t *testing.T) {
	Convey("testing numDifferentIntegers", t, func() {
		v := numDifferentIntegers("a123bc34d8ef34")
		So(v, ShouldEqual, 3)

		v = numDifferentIntegers("leet1234code234")
		So(v, ShouldEqual, 2)

		v = numDifferentIntegers("a1b01c001")
		So(v, ShouldEqual, 1)

		v = numDifferentIntegers("a")
		So(v, ShouldEqual, 0)

		v = numDifferentIntegers("123")
		So(v, ShouldEqual, 1)

		v = numDifferentIntegers("aa123aa0123")
		So(v, ShouldEqual, 1)
	})
}
