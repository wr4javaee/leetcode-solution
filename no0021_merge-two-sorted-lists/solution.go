package no0021_merge_two_sorted_lists

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 初始化两个头指针
	headL1 := &ListNode{}
	headL2 := &ListNode{}
	headL1.Next = l1
	headL2.Next = l2

	// 初始化两个遍历指针
	// 分别指向当前节点的前一个节点
	l1Prev := headL1
	l2Prev := headL2

	// 遍历两个链表
	// 将l1合并至l2
	for l1Prev.Next != nil || l2Prev.Next != nil {
		// l1已遍历完, 直接返回l2
		if l1Prev.Next == nil {
			return headL2.Next
		}

		// l2遍历完, 将l1拼接在l2后并返回l2
		if l2Prev.Next == nil {
			l2Prev.Next = l1Prev.Next
			return headL2.Next
		}

		// 比较l1与l2下一个节点值
		// 若l1.Next值较小或相等, 则l1.Next插入到l2.Next前, 移动l1并进入下一次循环
		// 否则移动l2, 进入下一次循环
		if l1Prev.Next.Val <= l2Prev.Next.Val {
			// l1.Next值较小或相等
			// 将l1.Next从l1中删除
			l1PrevNext := l1Prev.Next
			l1Prev.Next = l1Prev.Next.Next

			// 从l1中删除的节点插入到l2.Next前
			l2PrevNext := l2Prev.Next
			l2Prev.Next = l1PrevNext
			l1PrevNext.Next = l2PrevNext

			// 重置l2指针, 向下一个节点移动一次
			l2Prev = l2Prev.Next
			continue
		}

		// l1.Next.Val > l2.Next.Val
		// 移动l2指针, 向下一个节点移动一次
		l2Prev = l2Prev.Next
	}

	return headL2.Next
}
