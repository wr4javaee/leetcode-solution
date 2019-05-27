package no0104_maximum_depth_of_binary_tree

import "container/list"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	// 异常校验
	if root == nil {
		return 0
	}

	// BFS
	depth := 1
	// 借助list实现栈
	l := list.New()
	// 保存节点的深度
	nodeDepthMap := make(map[*TreeNode]int)
	nodeDepthMap[root] = depth

	// 初始化栈
	l.PushBack(root)
	maxDepthVal := depth
	for l.Len() > 0 {
		// 出栈
		element := l.Back()
		l.Remove(element)
		node := element.Value.(*TreeNode)
		nodeDepth, _ := nodeDepthMap[node]

		// 保存最大节点深度
		if maxDepthVal < nodeDepth {
			maxDepthVal = nodeDepth
		}

		if node.Left != nil {
			l.PushBack(node.Left)
			nodeDepthMap[node.Left] = nodeDepth + 1
		}

		if node.Right != nil {
			l.PushBack(node.Right)
			nodeDepthMap[node.Right] = nodeDepth + 1
		}
	}

	return maxDepthVal
}
