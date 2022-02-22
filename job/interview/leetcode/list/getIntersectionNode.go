package main

/**
相交链表
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 用于殊途同归的办法
	// 如果有相交的点，那么两个人在这条路上走，走完了再从头走对方的路，总有一天能遇到
	// 证明：相交前A的路有a个节点，B的有b个节点；而相交后都是c个节点
	// 故而 a+c+b = b+c+a
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}
