package main

type ListNode struct {
	val  int
	next *ListNode
}

// 给你两个非空的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
func addTwoNumbers(l1 *ListNode, l2 *ListNode) (output *ListNode) {
	// 忽略边界检查

	// 进位，用于记录上个节点超过10的部分
	carry := 0
	// 记录最后一次余数
	head := new(ListNode)
	cur := head
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
		cur.next = &ListNode{num % 10, nil}
		// 计算下一个
		cur = cur.next
	}

	return head
}
