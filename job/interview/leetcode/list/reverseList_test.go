package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReverseList(t *testing.T) {
	Convey("Testing reverseList", t, func() {

		l1 := &ListNode{3,
			&ListNode{
				2,
				&ListNode{
					1,
					nil,
				},
			},
		}
		x := reverseList(l1)
		So(x, ShouldResemble, &ListNode{1, &ListNode{
			2,
			&ListNode{
				3,
				nil,
			},
		}})
	})
}
