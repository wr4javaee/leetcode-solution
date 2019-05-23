package no0144_binary_tree_preorder_traversal

import (
	"container/list"
)

//
//给定一个二叉树，返回它的 前序 遍历。
//
//示例:
//
//输入: [1,null,2,3]
//		1
//		 \
//		  2
//		 /
//		3
//
//输出: [1,3,2]
//进阶: 递归算法很简单，你可以通过迭代算法完成吗？
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 前序遍历
func preorderTraversal(root *TreeNode) []int {
	// 定义返回结果切片
	res := make([]int, 0)

	// 递归方法
	preOrder4Recusion(root, &res)

	return res
}

// 递归方法
func preOrder4Recusion(node *TreeNode, res *[]int) {
	// 递归基线
	for node == nil {
		return
	}

	// 储存值
	*res = append(*res, node.Val)

	preOrder4Recusion(node.Left, res)
	preOrder4Recusion(node.Right, res)
}

// 递归算法很简单，你可以通过迭代算法完成吗？
// 执行用时 : 0 ms, 在Binary Tree Preorder Traversal的Go提交中击败了100.00% 的用户
func preOrder(root *TreeNode) []int {
	// 定义返回结果切片
	res := make([]int, 0)

	// 异常校验
	if root == nil {
		return res
	}

	// 借助golang的双向链表list实现栈的功能
	l := list.New()

	// 将root节点放入栈
	// 相当于push
	l.PushBack(root)

	// 模拟方法执行的栈过程
	for l.Len() > 0 {
		// 相当于pop
		element := l.Back()
		l.Remove(element)

		// 先序遍历结果存入返回切片中
		node := element.Value.(*TreeNode)
		res = append(res, node.Val)

		if node.Right != nil {
			l.PushBack(node.Right)
		}

		if node.Left != nil {
			l.PushBack(node.Left)
		}
	}

	return res
}
