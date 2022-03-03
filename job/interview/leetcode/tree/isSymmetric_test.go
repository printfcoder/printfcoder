package tree

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsSymmetricWithRecursion(t *testing.T) {
	Convey("testing isSymmetricWithRecursion", t, func() {
		v := isSymmetricWithRecursion(
			&TreeNode{
				Val: 3,
				Left: &TreeNode{Val: 20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
				Right: &TreeNode{Val: 20,
					Right: &TreeNode{Val: 15},
					Left:  &TreeNode{Val: 7},
				},
			},
		)
		So(v, ShouldEqual, true)

		v = isSymmetricWithRecursion(
			&TreeNode{
				Val: 3,
				Left: &TreeNode{Val: 20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
				Right: &TreeNode{Val: 20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
		)
		So(v, ShouldEqual, false)
	})
}

func TestIsSymmetricWithIteration(t *testing.T) {
	Convey("testing isSymmetricWithIteration", t, func() {
		v := isSymmetricWithIteration(
			&TreeNode{
				Val: 3,
				Left: &TreeNode{Val: 20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
				Right: &TreeNode{Val: 20,
					Right: &TreeNode{Val: 15},
					Left:  &TreeNode{Val: 7},
				},
			},
		)
		So(v, ShouldEqual, true)

		v = isSymmetricWithRecursion(
			&TreeNode{
				Val: 3,
				Left: &TreeNode{Val: 20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
				Right: &TreeNode{Val: 20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
		)
		So(v, ShouldEqual, false)
	})
}
