package list

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsPalindrome(t *testing.T) {
	Convey("Testing isPalindrome", t, func() {
		l1 := &ListNode{
			3,
			&ListNode{
				4,
				&ListNode{
					6,
					&ListNode{
						6,
						&ListNode{
							4,
							&ListNode{
								3,
								nil,
							},
						},
					},
				},
			}}
		x := isPalindrome(l1)
		So(x, ShouldEqual, true)

		l1 = &ListNode{
			3,
			&ListNode{
				4,
				&ListNode{
					6,
					&ListNode{
						5,
						&ListNode{
							6,
							&ListNode{
								4,
								&ListNode{
									3,
									nil,
								},
							},
						},
					},
				},
			}}
		x = isPalindrome(l1)
		So(x, ShouldEqual, true)
	})
}
