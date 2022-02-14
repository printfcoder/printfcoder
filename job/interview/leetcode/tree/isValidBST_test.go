package main

import "testing"

import . "github.com/smartystreets/goconvey/convey"

func TestIsValidBST(t *testing.T) {
	Convey("testing isValidBST", t, func() {
		v := isValidBST(&Node{
			Val:  3,
			Left: &Node{Val: 1},
			Right: &Node{Val: 20,
				Left:  &Node{Val: 15},
				Right: &Node{Val: 22},
			},
		})
		So(v, ShouldEqual, true)

		v = isValidBST(&Node{
			Val: 3,
		})
		So(v, ShouldEqual, true)

		v = isValidBST(&Node{
			Val:  3,
			Left: &Node{Val: 2},
		})
		So(v, ShouldEqual, true)

		v = isValidBST(&Node{
			Val:  3,
			Left: &Node{Val: 4},
			Right: &Node{Val: 20,
				Left:  &Node{Val: 15},
				Right: &Node{Val: 22},
			},
		})
		So(v, ShouldEqual, false)

		v = isValidBST(&Node{
			Val:   3,
			Left:  &Node{Val: 3},
			Right: &Node{Val: 3},
		})
		So(v, ShouldEqual, false)

		v = isValidBST(&Node{
			Val: 3,
			Right: &Node{Val: 20,
				Left:  &Node{Val: 15},
				Right: &Node{Val: 22},
			},
		})
		So(v, ShouldEqual, true)

		// [5,4,6,null,null,3,7]

		v = isValidBST(&Node{
			Val:  5,
			Left: &Node{Val: 4},
			Right: &Node{Val: 6,
				Left:  &Node{Val: 3},
				Right: &Node{Val: 7},
			},
		})
		So(v, ShouldEqual, false)
	})
}
