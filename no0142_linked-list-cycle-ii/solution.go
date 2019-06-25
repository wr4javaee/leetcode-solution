package no0142_linked_list_cycle_ii

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {

	if head == nil {
		return head
	}

	fast := head
	slow := head
	ptr := head

	for slow.Next != nil && slow.Next.Next != nil {
		fast = fast.Next
		slow = slow.Next.Next

		if fast == slow {
			// find cycle
			return findCycleEntrance(ptr, fast)
		}
	}

	return nil
}

func findCycleEntrance(ptr *ListNode, fast *ListNode) *ListNode {
	for ptr != fast {
		ptr = ptr.Next
		fast = fast.Next
	}
	return ptr
}
