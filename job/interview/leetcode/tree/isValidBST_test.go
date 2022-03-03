package tree

import "testing"

import . "github.com/smartystreets/goconvey/convey"

func TestIsValidBST(t *testing.T) {
	Convey("testing isValidBST", t, func() {
		v := isValidBST(&TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 1},
			Right: &TreeNode{Val: 20,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 22},
			},
		})
		So(v, ShouldEqual, true)

		v = isValidBST(&TreeNode{
			Val: 3,
		})
		So(v, ShouldEqual, true)

		v = isValidBST(&TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 2},
		})
		So(v, ShouldEqual, true)

		v = isValidBST(&TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 4},
			Right: &TreeNode{Val: 20,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 22},
			},
		})
		So(v, ShouldEqual, false)

		v = isValidBST(&TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 3},
		})
		So(v, ShouldEqual, false)

		v = isValidBST(&TreeNode{
			Val: 3,
			Right: &TreeNode{Val: 20,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 22},
			},
		})
		So(v, ShouldEqual, true)

		// [5,4,6,null,null,3,7]

		v = isValidBST(&TreeNode{
			Val:  5,
			Left: &TreeNode{Val: 4},
			Right: &TreeNode{Val: 6,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 7},
			},
		})
		So(v, ShouldEqual, false)
	})
}
