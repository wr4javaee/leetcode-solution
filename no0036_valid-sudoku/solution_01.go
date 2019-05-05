package no0036_valid_sudoku

import (
	"bytes"
	"strconv"
	"strings"

	//"strings"
)

// 执行用时: 12 ms, 在Valid Sudoku的Go提交中击败了29.76% 的用户
func isValidSudoku(board [][]byte) bool {
	// row 转义为 rowNo + 'r' + val
	// column 转义为 columnNo + 'r' + val
	// grid 转义为 rowNo + columnNo + val
	valMap := make(map[string]string, 243)
	// 按顺序遍历每一个九宫格数字
	for row, valArr := range board {
		for column, valByte := range valArr {
			valStr := string(valByte)
			// 不考虑空格.
			if strings.Compare(valStr, ".") == 0 {
				continue
			}
			// 对数字进行转义
			rowKey := convertRowKey(row, valStr)
			// 判断数字是否重复, 重复则为不合法九宫格
			if _, ok := valMap[rowKey]; ok {
				return false
			}

			columnKey := convertColumnKey(column, valStr)
			if _, ok := valMap[columnKey]; ok {
				return false
			}

			gridKey := convertGridKey(row, column, valStr)
			if _, ok := valMap[gridKey]; ok {
				return false
			}

			// 将转义后数字存至map
			valMap[rowKey] = rowKey
			valMap[columnKey] = columnKey
			valMap[gridKey] = gridKey
		}
	}
	return true
}

// 转义行数字,  rowNo + 'r' + val
func convertRowKey(rowNo int, val string) string {
	var b bytes.Buffer
	b.WriteString(strconv.Itoa(rowNo))
	b.WriteString("r")
	b.WriteString(val)
	return b.String()
}

// 转义列数字, columnNo + 'r' + val
func convertColumnKey(columnNo int, val string) string{
	var b bytes.Buffer
	b.WriteString(strconv.Itoa(columnNo))
	b.WriteString("c")
	b.WriteString(val)
	return b.String()
}

// 转义小九宫格数字, rowNo + columnNo + val
func convertGridKey(rowNo int, columnNo int, val string) string {
	rowNo = rowNo / 3
	columnNo = columnNo / 3
	var b bytes.Buffer
	b.WriteString(strconv.Itoa(rowNo))
	b.WriteString(strconv.Itoa(columnNo))
	b.WriteString(val)
	return b.String()
}