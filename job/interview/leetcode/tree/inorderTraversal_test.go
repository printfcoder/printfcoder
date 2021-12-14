package main

import "testing"

import . "github.com/smartystreets/goconvey/convey"

func TestInorderTraversalWithRecursion(t *testing.T) {
	Convey("testing inorderTraversalWithRecursion", t, func() {
		v := inorderTraversalWithRecursion(
			&Node{
				Val:  3,
				Left: &Node{Val: 9},
				Right: &Node{Val: 20,
					Left:  &Node{Val: 15},
					Right: &Node{Val: 7},
				},
			},
		)
		So(v, ShouldResemble, []int{9, 3, 15, 20, 7})
	})
}

func TestInorderTraversalWithIteration(t *testing.T) {
	Convey("testing inorderTraversalWithIteration", t, func() {
		v := inorderTraversalWithIteration(
			&Node{
				Val:  3,
				Left: &Node{Val: 9},
				Right: &Node{Val: 20,
					Left:  &Node{Val: 15},
					Right: &Node{Val: 7},
				},
			},
		)
		So(v, ShouldResemble, []int{9, 3, 15, 20, 7})
	})
}
