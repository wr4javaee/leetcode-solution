package no0224_basic_calculator

func calculate(s string) int {
	// step. 借助两个栈, 一个用来存放数字的数字栈，另一个用来存放操作符的操作符栈
	numStack := make([]int, 0, 1000)
	operStack := make([]string, 0, 1000)

	// 输入的数字大于9
	lastElementNum := false
	for _, ch := range s {
		// step. 将字符ascii转义为字符
		element := string(ch)
		numStackLen := len(numStack)
		operStackLen := len(operStack)

		// step. 判断是否是数字
		if 48 <= ch && ch <= 57 {
			currNum := int(ch) - 48
			if lastElementNum == true {
				newNum := numStack[numStackLen - 1] * 10 + currNum
				numStack[numStackLen - 1] = newNum
				continue
			}

			// 输入的是数字
			numStack = append(numStack, currNum)
			lastElementNum = true
			continue
		}

		// step. 判断输入的是操作符
		// 忽略空格
		if " " == element {
			continue
		}

		lastElementNum = false
		// (直接入栈
		if  "(" == element {
			operStack = append(operStack, element)
			continue
		}

		if "-" == element || "+" == element{
			// 和操作符栈顶比较
			// 若操作符栈为空, 直接入栈
			if operStackLen == 0 {
				operStack = append(operStack, element)
				continue
			}

			// 获得操作符栈栈顶
			operStackTop := operStack[operStackLen - 1]
			// 若栈顶优先级相等，则将数字栈的两个栈顶元素与操作符栈当前栈顶操作符出栈进行计算, 后将计算结果压入数字栈
			if operStackTop == "+" {
				// 数字栈的两个栈顶元素出栈并计算
				newNum := numStack[numStackLen - 1] + numStack[numStackLen - 2]
				numStack = numStack[: numStackLen - 2]
				// 当前栈顶操作符出栈
				operStack = operStack[: operStackLen - 1]
				// 计算结果压入数字栈
				numStack = append(numStack, newNum)
				// 新操作符入操作符栈
				operStack = append(operStack, element)
				continue
			}
			if operStackTop == "-" {
				// 数字栈的两个栈顶元素出栈并计算
				newNum := numStack[numStackLen - 2] - numStack[numStackLen - 1]
				numStack = numStack[: numStackLen - 2]
				// 当前栈顶操作符出栈
				operStack = operStack[: operStackLen - 1]
				// 计算结果压入数字栈
				numStack = append(numStack, newNum)
				// 新操作符入操作符栈
				operStack = append(operStack, element)
				continue
			}
			// 遇到(直接入栈
			if operStackTop == "(" {
				operStack = append(operStack, element)
				continue
			}
		}

		if ")" == element {
			// 和操作符栈顶比较
			// 若操作符栈顶为(, 则直接将(出栈
			// 若操作符栈顶为 + 或 - , 则将数字栈的两个栈顶元素与操作符栈当前栈顶操作符出栈进行计算, 后将计算结果压入数字栈
			// 获得操作符栈栈顶
			operStackTop := operStack[operStackLen - 1]

			// 若操作符栈顶为(, 则直接将(出栈
			if "(" == operStackTop {
				operStack = operStack[: operStackLen - 1]
				continue
			}

			// 若操作符栈顶为 + 或 - , 则将数字栈的两个栈顶元素与操作符栈当前栈顶操作符出栈进行计算, 后将计算结果压入数字栈
			if operStackTop == "+" {
				// 数字栈的两个栈顶元素出栈并计算
				newNum := numStack[numStackLen - 1] + numStack[numStackLen - 2]
				numStack = numStack[: numStackLen - 2]
				// 当前栈顶操作符出栈
				operStack = operStack[: operStackLen - 1]
				// 计算结果压入数字栈
				numStack = append(numStack, newNum)

				// 此时操作符栈顶为(, 新操作符舍弃，操作符栈顶出栈
				operStack = operStack[: operStackLen - 2]
				continue
			}
			if operStackTop == "-" {
				// 数字栈的两个栈顶元素出栈并计算
				newNum := numStack[numStackLen - 2] - numStack[numStackLen - 1]
				numStack = numStack[: numStackLen - 2]
				// 当前栈顶操作符出栈
				operStack = operStack[: operStackLen - 1]
				// 计算结果压入数字栈
				numStack = append(numStack, newNum)
				// 此时操作符栈顶为(, 新操作符舍弃，操作符栈顶出栈
				operStack = operStack[: operStackLen - 2]
				continue
			}
		}
	}

	// 循环后判定操作符栈是否为空
	operStackLen := len(operStack)
	numStackLen := len(numStack)
	resNum := 0
	if operStackLen > 0 {
		operStackTop := operStack[operStackLen - 1]
		if operStackTop == "+" {
			// 数字栈的两个栈顶元素出栈并计算
			resNum = numStack[0] + numStack[1]
			numStack = numStack[: numStackLen - 2]
		}
		if operStackTop == "-" {
			// 数字栈的两个栈顶元素出栈并计算
			resNum = numStack[0] - numStack[1]
			numStack = numStack[: numStackLen - 2]
		}
	}

	numStackLen = len(numStack)
	if numStackLen > 0 {
		resNum += numStack[0]
	}

	return resNum
}
