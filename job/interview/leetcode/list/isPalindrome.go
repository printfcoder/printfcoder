package list

func isPalindrome(head *ListNode) bool {
	cur := head
	// 找到mid节点
	mid := 0
	for fast := 0; cur != nil; fast++ {
		// 到头了
		if cur.Next == nil {
			mid = fast/2 + 1
		}

		cur = cur.Next
	}
	midNode := head
	for ; mid > 0; mid-- {
		midNode = midNode.Next
	}

	// 逆转midNode
	var reverseNode *ListNode
	for midNode != nil {
		next := midNode.Next
		midNode.Next = reverseNode
		reverseNode = midNode
		midNode = next
	}

	curFromHead := head
	for reverseNode != nil {
		if reverseNode.Val != curFromHead.Val {
			return false
		}

		curFromHead = curFromHead.Next
		reverseNode = reverseNode.Next
	}

	return true
}
