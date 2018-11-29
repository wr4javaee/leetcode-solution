# LeetCode.0002 add-two-numbers 两数相加
---

## 题目
---
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

```
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
```

## 方法一
---
#### 解题思路
参考
https://leetcode-cn.com/problems/add-two-numbers/solution/

#### 复杂度分析
* 时间复杂度 O( n )

#### 执行效率
* 执行用时: 24 ms, 在Add Two Numbers的Go提交中击败了87.50% 的用户

## 实现（Golang）
---

```
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 校验空链表
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	// 初始化返回链表
	resNode := new(ListNode)

	// 初始化两个新指针
	p1 := l1
	p2 := l2
	p3 := resNode

	// 按序遍历链表
	// 设置默认进位为0
	carry := 0
	for p1 != nil || p2 != nil || carry > 0 {
		p1Val := 0
		if p1 != nil {
			p1Val = p1.Val
		}
		p2Val := 0
		if p2 != nil {
			p2Val = p2.Val
		}
		p3Val := p1Val + p2Val + carry
		// 考虑进位问题
		if p3Val >= 10 {
			// 需要进位
			p3Val -= 10
			carry = 1
		} else {
			// 不用进位, 恢复进位flag
			carry = 0
		}

		// 链接返回节点
		tmpNode := new(ListNode)
		tmpNode.Val = p3Val
		p3.Next = tmpNode
		p3 = p3.Next

		// 遍历入参节点
		if p1 != nil {
			p1 = p1.Next
		}
		if p2 != nil {
			p2 = p2.Next
		}
	}

	return resNode.Next
}
```
## 单元测试（Golang）
---

```
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
```

