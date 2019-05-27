package no0109_convert_sorted_list_to_binary_search_tree

import "testing"

func Test(t *testing.T) {
	node1 := &ListNode{ Val: -10 }
	node2 := &ListNode{ Val: -3 }
	node3 := &ListNode{ Val: 0 }
	node4 := &ListNode{ Val: 5 }
	node5 := &ListNode{ Val: 9 }

	node4.Next = node5
	node3.Next = node4
	node2.Next = node3
	node1.Next = node2

	sortedListToBST(node1)
}
