package no0092_reverse_linked_list_ii

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {

	// validate
	if head == nil || head.Next == nil {
		return head
	}

	prev := &ListNode{}
	prev.Next = head
	mPre := prev

	// find m pre node
	for i := 0; i < m - 1; i ++ {
		mPre = mPre.Next
	}

	nNode := &ListNode{}
	nNode = nil
	mNode := mPre.Next
	cur := mPre.Next

	// reverse from n to m
	for i := 0; i < n - m + 1; i ++ {
		curNext := cur.Next
		cur.Next = nNode
		nNode = cur
		cur = curNext
	}

	mPre.Next = nNode
	mNode.Next = cur

	return prev.Next
}
