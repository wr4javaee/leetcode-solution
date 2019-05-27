package no0100_same_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// 基于dfs递归
	// 若两个二叉树相等, 则根节点、左子树、右子树都相等

	// 递归退出条件
	// p、q均为空
	if p == nil && q == nil {
		return true
	}

	// p、q中有一个为空
	if p == nil || q == nil {
		return false
	}

	// p、q节点值不同
	if p.Val != q.Val {
		return false
	}

	// p、q均不为空且p、q值相等
	// 比对两个二叉树的左右子树
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}


