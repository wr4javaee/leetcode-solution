package no0002_add_two_numbers

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 校验空链表
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	// 初始化返回链表
	resNode := new(ListNode)

	// 初始化两个新指针
	p1 := l1
	p2 := l2
	p3 := resNode

	// 按序遍历链表
	// 设置默认进位为0
	carry := 0
	for p1 != nil || p2 != nil || carry > 0 {
		p1Val := 0
		if p1 != nil {
			p1Val = p1.Val
		}
		p2Val := 0
		if p2 != nil {
			p2Val = p2.Val
		}
		p3Val := p1Val + p2Val + carry
		// 考虑进位问题
		if p3Val >= 10 {
			// 需要进位
			p3Val -= 10
			carry = 1
		} else {
			// 不用进位, 恢复进位flag
			carry = 0
		}

		// 链接返回节点
		tmpNode := new(ListNode)
		tmpNode.Val = p3Val
		p3.Next = tmpNode
		p3 = p3.Next

		// 遍历入参节点
		if p1 != nil {
			p1 = p1.Next
		}
		if p2 != nil {
			p2 = p2.Next
		}
	}

	return resNode.Next
}


