package selection_sort_list

//
//对链表进行选择排序。
//
//示例 1：
//
//输入: 4->2->1->3
//输出: 1->2->3->4
//示例 2：
//
//输入: -1->5->3->4->0
//输出: -1->0->3->4->5
type ListNode struct {
	Val int
	Next *ListNode
}

// 基于链表的插入排序, 由小到大排序
// 注意, 题目给的head为第一个节点的指针
func selectionSortList(head *ListNode) *ListNode {
	// 异常数据校验
	if head == nil || head.Next == nil {
		return head
	}

	// 创建头指针
	headPtr := &ListNode{ 0, nil }
	headPtr.Next = head

	// 从第一个节点开始遍历至倒数第二个节点
	// 因交换节点需要节点的前一个节点, 因此从头节点遍历
	for node:= headPtr; node.Next != nil && node.Next.Next != nil; node = node.Next {
		// 当前节点
		currNode := node.Next
		// 最小值节点, 默认为当前节点
		minValNode := currNode
		minValLeftNode := node

		// 从当前节点的下一个节点开始遍历至最后一个节点
		// 因节点交换需要左节点, 因此以左节点作为遍历条件
		change := false
		for jNode := currNode; jNode.Next != nil; jNode = jNode.Next {
			if jNode.Next.Val < minValNode.Val {
				// 找到较小值, 进行节点交换
				change = true
				minValNode = jNode.Next
				minValLeftNode = jNode
			}
		}

		// 循环结束后, 若找到较小节点则进行节点交换
		if change {
			// 若两个节点相邻（避免死循环）
			if currNode.Next == minValNode {
				minValRightNode := minValNode.Next
				minValNode.Next = currNode
				node.Next = minValNode
				currNode.Next = minValRightNode
				continue
			}

			// 若两个节点不相邻
			minValRightNode := minValNode.Next
			minValNode.Next = currNode.Next
			node.Next = minValNode
			currNode.Next = minValRightNode
			minValLeftNode.Next = currNode
		}
	}

	return headPtr.Next
}
