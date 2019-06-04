package no0019_remove_nth_node_from_end_of_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 异常校验
	if head == nil {
		return head
	}

	// 定义头指针
	headPtr := &ListNode{}
	headPtr.Next = head

	// 初始化first、second指针
	first := headPtr
	second := headPtr

	// first指针向后移动n+1个节点，使first指针与second指针保持n个间隔
	for n + 1 > 0 {
		first = first.Next
		n --
	}

	// 一步一步移动first、second指针，直到first指向nil
	for first != nil {
		first = first.Next
		second = second.Next
	}

	// 此时second指向的是应该删除节点的父节点
	// 如果second指向的是最后一个节点, 直接返回head
	if second.Next == nil {
		return head
	}

	second.Next = second.Next.Next
	return headPtr.Next
}
