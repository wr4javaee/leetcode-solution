package no0106_construct_binary_tree_from_inorder_and_postorder_traversal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inOrder []int, postOrder []int) *TreeNode {
	// 边界值校验
	inOrderLen := len(inOrder)
	postOrderLen := len(postOrder)
	if inOrderLen == 0 || postOrderLen == 0 {
		return nil
	}

	// 通过后序数组获得（子树）根节点的值
	// （后序数组最后一个值即根节点）
	nodeVal := postOrder[postOrderLen - 1]
	node := &TreeNode{ Val: nodeVal }
	// 递归退出条件
	// 若后序/中序数组长度为1, 说明当前节点已无左右子树了
	if postOrderLen == 1 || inOrderLen == 1 {
		return node
	}

	// 递归确认（子树）根节点的左子树和右子树
	// 通过（子树）根节点对应的中序数组获得左右子树的中序数组
	inOrderNodeInx := 0
	for i := 0; i < inOrderLen; i++ {
		if inOrder[i] == nodeVal {
			inOrderNodeInx = i
			break
		}
	}

	// 确认（子树）根节点的右子树
	childRightInOrder := make([]int, 0)
	if inOrderNodeInx < postOrderLen - 1 {
		childRightInOrder = inOrder[inOrderNodeInx + 1 :]
	}

	// 确认（子树）根节点的左子树
	childLeftInOrder := inOrder[:inOrderNodeInx]
	node.Left = buildTree(childLeftInOrder, postOrder[:postOrderLen - len(childRightInOrder) - 1])
	node.Right = buildTree(childRightInOrder, postOrder[:postOrderLen - 1])

	return node
}
