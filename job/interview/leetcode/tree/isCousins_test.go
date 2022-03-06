package tree

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsCousins(t *testing.T) {
	Convey("testing isCousins", t, func() {
		v := isCousins(&TreeNode{
			Val: 1,
			Left: &TreeNode{Val: 2,
				Left: &TreeNode{Val: 4},
			},
			Right: &TreeNode{Val: 3},
		}, 4, 3)
		So(v, ShouldEqual, false)

		// 输入：root = [1,2,3,null,4,null,5], x = 5, y = 4
		v = isCousins(&TreeNode{
			Val: 1,
			Left: &TreeNode{Val: 2,
				Right: &TreeNode{Val: 4},
			},
			Right: &TreeNode{Val: 3,
				Right: &TreeNode{Val: 5},
			},
		}, 5, 4)
		So(v, ShouldEqual, true)

		// 输入：root = [1,2,3,null,4], x = 2, y = 3
		v = isCousins(&TreeNode{
			Val: 1,
			Left: &TreeNode{Val: 2,
				Right: &TreeNode{Val: 4},
			},
			Right: &TreeNode{Val: 3},
		}, 2, 3)
		So(v, ShouldEqual, false)
	})
}
