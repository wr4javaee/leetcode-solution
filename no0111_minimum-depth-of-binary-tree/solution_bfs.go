package no0111_minimum_depth_of_binary_tree

import "container/list"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	// 思路: BFS

	// 异常校验
	if root == nil {
		return 0
	}

	// 记录到达节点的最近节点
	edgeTo := make(map[*TreeNode]*TreeNode, 0)

	// 初始化队列
	l := list.New()
	l.PushBack(root)

	// 遍历队列
	var minDepthNode *TreeNode
	for l.Len() > 0 {
		// 出队列
		element := l.Front()
		l.Remove(element)
		node := element.Value.(*TreeNode)

		// 判断是否找到最小深度节点
		if node.Left == nil && node.Right == nil {
			// 找到了最小深度节点
			minDepthNode = node
			break
		}

		if node.Left != nil {
			edgeTo[node.Left] = node
			l.PushBack(node.Left)
		}

		if node.Right != nil {
			edgeTo[node.Right] = node
			l.PushBack(node.Right)
		}
	}

	// 求解根节点到最小深度节点的最小路径
	if minDepthNode == root {
		return 1
	}

	tmpNode := edgeTo[minDepthNode]
	res := 2
	for tmpNode != root {
		tmpNode = edgeTo[tmpNode]
		res ++
	}

	return res
}
