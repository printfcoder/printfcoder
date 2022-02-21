package array

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMergeWithTwoPointsFromHead(t *testing.T) {
	Convey("Testing MergeWithTwoPointsFromHead", t, func() {
		sorted := mergeWithTwoPointsFromHead([]int{1, 2, 2, 4, 4, 8}, []int{3, 5, 7, 9, 10})
		So(sorted, ShouldResemble, []int{1, 2, 2, 3, 4, 4, 5, 7, 8, 9, 10})
	})
}

func TestMergeWithTowPointsFromTail(t *testing.T) {
	Convey("Testing MergeWithTowPointsFromTail", t, func() {
		sorted := mergeWithTowPointsFromTail([]int{1, 2, 2, 4, 4, 8}, []int{3, 5, 7, 9, 10})
		So(sorted, ShouldResemble, []int{1, 2, 2, 3, 4, 4, 5, 7, 8, 9, 10})
	})
}
