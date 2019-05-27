package no0105_construct_binary_tree_from_preorder_and_inorder_traversal

import (
	"fmt"
	"testing"
)

func TestBuildTree(t *testing.T) {
	preorder := []int{ 3,2,1,4,5 }
	inorder := []int{ 1,2,3,4,5 }
	res := buildTree(preorder, inorder)
	fmt.Printf("%v", res)
}
