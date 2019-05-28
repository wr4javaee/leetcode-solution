package no0114_flatten_binary_tree_to_linked_list

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
	// 基于前序遍历
	preOrder(root, root)

	// 将左子树都移动到右子树中
	moveToRight(root)
}

var lastNode *TreeNode

// 基于前序遍历
// 将遍历结果原地追加到左子树
func preOrder(root *TreeNode, currNode *TreeNode) {
	// 递归基线条件
	if currNode == nil {
		return
	}

	// 获得左右子树
	leftChildNode := currNode.Left
	rightChildNode := currNode.Right

	// 将每一次遍历的子树都作为上一个节点的左子树
	if lastNode == nil {
		lastNode = currNode
	} else {
		lastNode.Left = currNode
		lastNode = lastNode.Left
	}

	if leftChildNode != nil {
		preOrder(root, currNode.Left)
	}

	if rightChildNode != nil {
		preOrder(root, currNode.Right)
	}
}

// 将左子树都移动到右子树中
func moveToRight(node *TreeNode) {
	// 递归基线条件
	if node == nil {
		return
	}

	if node.Left != nil {
		node.Right = node.Left
		node.Left = nil
		moveToRight(node.Right)
	}
}
