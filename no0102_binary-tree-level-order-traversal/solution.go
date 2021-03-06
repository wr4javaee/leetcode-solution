package no0102_binary_tree_level_order_traversal

import "container/list"

/**
给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。

例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]
*/
func levelOrder(root *TreeNode) [][]int {
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
	for k, v := range resMap {
		res[k] = *v
	}

	return res
}
