package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSortedArrayToBST(t *testing.T) {
	Convey("testing sortedArrayToBST", t, func() {
		v := sortedArrayToBST([]int{1, 2, 3, 4, 5, 6, 7})
		So(v, ShouldResemble, &Node{
			Val: 4,
			Left: &Node{
				Val:   2,
				Left:  &Node{Val: 1},
				Right: &Node{Val: 3},
			},
			Right: &Node{
				Val:   6,
				Left:  &Node{Val: 5},
				Right: &Node{Val: 7},
			},
		})

		v = sortedArrayToBST([]int{1})
		So(v, ShouldResemble, &Node{
			Val: 1,
		})

		v = sortedArrayToBST([]int{1, 2})
		So(v, ShouldResemble, &Node{
			Val: 2,
			Left: &Node{
				Val: 1,
			},
		})
	})
}
