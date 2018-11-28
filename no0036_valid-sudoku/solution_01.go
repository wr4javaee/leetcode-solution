package no0036_valid_sudoku

import (
	"strconv"
	"strings"

	//"strings"
)

func isValidSudoku(board [][]byte) bool {
	// row 转义为 rowNo + 'r' + val
	// column 转义为 columnNo + 'r' + val
	// grid 转义为 rowNo + columnNo + val
	valMap := make(map[string]string, 89)
	// 按顺序遍历每一个九宫格数字
	for row, valArr := range board {
		for column, valByte := range valArr {
			valStr := convertByteToStr(valByte)
			// 不考虑空格.
			if strings.Compare(valStr, ".") == 0 {
				continue
			}
			// 对数字进行转义
			rowKey := convertRowKey(row, valStr)
			columnKey := convertColumnKey(column, valStr)
			gridKey := convertGridKey(row, column, valStr)
			// 判断数字是否重复, 重复则为不合法九宫格
			if _, ok := valMap[rowKey]; ok {
				return false
			}
			if _, ok := valMap[columnKey]; ok {
				return false
			}
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

// byte to string
func convertByteToStr(b byte) string {
	return string(b)
}

// 转义行数字
func convertRowKey(rowNo int, val string) (rowKey string) {
	rowKey = strconv.Itoa(rowNo) + "r" + val
	return
}

// 转义列数字
func convertColumnKey(columnNo int, val string) (columnKey string) {
	columnKey = strconv.Itoa(columnNo) + "c" + val
	return
}

// 转义小九宫格数字
func convertGridKey(rowNo int, columnNo int, val string) (gridKey string) {
	rowNo = rowNo / 3
	columnNo = columnNo / 3
	gridKey = strconv.Itoa(rowNo) + strconv.Itoa(columnNo) + val
	return
}