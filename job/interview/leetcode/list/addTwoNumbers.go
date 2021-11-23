package main

type ListNode struct {
	val  int
	next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (output int) {
	// 忽略边界检查

	// 进位，用于记录上个节点超过10的部分
	carry := 0
	// 记录最后一次余数
	mod := 0
	// 链表不空则继续加，进位不空也加
	for l1 != nil || l2 != nil || carry != 0 {
		// 取两个链表的左边节点，取完值后需要游到下一节点
		n1, n2 := 0, 0
		if l1 != nil {
			n1, l1 = l1.val, l1.next
		}

		if l2 != nil {
			n2, l2 = l2.val, l2.next
		}

		// 同位相加，带上上次进位
		num := n1 + n2 + carry
		// 记录当前进位
		carry = num / 10
		// 记录余数
		mod = num % 10
	}

	// 	return mod + carry
}
