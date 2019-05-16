package no0147_insertion_sort_list

import (
	"fmt"
	"testing"
)

func Test(test *testing.T) {
	node1 := &ListNode{4, nil}
	node2 := &ListNode{2, nil}
	node3 := &ListNode{1, nil}
	node4 := &ListNode{3, nil}

	node3.Next = node4
	node2.Next = node3
	node1.Next = node2

	res := insertionSortList(node1)
	for res.Next != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
	fmt.Println(res.Val)
}
