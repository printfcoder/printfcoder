package list

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRemoveElements(t *testing.T) {
	// [1,2,3,4,5]
	Convey("Testing removeElements", t, func() {
		x := removeElements(arrayToList([]int{1, 2, 6, 3, 4, 5, 6}), 6)
		So(x, ShouldResemble, arrayToList([]int{1, 2, 3, 4, 5}))
		x = removeElements(arrayToList([]int{}), 1)
		So(x, ShouldResemble, arrayToList([]int{}))
		x = removeElements(arrayToList([]int{7, 7, 7, 7}), 7)
		So(x, ShouldResemble, arrayToList([]int{}))
	})
}
