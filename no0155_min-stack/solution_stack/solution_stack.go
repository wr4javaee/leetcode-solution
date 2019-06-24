package solution_stack

import "container/list"

type MinStack struct {
	// 正常的栈
	ValStack *list.List
	// 保留最小值的栈
	MinValStack *list.List
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{ list.New(), list.New() }
}


func (this *MinStack) Push(x int)  {
	// 入栈规则：
	// 1、x正常入ValStack
	this.ValStack.PushBack(x)
	// 2、判断MinValStack是否为空, 若为空则入MinValStack
	if this.MinValStack.Len() == 0 {
		this.MinValStack.PushBack(x)
		return
	}

	// 3、MinValStack不为空，则比较MinValStack栈顶和x谁小，若x小于等于栈顶，则入栈，否则不做操作
	minValStackTop := this.MinValStack.Back().Value.(int)
	if x <= minValStackTop {
		this.MinValStack.PushBack(x)
	}
}


func (this *MinStack) Pop()  {
	// 出栈规则：
	// 1、正常出ValStack
	valStackTopElement := this.ValStack.Back()
	this.ValStack.Remove(valStackTopElement)
	valStackTop := valStackTopElement.Value.(int)
	// 2. 比较valStackTop与MinValStack的栈顶元素，若valStackTop小于等于MinValStack的栈顶，则也弹出MinValStack的栈顶
	minValStackTopElement := this.MinValStack.Back()
	minValStackTop := minValStackTopElement.Value.(int)
	if valStackTop <= minValStackTop {
		this.MinValStack.Remove(minValStackTopElement)
	}
}


func (this *MinStack) Top() int {
	if this.ValStack.Len() == 0 {
		return 0
	}
	// 直接返回ValStack栈顶
	return this.ValStack.Back().Value.(int)
}


func (this *MinStack) GetMin() int {
	if this.MinValStack.Len() == 0 {
		return 0
	}
	// 直接返回ValStack栈顶
	return this.MinValStack.Back().Value.(int)
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
