package main

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAddTwoNums(t *testing.T) {
	Convey("Testing two nums", t, func() {
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
		v := ""
		for x.next != nil {
			v = v + fmt.Sprintf("%d", x.next.val)
			x.next = x.next.next
		}

		So(v, ShouldEqual, "1911")
	})
}
