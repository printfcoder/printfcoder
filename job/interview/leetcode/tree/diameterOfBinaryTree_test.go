package tree

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	Convey("testing diameterOfBinaryTree", t, func() {
		v := diameterOfBinaryTree(
			&TreeNode{
				Val: 1,
				Left: &TreeNode{Val: 2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{Val: 3},
			},
		)
		So(v, ShouldEqual, 3)

		v = diameterOfBinaryTree(
			&TreeNode{
				Val: 1,
			},
		)
		So(v, ShouldEqual, 0)

		v = diameterOfBinaryTree(
			&TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 3},
			},
		)
		So(v, ShouldEqual, 1)
	})
}
