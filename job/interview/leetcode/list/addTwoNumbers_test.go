package list

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAddTwoNums(t *testing.T) {
	Convey("Testing two nums", t, func() {
		x := addTwoNumbers(arrayToList([]int{2, 4, 3}), arrayToList([]int{5, 6, 4}))
		So(x, ShouldResemble, arrayToList([]int{7, 0, 8}))
		x = addTwoNumbers(arrayToList([]int{9, 9, 9, 9, 9, 9, 9}), arrayToList([]int{9, 9, 9, 9}))
		So(x, ShouldResemble, arrayToList([]int{8, 9, 9, 9, 0, 0, 0, 1}))
		x = addTwoNumbers(arrayToList([]int{9}), arrayToList([]int{9}))
		So(x, ShouldResemble, arrayToList([]int{8, 1}))
		x = addTwoNumbers(arrayToList([]int{9}), nil)
		So(x, ShouldResemble, arrayToList([]int{9}))
		x = addTwoNumbers(arrayToList([]int{9}), arrayToList([]int{0}))
		So(x, ShouldResemble, arrayToList([]int{9}))
	})
}
