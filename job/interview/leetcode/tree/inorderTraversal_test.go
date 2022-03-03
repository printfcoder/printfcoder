package tree

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInorderTraversalWithRecursion(t *testing.T) {
	Convey("testing inorderTraversalWithRecursion", t, func() {
		v := inorderTraversalWithRecursion(
			&TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{Val: 20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
		)
		So(v, ShouldResemble, []int{9, 3, 15, 20, 7})
	})
}

func TestInorderTraversalWithIteration(t *testing.T) {
	Convey("testing inorderTraversalWithIteration", t, func() {
		v := inorderTraversalWithIteration(
			&TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{Val: 20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
		)
		So(v, ShouldResemble, []int{9, 3, 15, 20, 7})
	})
}
