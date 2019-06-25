package no0206_reverse_linked_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	// 异常校验
	if head == nil || head.Next == nil {
		return head
	}

	currNode := head
	prevNode := &ListNode{}
	prevNode = nil

	for currNode != nil {
		currNext := currNode.Next

		// 当前节点指向前驱节点
		currNode.Next = prevNode

		// 前驱指针指向当前节点
		prevNode = currNode

		// 当前节点指针指向下一个节点
		currNode = currNext
	}

	return prevNode
}
