package no0141_linked_list_cycle

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	// 边界值校验
	if head == nil || head.Next == nil {
		return false
	}

	// 只有一个节点的自循环
	if head.Next == head {
		return true
	}

	// 定义快慢指针
	headPtr := &ListNode{}
	headPtr.Next = head
	fast := headPtr
	slow := headPtr

	// 遍历链表
	// 快指针每次移动两次，慢指针每次移动一次
	// 对于有环的链表来说，快指针一定会在某个时间点追上慢指针
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if fast == nil {
			// 链表无环，fast已经指向尾节点
			return false
		}

		if fast == slow {
			return true
		}
	}

	// 链表无环，fast已经指向尾节点
	return false
}
