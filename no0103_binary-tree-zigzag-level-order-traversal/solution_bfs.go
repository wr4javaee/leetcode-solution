package no0103_binary_tree_zigzag_level_order_traversal

import "container/list"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	// bfs

	// 异常校验
	if root == nil {
		return [][]int{}
	}

	// 保存每一深度对应的返回结果
	depthListMap := make(map[int]*list.List)
	// 保存每一节点的深度
	depthValMap := make(map[*TreeNode]int)

	// bfs遍历
	// 初始化默认深度
	depthVal := 0
	depthValMap[root] = 0
	l := list.New()
	l.PushBack(root)

	for l.Len() > 0 {
		element := l.Front()
		l.Remove(element)
		node := element.Value.(*TreeNode)
		depthVal = depthValMap[node]

		depthList, has := depthListMap[depthVal]
		if !has {
			depthList = list.New()
		}

		depthList.PushBack(node.Val)
		depthListMap[depthVal] = depthList

		if node.Left != nil {
			depthValMap[node.Left] = depthVal + 1
			l.PushBack(node.Left)
		}
		if node.Right != nil {
			depthValMap[node.Right] = depthVal + 1
			l.PushBack(node.Right)
		}
	}

	res := make([][]int, len(depthListMap))
	for k, v := range depthListMap {
		res[k] = make([]int, v.Len())

		for i := 0; v.Len() > 0; i ++ {
			element := &list.Element{}
			if k % 2 != 0 {
				// 单竖行从右往左
				element = v.Back()
			} else {
				// 偶数行从左往右
				element = v.Front()
			}
			v.Remove(element)
			res[k][i] = element.Value.(int)
		}
	}

	return res
}
