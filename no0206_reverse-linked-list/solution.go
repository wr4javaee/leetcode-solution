package no0206_reverse_linked_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//func reverseList(head *ListNode) *ListNode {
//	// 异常校验
//	if head == nil || head.Next == nil {
//		return head
//	}
//
//	// 定义头指针
//	headPtr := &ListNode{}
//	headPtr.Next = head
//
//	// 遍历链表并原地反转
//	// 从第二个节点开始遍历
//	node := headPtr.Next.Next
//	for node != nil {
//		nextNode := node.Next
//		// 当前遍历的是第一个节点时, 需要将第一个节点的next置为空
//		if headPtr.Next.Next == node {
//			headPtr.Next.Next = nil
//		}
//		if nextNode == nil {
//			// node为最后一个节点（因前面的异常校验，这里不用再校验链表仅有一个节点的情况）
//			node.Next = headPtr.Next
//			headPtr.Next = node
//			break
//		}
//
//
//		node.Next = headPtr.Next
//		headPtr.Next = node
//		node = nextNode
//	}
//
//	return headPtr.Next
//}
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
