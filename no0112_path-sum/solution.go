package no0112_path_sum

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
	return dfs(root, 0, sum)
}

// dfs
func dfs(node *TreeNode, currSum int, targetSum int) bool {
	// 边界条件判定
	if node == nil {
		return false
	}

	// 累加自根节点以来的路径和
	currSum += node.Val

	// 判断当前节点是否为叶子节点
	if node.Left == nil && node.Right == nil {
		return currSum == targetSum
	}

	// 非叶子节点, 继续递归
	// 只要两个节点有一个满足要求, 即视为满足要求
	return dfs(node.Left, currSum, targetSum) || dfs(node.Right, currSum, targetSum)
}
