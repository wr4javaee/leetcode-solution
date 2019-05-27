package no0107_binary_tree_level_order_traversal_ii

import "container/list"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	// BFS
	// 用list模拟队列
	l := list.New()

	resMap:= make(map[int]*[]int, 0)

	// 异常校验
	if root == nil {
		return [][]int{}
	}

	// 初始化队列
	l.PushBack(root)

	// 初始化节点深度
	depth := make(map[*TreeNode]int, 0)
	depth[root] = 0

	// BFS遍历
	for l.Len() > 0 {
		// 出队
		element := l.Front()
		node := element.Value.(*TreeNode)
		l.Remove(element)

		dep, _ := depth[node]
		r, has := resMap[dep]
		if !has {
			tmp := make([]int, 0)
			r = &tmp
			resMap[dep] = r
		}
		*r = append(*r, node.Val)

		// 获得该节点的邻接表, 即左右子节点
		if node.Left != nil {
			depth[node.Left] = dep + 1
			l.PushBack(node.Left)
		}

		if node.Right != nil {
			depth[node.Right] = dep + 1
			l.PushBack(node.Right)
		}
	}

	res := make([][]int, len(resMap))
	maxDepth := len(resMap) - 1
	for k, v := range resMap {
		res[maxDepth - k] = *v
	}

	return res
}
