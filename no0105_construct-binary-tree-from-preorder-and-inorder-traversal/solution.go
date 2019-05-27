package no0105_construct_binary_tree_from_preorder_and_inorder_traversal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preOrder []int, inOrder []int) *TreeNode {
	preOrderLen := len(preOrder)
	inOrderLen := len(inOrder)

	// 递归结束条件
	// 异常条件
	if preOrderLen == 0 || inOrderLen == 0 {
		return nil
	}

	// 获得（子树）根节点值
	nodeVal := preOrder[0]
	node := &TreeNode{ Val: nodeVal }

	// 当前序数组长度为1时, 说明该节点已无子树
	if preOrderLen == 1 {
		return node
	}

	// 当中序数组为1时, 说明该节点也已无子树
	if inOrderLen == 1 {
		return node
	}

	// 从中序数组中定位（子树）根节点索引
	nodeInx := 0
	for i := 0; i < inOrderLen; i++ {
		if inOrder[i] == nodeVal {
			nodeInx = i
			break
		}
	}

	// 定位左子树
	preOrder = preOrder[1:]
	leftInOrder := inOrder[:nodeInx]
	rightInOrder := make([]int, 0)
	node.Left = buildTree(preOrder, leftInOrder)

	// 定位右子树
	// 关键点：注意右子树的先序数组起始位置
	preOrder = preOrder[len(leftInOrder):]
	if nodeInx < inOrderLen - 1 {
		rightInOrder = inOrder[nodeInx + 1 :]
	}
	node.Right = buildTree(preOrder, rightInOrder)
	return node
}
