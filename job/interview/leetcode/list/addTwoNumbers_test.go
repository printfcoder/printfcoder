package list

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
		for x.Next != nil {
			v = v + fmt.Sprintf("%d", x.Next.Val)
			x.Next = x.Next.Next
		}

		So(v, ShouldEqual, "1911")
	})
}
