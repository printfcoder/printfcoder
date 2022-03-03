package tree

import "testing"

import . "github.com/smartystreets/goconvey/convey"

func TestMaxDepthWithDepthFirst(t *testing.T) {
	Convey("testing maxDepthWithDepthFirst", t, func() {
		v := maxDepthWithDepthFirst(
			&TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{Val: 20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
		)
		So(v, ShouldEqual, 3)
	})
}

func TestMaxDepthWithBreadthFirst(t *testing.T) {
	Convey("testing maxDepthWithBreadthFirst", t, func() {
		v := maxDepthWithBreadthFirst(
			&TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{Val: 20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
		)
		So(v, ShouldEqual, 3)
	})
}
