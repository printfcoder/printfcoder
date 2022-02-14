package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRemoveNthFromEnd(t *testing.T) {
	Convey("Testing removeNthFromEnd", t, func() {
		l1 := &ListNode{3, nil}
		x := removeNthFromEnd(l1, 1)
		So(x, ShouldEqual, nil)

		l1 = &ListNode{3,
			&ListNode{
				2,
				&ListNode{
					1,
					nil,
				},
			},
		}
		x = removeNthFromEnd(l1, 2)
		So(x, ShouldResemble, &ListNode{3, &ListNode{
			1,
			nil,
		}})

		l1 = &ListNode{3,
			&ListNode{
				2,
				nil,
			},
		}
		x = removeNthFromEnd(l1, 2)
		So(x, ShouldResemble, &ListNode{2, nil})

		l1 = &ListNode{3,
			&ListNode{
				2,
				nil,
			},
		}
		x = removeNthFromEnd(l1, 1)
		So(x, ShouldResemble, &ListNode{3, nil})
	})
}
