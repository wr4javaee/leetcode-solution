package no0145_binary_tree_postorder_traversal

import "container/list"

//
//给定一个二叉树，返回它的 后序 遍历。
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
//输出: [3,2,1]
//进阶: 递归算法很简单，你可以通过迭代算法完成吗？
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 后序遍历
func postorderTraversal(root *TreeNode) []int {
	// 返回结果切片
	res := make([]int, 0)

	// 递归
	postorderTraversal4Recusion(root, &res)

	// 返回结果
	return res
}

// 递归方法
func postorderTraversal4Recusion(node *TreeNode, res *[]int) {
	// 递归退出条件
	if node == nil {
		return
	}

	// 左子树
	if node.Left != nil {
		postorderTraversal4Recusion(node.Left, res)
	}

	// 右子树
	if node.Right != nil {
		postorderTraversal4Recusion(node.Right, res)
	}

	// 输出值
	*res = append(*res, node.Val)
}

// 非递归
func  postorder(root *TreeNode) []int {
	// 返回结果切片
	res := make([]int, 0)

	// 异常校验
	if root == nil {
		return res
	}

	// 初始化方法栈
	// 采用双向链表list模拟栈
	l := list.New()

	// 初始化栈角色
	// find left | find right | print and end
	roles := make(map[*TreeNode]string, 0)

	// 初始化方法栈值
	l.PushBack(root)
	roles[root] = "find left"

	// 模拟方法栈
	for l.Len() > 0 {
		// 找到栈底数据
		element := l.Back()
		// convert
		node := element.Value.(*TreeNode)
		role, _ := roles[node]

		// 根据栈对象角色进行后续处理
		if role == "print and end" {
			// 输出值
			res = append(res, node.Val)

			// 出栈
			l.Remove(element)
			delete(roles, node)
			continue
		}

		if role == "find left" {
			roles[node] = "find right"
			if node.Left != nil {
				roles[node.Left] = "find left"
				l.PushBack(node.Left)
			}
			continue
		}

		if role == "find right" {
			roles[node] = "print and end"
			if node.Right != nil {
				roles[node.Right] = "find left"
				l.PushBack(node.Right)
			}
			continue
		}

	}

	return res
}
