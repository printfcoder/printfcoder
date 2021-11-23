package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTwoSumWithMap(t *testing.T) {
	Convey("Testing two sum", t, func() {
		l1 := &ListNode{
			3,
			&ListNode{
				4,
				&ListNode{
					6,
					nil,
				},
			},
		}
		l2 := &ListNode{
			8,
			&ListNode{
				4,
				&ListNode{
					5,
					nil,
				},
			},
		}
		x := addTwoNumbers(l1, l2)
		So(x, ShouldEqual, 1191)
	})
}
