package no0025_reverse_nodes_in_k_group

func reverseKGroup4Review(head *ListNode, k int) *ListNode {

	// validate
	if head == nil || head.Next == nil {
		return head
	}

	// init sentinel
	prev := &ListNode{}
	prev.Next = head
	tmp := prev
	cur := prev

	// get length from list
	listLen := 0
	for cur != nil {
		cur = cur.Next
		listLen ++
	}

	// reverse
	for listLen > k {
		cur = prev.Next

		// k group
		for i := 1; i < k; i ++ {
			t := cur.Next
			cur.Next = t.Next
			t.Next = prev.Next
			prev.Next = t
		}

		prev = cur
		listLen -= k
	}

	return tmp.Next

}
