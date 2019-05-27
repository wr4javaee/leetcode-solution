package no0101_symmetric_tree

import "container/list"

// 定义方法栈元素
type StackNode struct {
	leftNode *TreeNode
	rightNode *TreeNode
}

// iteration
func isSymmetric(root *TreeNode) bool {
	// 异常校验
	if root == nil {
		return true
	}

	// 模拟方法栈
	l := list.New()

	// 初始化方法栈
	l.PushBack(&StackNode{root.Left, root.Right})

	// 初始化结果
	res := false

	// 遍历方法栈
	for l.Len() > 0 {
		// 出栈
		element := l.Back()
		l.Remove(element)
		stackNode := element.Value.(*StackNode)
		leftNode := stackNode.leftNode
		rightNode := stackNode.rightNode

		// 设置循环结束条件
		if leftNode == nil && rightNode == nil {
			res = true
			continue
		}

		if leftNode == nil || rightNode == nil {
			res = false
			break
		}

		if leftNode.Val != rightNode.Val {
			res = false
			break
		}

		// 入栈
		l.PushBack(&StackNode{leftNode.Left, rightNode.Right})
		l.PushBack(&StackNode{leftNode.Right, rightNode.Left})
	}

	return res
}
