package no0129_sum_root_to_leaf_numbers

import "container/list"

// val为最近到达key节点的上一个节点
var edgeToMap = make(map[*TreeNode]*TreeNode)
// 根节点到叶子节点的栈
var leafStackMap = make(map[*TreeNode]*list.List)


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	// 基于DFS, 保存所有根节点到叶子节点的路径
	dfs(root)

	// 计算题目要求结果
	res := 0
	for _, nodeStack := range leafStackMap {
		nodeSum := 0
		for nodeStack.Len() > 0 {
			nodeElement := nodeStack.Back()
			nodeStack.Remove(nodeElement)
			node := nodeElement.Value.(*TreeNode)
			nodeSum = nodeSum * 10 + node.Val
		}
		res += nodeSum
	}

	return res
}

// 基于DFS保存根节点到叶子节点的路径
func dfs(node *TreeNode) {
	// 递归基线
	if node == nil {
		return
	}

	// 判定是否是叶子节点
	if node.Left == nil && node.Right == nil {
		// 根节点到叶子节点的栈
		nodeStack := list.New()
		nodeStack.PushBack(node)
		// 对于叶子节点, 从edgeToMap中回溯栈并保存到leafMap中
		lastNode, has := edgeToMap[node]
		edgeToStack := list.New()
		if has {
			edgeToStack.PushBack(lastNode)
		}
		// 回溯edgeToMap
		for edgeToStack.Len() > 0 {
			element := edgeToStack.Back()
			edgeToStack.Remove(element)
			tmpNode := element.Value.(*TreeNode)
			nodeStack.PushBack(tmpNode)

			tmpLastNode, has := edgeToMap[tmpNode]
			if has {
				edgeToStack.PushBack(tmpLastNode)
			}
		}
		leafStackMap[node] = nodeStack
		return
	}

	// 继续下一次递归
	if node.Left != nil {
		edgeToMap[node.Left] = node
		dfs(node.Left)
	}
	if node.Right != nil {
		edgeToMap[node.Right] = node
		dfs(node.Right)
	}
}
