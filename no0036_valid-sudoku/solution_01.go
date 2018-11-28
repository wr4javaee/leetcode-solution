package no0036_valid_sudoku

import (
	"strconv"
	"strings"
)

func isValidSudoku(board [][]byte) bool {
	// row
	// column
	// grid
	valMap := make(map[string]string, 89)
	for row, valArr := range board {
		for column, valByte := range valArr {
			valStr := convertByteToStr(valByte)
			if strings.Compare(valStr, ".") == 0 {
				continue
			}
			rowKey := convertRowKey(row, valStr)
			columnKey := convertColumnKey(column, valStr)
			gridKey := convertGridKey(row, column, valStr)
			if _, ok := valMap[rowKey]; ok {
				return false
			}
			if _, ok := valMap[columnKey]; ok {
				return false
			}
			if _, ok := valMap[gridKey]; ok {
				return false
			}
			valMap[rowKey] = rowKey
			valMap[columnKey] = columnKey
			valMap[gridKey] = gridKey
		}
	}
	return true
}

func convertByteToStr(b byte) string {
	return string(b)
}

func convertRowKey(rowNo int, val string) (rowKey string) {
	rowKey = strconv.Itoa(rowNo) + "r" + val
	return
}

func convertColumnKey(columnNo int, val string) (columnKey string) {
	columnKey = strconv.Itoa(columnNo) + "c" + val
	return
}

func convertGridKey(rowNo int, columnNo int, val string) (gridKey string) {
	rowNo = rowNo / 3
	columnNo = columnNo / 3
	gridKey = strconv.Itoa(rowNo) + strconv.Itoa(columnNo) + "g" + val
	return
}