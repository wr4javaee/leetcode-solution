package no0101_symmetric_tree

/**
 * 分而治之思想：递归
 * 思路：如果一个二叉树是对称二叉树，则它的左子树和右子树一定互为镜像
 * 两个二叉树互为镜像的条件：二叉树1的左子树 与 二叉树2的右子树互为镜像，且 二叉树2的左子树 与 二叉树1的右子树互为镜像
 */
func isSymmetric(root *TreeNode) bool {

	// 异常校验
	if root == nil {
		return true
	}

	return isMirror(root.Left, root.Right)
}

func isMirror(node1 *TreeNode, node2 *TreeNode) bool {
	// 两个二叉树都为空
	if node1 == nil && node2 == nil {
		return true
	}

	// 其中一个二叉树不为空
	if node1 == nil || node2 == nil {
		return false
	}

	// 判断两个子树节点的值
	if node1.Val != node2.Val {
		return false
	}

	// 两个子树根节点值相等，且两个子树的孩子对称则互为对称
	return isMirror(node1.Left, node2.Right) && isMirror(node1.Right, node2.Left)
}


