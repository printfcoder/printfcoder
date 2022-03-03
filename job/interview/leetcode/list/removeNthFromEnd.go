package list

/**
给你一个链表，删除链表的倒数第n个结点，并且返回链表的头结点。

示例 1：

输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
示例 2：

输入：head = [1], n = 1
输出：[]
示例 3：

输入：head = [1,2], n = 1
输出：[1]

提示：

链表中结点的数目为 sz
1 <= sz <= 30
0 <= Node.Val <= 100
1 <= n <= sz
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	// 对于只有一个节点的情况
	if head.Next == nil && n == 1 {
		return nil
	}

	fast, del := head, head
	// 用于保存删除前的一个，因为可能删除是最后一个成员，此时del已经游到被删除者了，要通过前置成员把它删除
	var delPre *ListNode
	for i, j := 0, 0; fast != nil; j++ {
		// 到底了
		// 找到第n个，开始删除
		if fast.Next == nil && del != nil {
			// 删除
			if del.Next != nil {
				del.Val = del.Next.Val
				del.Next = del.Next.Next
			} else {
				// 说明删除的是最后一个
				// 此时del指针还在上一个
				delPre.Next = nil
			}
			break
		} else if j-i == n-1 {
			delPre = del
			del = del.Next
			i++
		}
		fast = fast.Next
	}

	return head
}
