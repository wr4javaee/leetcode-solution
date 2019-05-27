package no0110_balanced_binary_tree

import "container/list"

func isBalanced(root *TreeNode) bool {
	// 边界值校验
	if root == nil {
		return true
	}

	// 边界值校验
	// 若一个二叉树的深度小于2, 则一定为高度平衡二叉树
	depth := treeDepth(root)
	if depth < 2 {
		return true
	}

	// 判断左右子树的高度差是否大于1
	leftDepth := treeDepth(root.Left)
	rightDepth := treeDepth(root.Right)
	if absSub(leftDepth, rightDepth) > 1 {
		return false
	}

	// 递归判断每一个节点是否是高度平衡二叉树
	return isBalanced(root.Left) && isBalanced(root.Right)
}

// BFS求树最大深度
// 同时判定是否是高度平衡二叉树
func treeDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	l := list.New()
	nodeDepthMap := make(map[*TreeNode]int, 0)
	nodeDepthMap[node] = 1
	maxDepth := 0
	l.PushBack(node)
	for l.Len() > 0 {
		element := l.Back()
		l.Remove(element)
		n := element.Value.(*TreeNode)
		depth := nodeDepthMap[n]
		if maxDepth < depth {
			maxDepth = depth
		}

		if n.Left != nil {
			l.PushBack(n.Left)
			nodeDepthMap[n.Left] = depth + 1
		}
		if n.Right != nil {
			l.PushBack(n.Right)
			nodeDepthMap[n.Right] = depth + 1
		}
	}

	return maxDepth
}

func absSub(leftDepth int, rightDepth int) int {
	if leftDepth >= rightDepth {
		return leftDepth - rightDepth
	} else {
		return rightDepth - leftDepth
	}
}
