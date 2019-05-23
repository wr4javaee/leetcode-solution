package merge_sort_list_bottom_to_top

//对链表进行归并排序（自底向上）。
//
//示例 1：
//
//输入: 4->2->1->3
//输出: 1->2->3->4
//示例 2：
//
//输入: -1->5->3->4->0
//输出: -1->0->3->4->5
// 递归算法的空间复杂度：递归深度N*每次递归所要的辅助空间，如果每次递归所需的辅助空间是常数，则递归的空间复杂度是 O(N).
type ListNode struct {
	Val int
	Next *ListNode
}

// 基于链表的归并排序（自底向上）, 由小到大排序
// 注意, 题目给的head为第一个节点的指针
//
// 注意, 链表的归并排序不需要额外的空间, 即空间复杂度为常量O(1)
// 利用归并的思想，递归地将当前链表分为两段，然后merge
// 分两段的方法是使用 fast-slow 法，用两个指针，一个每次走两步，一个走一步，直到快的走到了末尾
// 然后慢的所在位置就是中间位置，这样就分成了两段
// merge时，把两段头部节点值比较，用一个 p 指向较小的
// 且记录第一个节点，然后两段的头一步一步向后走，p也一直向后走，总是指向较小节点，
// 直至其中一个头为NULL，处理剩下的元素。最后返回记录的头即可。
func mergeSortList(head *ListNode) *ListNode {
	// 异常数据校验
	if head == nil || head.Next == nil {
		return head
	}

	// 创建头指针
	headPtr := &ListNode{ 0, nil }
	headPtr.Next = head

	// 11 -> 22 -> 44 -> 88 -> 1616
	slowNode := headPtr
	fastNode := headPtr
	// 初始化子数组大小为1
	for sz := 1; ; sz = sz * 2 {

	}



}

// 将链表进行归并
func merge(lowNode *ListNode, highNode *ListNode) *ListNode {
	// 定义头节点
	headPtr := &ListNode{ Next:nil }
	currNode := headPtr

	// 遍历左右节点
	for lowNode.Next != nil && highNode.Next != nil {
		if lowNode.Val < highNode.Val {
			currNode.Next = lowNode
			currNode = currNode.Next
			lowNode = lowNode.Next
			continue
		}

		currNode.Next = highNode
		currNode = currNode.Next
		highNode = highNode.Next
	}

	// 处理左节点或右节点的下一个节点为空的情况
	if lowNode.Next != nil {
		currNode.Next = lowNode
	}
	if highNode.Next != nil {
		currNode.Next = highNode
	}

	return headPtr.Next
}
