package no0876_middle_of_the_linked_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	// 异常数据校验
	if head == nil || head.Next == nil {
		return head
	}

	// 定义头指针
	headPtr := &ListNode{}
	headPtr.Next = head

	// 定义快慢指针
	slow := headPtr
	fast := headPtr

	// 采用快慢指针法
	// 慢指针每次移动一次, 快指针每次移动两次
	// 当快指针第二步移动到nil时, 说明链表为单数, 返回慢指针指向的节点
	// 当快指针第一步移动到nil时, 说明链表为双数, 返回慢指针指向的节点
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			// 快指针第一步移动到nil时, 说明链表为双数, 返回慢指针指向的下一个节点
			return slow
		}

		fast = fast.Next
	}

	// 当快指针第二步移动到nil时, 说明链表为单数, 返回慢指针指向的节点
	return slow
}
