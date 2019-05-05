package no0006_zigzag_conversion

import "fmt"

// 执行用时: 128 ms, 在ZigZag Conversion的Go提交中击败了3.79% 的用户
func convert(s string, numRows int) string {

	if numRows <= 1 {
		return s
	}
	// 初始化保存Z结果集的map, 其中key为行号, value为每一行的字符串
	resultMap := make(map[int]string, numRows)

	// 定义当前行号, 1代表第一行
	currRowNo := 1
	// 定义方向标识, 默认向下走
	goingDown := true

	// 循环遍历字符串
	for _, v := range s {
		// 按照行号从map中取出字符串
		rowStr, isPresent := resultMap[currRowNo]
		if ! isPresent {
			rowStr = ""
		}
		// 从转义前的字符串中取出字符
		srcStr := string(v)

		// step. 如果向下走
		if goingDown {
			// 向下走时, 原始字符直接拼接到新转义的字符串之后
			rowStr += srcStr
			resultMap[currRowNo] = rowStr

			// 增到numRows, 变换方向
			if currRowNo == numRows {
				goingDown = false
				currRowNo --
			} else {
				// 行号自增, 直到增到numRows
				currRowNo ++
			}
			continue
		}

		// step. 向上走
		if ! goingDown {
			// 向上走时
			rowStr += srcStr
			resultMap[currRowNo] = rowStr

			// 减到1, 变换方向
			if currRowNo == 1 {
				goingDown = true
				currRowNo ++
			} else {
				// 行号自减, 直到减到1
				currRowNo --
			}
			continue
		}
	}

	resultStr := ""
	for i := 1; i <= numRows; i++ {
		resultStr += resultMap[i]
		fmt.Println(resultMap[i])
	}
	return resultStr
}