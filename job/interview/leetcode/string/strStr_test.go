package main

import "testing"

import . "github.com/smartystreets/goconvey/convey"

func TestStrStr(t *testing.T) {
	Convey("testing strStr", t, func() {
		v := strStr("abadifuiopd", "ifuiopd")
		So(v, ShouldEqual, 4)
	})
}