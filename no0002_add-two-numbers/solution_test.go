package no0002_add_two_numbers

import (
	"fmt"
	"testing"
)

// 定义单元测试数据结构体
type testCaseModel struct {
	requestModel
	responseModel
}

// 定义被测试方法请求参数结构体
type requestModel struct {
	l1 *ListNode
	l2 *ListNode
}

// 定义被测试方法返回参数结构体
type responseModel struct {
	result *ListNode
}

// 单元测试
func TestTwoSum(t *testing.T) {
	// 构建测试数据
	testCaseArr := []testCaseModel {
		// 正常数据
		{
			requestModel {
				initListNode(&[]int{5}),
				initListNode(&[]int{5}),
			},
			responseModel {
				initListNode(&[]int{0, 1}),
			},
		},
		// 正常数据
		{
			requestModel {
				initListNode(&[]int{1, 8}),
				initListNode(&[]int{0}),
			},
			responseModel {
				initListNode(&[]int{1, 8}),
			},
		},
		// 正常数据
		{
			requestModel {
				initListNode(&[]int{2, 4, 3}),
				initListNode(&[]int{5, 6, 4}),
			},
			responseModel {
				initListNode(&[]int{7, 0, 8}),
			},
		},
	}

	// 执行测试方法
	// 测试方法一
	for i, v := range testCaseArr {
		expectRes := addTwoNumbers(v.l1, v.l2)
		if equals(expectRes, v.responseModel.result) {
			t.Log("测试通过", i)
		} else {
			t.Log("测试不通过", i)
			printListNode(expectRes)
			printListNode(v.responseModel.result)
		}
	}

}

// 初始化链表
func initListNode(vals *[]int) (listNode *ListNode) {
	// 定义一个哑节点
	listNode = new(ListNode)

	// 参数校验
	if *vals == nil {
		return
	}

	// 指向当前节点
	p := listNode

	for _, v := range *vals {
		l := &ListNode{v, nil}
		p.Next = l
		p = p.Next
	}

	return listNode.Next
}

// 打印链表
func printListNode(p *ListNode) {
	for p != nil {
		fmt.Print(p.Val)
		fmt.Print("-")
		p = p.Next
	}
	fmt.Println()
}

// 链表值比较
func equals(p1 *ListNode, p2 *ListNode) bool {
	if p1 == nil && p2 == nil {
		return true
	}

	for p1 != nil || p2 != nil {
		if (p1 == nil && p2 != nil) ||
			(p1 != nil && p2 == nil) {
			return false
		}
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	return true
}