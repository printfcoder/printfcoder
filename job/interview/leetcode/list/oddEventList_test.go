package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOddEvenList(t *testing.T) {
	// [1,2,3,4,5]
	Convey("Testing oddEvenList", t, func() {
		// [1,2,3,4,5]
		l1 := &ListNode{
			1,
			&ListNode{
				2,
				&ListNode{
					3,
					&ListNode{
						4,
						&ListNode{
							5,
							nil,
						},
					},
				},
			},
		}
		x := oddEvenList2(l1)
		So(x, ShouldResemble, &ListNode{
			1,
			&ListNode{
				3,
				&ListNode{
					5,
					&ListNode{
						2,
						&ListNode{
							4,
							nil,
						},
					},
				},
			},
		})
	})
}
