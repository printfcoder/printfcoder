package main

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMergeTwoLists(t *testing.T) {
	Convey("Testing merge two lists", t, func() {
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
			5,
			&ListNode{
				7,
				&ListNode{
					8,
					nil,
				},
			},
		}
		x := mergeTwoLists(l1, l2)
		v := fmt.Sprintf("%d", x.val)
		for x.next != nil {
			v = v + fmt.Sprintf("%d", x.next.val)
			x.next = x.next.next
		}

		So(v, ShouldEqual, "345678")
	})
}
