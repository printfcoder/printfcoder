package ut

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAdd(t *testing.T) {
	ret := Add(1, 3)
	if ret != 4 {
		t.Fatalf("the sum should be %d but got %d", 4, ret)
	}
}

func TestAddWithConvey(t *testing.T) {
	Convey("testing Add", t, func() {
		ret := Add(1, 3)
		Convey("The value should be 4", func() {
			So(ret, ShouldEqual, ret)
		})
	})
}
