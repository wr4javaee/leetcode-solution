package no0094_binary_tree_inorder_traversal

import "container/list"

//给定一个二叉树，返回它的 中序 遍历。
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
//输出: [1,2,3]
//进阶: 递归算法很简单，你可以通过迭代算法完成吗？
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)

	inOrderRecusion(root, &res)

	return res
}

// 递归方法
func inOrderRecusion(node *TreeNode, res *[]int) {

	// 递归基线
	if node == nil {
		return
	}

	// 先递归左子树
	inOrderRecusion(node.Left, res)

	// 存储
	*res = append(*res, node.Val)

	// 递归右子树
	inOrderRecusion(node.Right, res)
}

// 非递归方法
func inOrder(root *TreeNode) []int {
	// 非递归, 引入双向链表list作为栈
	l := list.New()
	res := make([]int, 0)

	// 异常校验
	if root == nil {
		return res
	}

	// 为即将入栈的根节点分配角色
	// find left | print | end
	roles := make(map[*TreeNode]string, 0)

	// 入栈
	l.PushBack(root)
	roles[root] = "find left"

	// 模拟递归方法栈
	for l.Len() > 0 {
		// 出栈, 等同于pop
		element := l.Back()

		// 类型转化
		node := element.Value.(*TreeNode)
		// 判断node的角色
		// 输出值
		role, _ := roles[node]

		// 需要找左子树, 入栈
		if role == "find left" {
			roles[node] = "print"

			if node.Left != nil {
				roles[node.Left] = "find left"
				l.PushBack(node.Left)
			}
			continue
		}

		// 需要输出值
		if role == "print" {
			res = append(res, node.Val)
			roles[node] = "end"
			// 需要找到右子树
			if node.Right != nil {
				roles[node.Right] = "find left"
				l.PushBack(node.Right)
			}
			continue
		}

		// 需要出栈
		if role == "end" {
			delete(roles, node)
			l.Remove(element)
		}

	}

	return res
}
