package main

import "testing"

import . "github.com/smartystreets/goconvey/convey"

func TestMaxDepthWithDepthFirst(t *testing.T) {
	Convey("testing maxDepthWithDepthFirst", t, func() {
		v := maxDepthWithDepthFirst(
			&Node{
				Val:  3,
				Left: &Node{Val: 9},
				Right: &Node{Val: 20,
					Left:  &Node{Val: 15},
					Right: &Node{Val: 7},
				},
			},
		)
		So(v, ShouldEqual, 3)
	})
}

func TestMaxDepthWithBreadthFirst(t *testing.T) {
	Convey("testing maxDepthWithBreadthFirst", t, func() {
		v := maxDepthWithBreadthFirst(
			&Node{
				Val:  3,
				Left: &Node{Val: 9},
				Right: &Node{Val: 20,
					Left:  &Node{Val: 15},
					Right: &Node{Val: 7},
				},
			},
		)
		So(v, ShouldEqual, 3)
	})
}
