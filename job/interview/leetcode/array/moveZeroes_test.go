package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMoveZeroes(t *testing.T) {
	Convey("Testing moveZeroes", t, func() {
		input := []int{1, 0, 0, 4, 4, 8}
		moveZeroesWithTwoPoints(input)
		So(input, ShouldResemble, []int{1, 4, 4, 8, 0, 0})
	})
}
