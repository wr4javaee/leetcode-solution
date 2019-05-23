package no0144_binary_tree_preorder_traversal

import (
"fmt"
"testing"
)

func TestPreorderTraversal(test *testing.T) {
	root := &TreeNode{1, nil, nil}
	node2 := &TreeNode{2, nil, nil}
	node3 := &TreeNode{3, nil, nil}

	node2.Left = node3
	root.Right = node2

	res := preorderTraversal(root)
	fmt.Printf("%v", res)
}

func TestPreOrder(test *testing.T) {
	root := &TreeNode{1, nil, nil}
	node2 := &TreeNode{2, nil, nil}
	node3 := &TreeNode{3, nil, nil}

	node2.Left = node3
	root.Right = node2

	res := preOrder(root)
	fmt.Printf("%v", res)
}
