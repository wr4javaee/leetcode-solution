package no0206_reverse_linked_list

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	node := reverseList(head.Next)

	head.Next.Next = head
	head.Next = nil
	return node
}
