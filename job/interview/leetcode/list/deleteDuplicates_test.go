package list

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDeleteDuplicates(t *testing.T) {
	Convey("Testing deleteDuplicates", t, func() {
		x := deleteDuplicates(arrayToList([]int{1, 2, 3, 3, 4, 4, 5}))
		So(x, ShouldResemble, arrayToList([]int{1, 2, 5}))
		x = deleteDuplicates(arrayToList([]int{1, 1, 1, 2, 3}))
		So(x, ShouldResemble, arrayToList([]int{2, 3}))

	})
}

func TestDeleteDuplicatesSaveOne(t *testing.T) {
	Convey("Testing deleteDuplicatesSaveOne", t, func() {
		x := deleteDuplicatesSaveOne(arrayToList([]int{1, 2, 3, 3, 4, 4, 5}))
		So(x, ShouldResemble, arrayToList([]int{1, 2, 3, 4, 5}))
		x = deleteDuplicatesSaveOne(arrayToList([]int{1, 1, 1, 2, 3}))
		So(x, ShouldResemble, arrayToList([]int{1, 2, 3}))
		x = deleteDuplicatesSaveOne(arrayToList([]int{1}))
		So(x, ShouldResemble, arrayToList([]int{1}))
		x = deleteDuplicatesSaveOne(nil)
		So(x, ShouldEqual, nil)
		x = deleteDuplicatesSaveOne(arrayToList([]int{1, 1, 1, 1, 1}))
		So(x, ShouldResemble, arrayToList([]int{1}))

		x = deleteDuplicatesSaveOne(arrayToList([]int{1, 1, 1, 2, 2}))
		So(x, ShouldResemble, arrayToList([]int{1, 2}))
	})
}

func arrayToList(arr []int) *ListNode {
	ret := &ListNode{}
	cur := ret
	for i, v := range arr {
		cur.Val = v
		if i != len(arr)-1 {
			cur.Next = &ListNode{}
			cur = cur.Next
		}
	}

	return ret
}
