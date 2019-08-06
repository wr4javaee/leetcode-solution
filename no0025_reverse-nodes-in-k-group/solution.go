package no0025_reverse_nodes_in_k_group

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	// validate
	if head == nil || head.Next == nil {
		return head
	}

	preNode := &ListNode{}
	preNode.Next = head
	tmpNode := preNode
	currNode := preNode

	listLen := 0
	for currNode != nil {
		currNode = currNode.Next
		listLen ++
	}


	for listLen > k {
		currNode = preNode.Next
		for i := 1; i < k; i ++ {
			t := currNode.Next
			currNode.Next = t.Next
			t.Next = preNode.Next
			preNode.Next = t
		}
		preNode = currNode
		listLen = listLen - k
	}

	return tmpNode.Next
}
