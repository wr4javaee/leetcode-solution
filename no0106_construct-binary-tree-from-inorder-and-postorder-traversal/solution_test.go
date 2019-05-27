package no0106_construct_binary_tree_from_inorder_and_postorder_traversal

import "testing"

func Test(t *testing.T) {
	inOrder := []int{ 9,3,15,20,7 }
	postOrder := []int{ 9,15,7,20,3 }

	buildTree(inOrder, postOrder)
}
