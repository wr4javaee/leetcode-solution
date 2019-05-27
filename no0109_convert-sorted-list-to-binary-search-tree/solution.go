package no0109_convert_sorted_list_to_binary_search_tree

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	// 边界值校验
	if head == nil {
		return nil
	}

	// 递归退出条件
	// 当前链表节点无右临节点
	if head.Next == nil {
		return &TreeNode{ Val: head.Val }
	}

	// 初始化二叉树节点 - 链表中间节点
	// 取快慢指针确定中间节点
	slowLeft := head
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		// 慢指针每次只移动一次
		slowLeft = slow
		slow = slow.Next

		// 快指针每次移动两次
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	// 循环结束时快指针指向链表队尾
	// 慢指针指向链表中心

	// 初始化节点
	treeNodeVal := slow.Val
	treeNode := &TreeNode{ Val:treeNodeVal }

	// 递归设置右子树
	treeNode.Right = sortedListToBST(slow.Next)

	slowLeft.Next = nil
	// 递归设置左子树
	treeNode.Left = sortedListToBST(head)

	return treeNode
}
