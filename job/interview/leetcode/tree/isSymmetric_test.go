package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsSymmetricWithRecursion(t *testing.T) {
	Convey("testing isSymmetricWithRecursion", t, func() {
		v := isSymmetricWithRecursion(
			&Node{
				Val: 3,
				Left: &Node{Val: 20,
					Left:  &Node{Val: 15},
					Right: &Node{Val: 7},
				},
				Right: &Node{Val: 20,
					Right: &Node{Val: 15},
					Left:  &Node{Val: 7},
				},
			},
		)
		So(v, ShouldEqual, true)

		v = isSymmetricWithRecursion(
			&Node{
				Val: 3,
				Left: &Node{Val: 20,
					Left:  &Node{Val: 15},
					Right: &Node{Val: 7},
				},
				Right: &Node{Val: 20,
					Left: &Node{Val: 15},
					Right:  &Node{Val: 7},
				},
			},
		)
		So(v, ShouldEqual, false)
	})
}

func TestIsSymmetricWithIteration(t *testing.T) {
	Convey("testing isSymmetricWithIteration", t, func() {
		v := isSymmetricWithIteration(
			&Node{
				Val: 3,
				Left: &Node{Val: 20,
					Left:  &Node{Val: 15},
					Right: &Node{Val: 7},
				},
				Right: &Node{Val: 20,
					Right: &Node{Val: 15},
					Left:  &Node{Val: 7},
				},
			},
		)
		So(v, ShouldEqual, true)

		v = isSymmetricWithRecursion(
			&Node{
				Val: 3,
				Left: &Node{Val: 20,
					Left:  &Node{Val: 15},
					Right: &Node{Val: 7},
				},
				Right: &Node{Val: 20,
					Left: &Node{Val: 15},
					Right:  &Node{Val: 7},
				},
			},
		)
		So(v, ShouldEqual, false)
	})
}