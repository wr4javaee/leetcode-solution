package no0113_path_sum_ii

import (
	"container/list"
)

// 定义子节点 - 父节点map
var nodeMap = make(map[*TreeNode]*TreeNode, 0)
var resMap = make(map[*TreeNode]*list.List, 0)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	// 基于DFS关联节点路径
	dfs(root, nil, sum)

	// 返回结果数组
	res := make([][]int, len(resMap))
	for _, nodeValStack := range resMap {
		arr := make([]int, nodeValStack.Len())
		j := 0
		for nodeValStack.Len() > 0 {
			element := nodeValStack.Back()
			nodeValStack.Remove(element)
			arr[j] = element.Value.(int)
			j ++
		}
		res = append(res, arr)
	}

	return res
}

// 深度优先记录根节点到每一个叶子节点的路径
// 同时算路径节点val和与目标值是否一致
func dfs(currNode *TreeNode, lastNode *TreeNode, sum int) {
	// 递归基线条件
	if currNode == nil {
		return
	}

	// 关联节点路径
	if lastNode != nil {
		nodeMap[currNode] = lastNode
	}

	// 叶子节点, 查找根节点到叶子节点路径总和是否满足要求
	if currNode.Left == nil && currNode.Right == nil {
		// 从nodeMap中寻找根节点到叶子节点的完整路径
		nodeValStack := list.New()
		nodeValStack.PushBack(currNode.Val)
		nodeSum := currNode.Val

		// 借助栈对nodeMap遍历
		l := list.New()
		lastNode, has := nodeMap[currNode]
		if has {
			nodeSum += lastNode.Val
			l.PushBack(lastNode)
			nodeValStack.PushBack(lastNode.Val)
		}
		for l.Len() > 0 {
			tmpNodeElement := l.Back()
			l.Remove(tmpNodeElement)
			tmpNode := tmpNodeElement.Value.(*TreeNode)
			tmpParentNode, has := nodeMap[tmpNode]
			if has {
				nodeSum += tmpParentNode.Val
				nodeValStack.PushBack(tmpParentNode.Val)
				l.PushBack(tmpParentNode)
			}
		}

		// 判断从根节点到叶子节点总和是否和目标值相同
		if nodeSum == sum {
			// 将结果保存
			resMap[currNode] = nodeValStack
		}
	}

	// 非叶子节点, 继续递归
	if currNode.Left != nil {
		dfs(currNode.Left, currNode, sum)
	}

	if currNode.Right != nil {
		dfs(currNode.Right, currNode, sum)
	}
}
