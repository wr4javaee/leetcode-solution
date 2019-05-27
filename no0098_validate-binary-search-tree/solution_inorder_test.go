package no0098_validate_binary_search_tree

import "testing"

func TestIsValidBST(t *testing.T) {

	rootNode := &TreeNode{ Val:1 }
	leftNode := &TreeNode{ Val:1 }
	rootNode.Left = leftNode
	res := isValidBST(rootNode)
	if res {
		t.Fatal("error result")
	}
}
