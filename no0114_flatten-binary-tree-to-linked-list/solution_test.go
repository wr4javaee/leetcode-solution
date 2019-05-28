package no0114_flatten_binary_tree_to_linked_list

import "testing"

func Test(t *testing.T) {
	node1 := &TreeNode{ Val:1 }
	node2 := &TreeNode{ Val:2 }
	node5 := &TreeNode{ Val:5 }
	node3 := &TreeNode{ Val:3 }
	node4 := &TreeNode{ Val:4 }
	node6 := &TreeNode{ Val:6 }
	node2.Left = node3
	node2.Right = node4
	node5.Right = node6
	node1.Left = node2
	node1.Right = node5
	flatten(node1)
}
