package no0024_swap_nodes_in_pairs

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	prev := &ListNode{}
	prev.Next = head
	res := head.Next
	for prev.Next != nil && prev.Next.Next != nil {
		a := prev.Next
		b := a.Next
		c := b.Next

		b.Next = a
		a.Next = c
		prev.Next = b
		prev = a
	}

	return res
}
