package no0147_insertion_sort_list

//
//对链表进行插入排序。
//
//
//插入排序的动画演示如上。从第一个元素开始，该链表可以被认为已经部分排序（用黑色表示）。
//每次迭代时，从输入数据中移除一个元素（用红色表示），并原地将其插入到已排好序的链表中。
//
//
//
//插入排序算法：
//
//插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
//每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
//重复直到所有输入数据插入完为止。
//
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
func insertionSortList(head *ListNode) *ListNode {
	// 异常数据校验
	if head == nil || head.Next == nil {
		return head
	}

	// 创建头指针
	headPtr := &ListNode{ 0, nil }
	headPtr.Next = head

	// 基于链表的插入排序, 从第一个node遍历到最后一个元素
	for currNode := head; currNode.Next != nil; {
		// 比较当前节点值与当前节点的下一个节点值
 		nextNode := currNode.Next
		// 若当前节点的下一个节点值小于当前节点
		// 将当前节点的下一个节点从链表中移出
		if nextNode.Val > currNode.Val {
			// 当前节点的下一个节点值比当前节点大, 进行下一次循环
			currNode = currNode.Next
			continue
		}

		// 将当前节点的下一个节点从链表中移出
		currNode.Next = nextNode.Next

		// 将移出的nextNode插入到合适的地方：
		// 1、若nextNode的值比头节点还小, 则将nextNode值为头节点并修改头节点指针
		// 2、否则从头节点开始迭代, 直到找到迭代节点值小于nextNode且迭代节点的下一个节点的值大于等于nextNode
		// 将nextNode插入到该位置
		node := headPtr.Next
		for true {
			if headPtr.Next.Val >= nextNode.Val {
				nextNode.Next = headPtr.Next
				headPtr.Next = nextNode
				break
			}

			if node.Val < nextNode.Val && node.Next.Val >= nextNode.Val {
				nextNode.Next = node.Next
				node.Next = nextNode
				break
			}

			node = node.Next
		}
	}

	return headPtr.Next
}


