package tree

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSortedArrayToBST(t *testing.T) {
	Convey("testing sortedArrayToBST", t, func() {
		v := sortedArrayToBST([]int{1, 2, 3, 4, 5, 6, 7})
		So(v, ShouldResemble, &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			Right: &TreeNode{
				Val:   6,
				Left:  &TreeNode{Val: 5},
				Right: &TreeNode{Val: 7},
			},
		})

		v = sortedArrayToBST([]int{1})
		So(v, ShouldResemble, &TreeNode{
			Val: 1,
		})

		v = sortedArrayToBST([]int{1, 2})
		So(v, ShouldResemble, &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
		})
	})
}
