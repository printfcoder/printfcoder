package main

// SinglyLinkedListNode Definition for singly-linked list.
type SinglyLinkedListNode struct {
	Val  int
	Next *SinglyLinkedListNode
}

func deleteNodeInSinglyLinked(node *SinglyLinkedListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func deleteNodeInSinglyLinked2(node *SinglyLinkedListNode) {
	if node == nil {
		return
	}
	cur := node
	for cur.Next.Next != nil {
		cur.Val = cur.Next.Val
		cur = cur.Next
	}
	cur.Val = cur.Next.Val
	cur.Next = nil
}
