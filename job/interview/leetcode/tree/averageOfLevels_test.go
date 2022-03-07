package tree

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAverageOfLevels(t *testing.T) {
	Convey("testing averageOfLevels", t, func() {
		//
		// 输入：root = [3,9,20,null,null,15,7]
		//输出：[3.00000,14.50000,11.00000]
		v := averageOfLevels(
			&TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{Val: 20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
		)
		So(v, ShouldResemble, []float64{3.00000, 14.50000, 11.00000})

		//输入：root = [3,9,20,15,7]
		//输出：[3.00000,14.50000,11.00000]
		v = averageOfLevels(
			&TreeNode{
				Val: 3,
				Left: &TreeNode{Val: 9,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
				Right: &TreeNode{Val: 20},
			},
		)
		So(v, ShouldResemble, []float64{3.00000, 14.50000, 11.00000})

		v = averageOfLevels(
			&TreeNode{
				Val: 1,
			},
		)
		So(v, ShouldResemble, []float64{1.0})
	})
}
