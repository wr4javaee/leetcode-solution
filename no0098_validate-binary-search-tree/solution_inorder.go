package no0098_validate_binary_search_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	// 中序, 判定数组是否升序
	// 异常校验
	if root == nil {
		return true
	}

	arr := make([]int, 0)

	// 中序遍历判定
	isValid := true
	inOrderValid(root, &arr, &isValid)
	if isValid {
		// 递归结束后依然要判定最后的结果位
		arrLen := len(arr)
		// 数组不是严格升序,  判定为非二叉搜索树
		if arrLen > 1 && arr[arrLen - 1] <= arr[arrLen - 2] {
			isValid = false
		}
	}

	return isValid
}

func inOrderValid(node *TreeNode, arr *[]int, isValid *bool) {
	if ! *isValid {
		// 已经判断出不满足二叉搜索树, 直接返回结果
		return
	}

	if node == nil {
		return
	}

	// 判断左子树
	if node.Left != nil {
		inOrderValid(node.Left, arr, isValid)
	}

	// 判断根节点
	// 判定数组是否有序
	arrLen := len(*arr)
	// 数组不是严格升序,  判定为非二叉搜索树
	if arrLen > 1 && (*arr)[arrLen - 1] <= (*arr)[arrLen - 2] {
		*isValid = false
	}

	*arr = append(*arr, node.Val)
	if node.Right != nil {
		inOrderValid(node.Right, arr, isValid)
	}
}
