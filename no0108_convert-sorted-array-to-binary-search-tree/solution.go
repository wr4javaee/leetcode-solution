package no0108_convert_sorted_array_to_binary_search_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	numsLen := len(nums)

	// 边界值校验
	if numsLen == 0 {
		return nil
	}

	// 因为是高度平衡二叉树, 每次取数组中间作为节点即可满足要求
	midInx := numsLen / 2
	nodeVal := nums[midInx]
	node := &TreeNode{ Val:nodeVal }

	// 递归退出条件
	// 若数组仅有一个数据, 则该节点没有左右子树
	if numsLen == 1 {
		return node
	}

	// 递归设置左子树
	node.Left = sortedArrayToBST(nums[0: midInx])

	// 递归设置右子树
	if midInx < numsLen - 1 {
		node.Right = sortedArrayToBST(nums[midInx + 1 :])
	}

	return node
}
