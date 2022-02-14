package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFirstBadVersion(t *testing.T) {
	Convey("Testing firstBadVersion", t, func() {
		bad := firstBadVersion(2126753390)
		So(bad, ShouldEqual, 1702766719)
	})
}
