package no0024_swap_nodes_in_pairs

func swapPairs4Review(head *ListNode) *ListNode {
	// validate
	if head == nil || head.Next == nil {
		return head
	}

	prev := &ListNode{}
	prev.Next = head
	headPtr := head.Next

	for prev.Next != nil && prev.Next.Next != nil {
		a := prev.Next
		b := a.Next
		c := b.Next

		b.Next = a
		a.Next = c
		prev.Next = b
		prev = a
	}

	return headPtr
}
