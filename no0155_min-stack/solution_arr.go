package main

type MinStack struct {
	ValStack []int
	MinValStack []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{make([]int, 0, 1000), make([]int, 0, 1000)}
}


func (this *MinStack) Push(x int)  {
	this.ValStack = append(this.ValStack, x)
	minValStackLen := len(this.MinValStack)

	// 若MinValStack为空，则x入MinValStack
	if minValStackLen == 0 {
		this.MinValStack = append(this.MinValStack, x)
		return
	}

	// 若MinValStack不为空，且x小于等于MinValStack的栈顶，x入MinValStack
	minVal := this.MinValStack[minValStackLen - 1]
	if x <= minVal {
		this.MinValStack = append(this.MinValStack, x)
		return
	}
}


func (this *MinStack) Pop()  {
	// ValStack出栈. 出栈元素设为x
	valStackLen := len(this.ValStack)
	if valStackLen == 0 {
		return
	}
	x := this.ValStack[valStackLen - 1]
	this.ValStack = this.ValStack[: valStackLen - 1]

	// 若x小于等于MinValStack栈顶， 则MinValStack栈顶出栈
	minValStackLen := len(this.MinValStack)
	if minValStackLen == 0 {
		return
	}
	y := this.MinValStack[minValStackLen - 1]
	if x <= y {
		this.MinValStack = this.MinValStack[: minValStackLen - 1]
		return
	}
}


func (this *MinStack) Top() int {
	valStackLen := len(this.ValStack)
	if valStackLen == 0 {
		return 0
	}
	return this.ValStack[valStackLen - 1]
}


func (this *MinStack) GetMin() int {
	minValStackLen := len(this.MinValStack)
	if minValStackLen == 0 {
		return 0
	}
	return this.MinValStack[minValStackLen - 1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
