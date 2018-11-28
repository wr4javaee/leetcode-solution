# LeetCode.0036 valid-sudoku/有效的数独

## 题目
---
判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。
数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。

![](media/15432126143456/15432126961639.png)

上图是一个部分填充的有效的数独。
数独部分空格内已填入了数字，空白格用 '.' 表示。

示例 1:

```
输入:
[
  ["5","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
输出: true
```
示例 2:

```
输入:
[
  ["8","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
输出: false
解释: 除了第一行的第一个数字从 5 改为 8 以外，空格内其他数字均与 示例1 相同。
     但由于位于左上角的 3x3 宫内有两个 8 存在, 因此这个数独是无效的
```

说明:

* 一个有效的数独（部分已被填充）不一定是可解的。
* 只需要根据以上规则，验证已经填入的数字是否有效即可。
* 给定数独序列只包含数字 1-9 和字符 '.' 。
* 给定数独永远是 9x9 形式的。

## 解题方法

### 方法一
#### 算法时间复杂度
本方法时间复杂度 O(n^2)
#### 思路
* 通过两层循环，按顺序一次遍历每一个数，并按照以下约定进行分类：

```
  ["5","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
```


* 约定第一行序号为0，以第一行数字为例，转义成：
0r5	0r3	0r.	0r.	0r7	0r.	0r.	0r.	0r.	
样例数据转义后变为：

```
0r5	0r3	0r.	0r.	0r7	0r.	0r.	0r.	0r.	
1r6	1r.	1r.	1r1	1r9	1r5	1r.	1r.	1r.	
2r.	2r9	2r8	2r.	2r.	2r.	2r.	2r6	2r.	
3r8	3r.	3r.	3r.	3r6	3r.	3r.	3r.	3r3	
4r4	4r.	4r.	4r8	4r.	4r3	4r.	4r.	4r1	
5r7	5r.	5r.	5r.	5r2	5r.	5r.	5r.	5r6	
6r.	6r6	6r.	6r.	6r.	6r.	6r2	6r8	6r.	
7r.	7r.	7r.	7r4	7r1	7r9	7r.	7r.	7r5	
8r.	8r.	8r.	8r.	8r8	8r.	8r.	8r7	8r9	
```


* 约定第一列序号为0，以第一行数字为例，转义成：
0c5	1c3	2c.	3c.	4c7	5c.	6c.	7c.	8c.	
样例数据转义后变为：

```
0c5	1c3	2c.	3c.	4c7	5c.	6c.	7c.	8c.	
0c6	1c.	2c.	3c1	4c9	5c5	6c.	7c.	8c.	
0c.	1c9	2c8	3c.	4c.	5c.	6c.	7c6	8c.	
0c8	1c.	2c.	3c.	4c6	5c.	6c.	7c.	8c3	
0c4	1c.	2c.	3c8	4c.	5c3	6c.	7c.	8c1	
0c7	1c.	2c.	3c.	4c2	5c.	6c.	7c.	8c6	
0c.	1c6	2c.	3c.	4c.	5c.	6c2	7c8	8c.	
0c.	1c.	2c.	3c4	4c1	5c9	6c.	7c.	8c5	
0c.	1c.	2c.	3c.	4c8	5c.	6c.	7c7	8c9	
```


* 约定第一个九宫格序号为0, 对于每一个九宫格来说，可以判断行号或列号分别对3求余，即可得出第一个九宫格序号，样例数据转义后变为：

```
005	003	00.	01.	017	01.	02.	02.	02.	
006	00.	00.	011	019	015	02.	02.	02.	
00.	009	008	01.	01.	01.	02.	026	02.	
108	10.	10.	11.	116	11.	12.	12.	123	
104	10.	10.	118	11.	113	12.	12.	121	
107	10.	10.	11.	112	11.	12.	12.	126	
20.	206	20.	21.	21.	21.	222	228	22.	
20.	20.	20.	214	211	219	22.	22.	225	
20.	20.	20.	21.	218	21.	22.	227	229	
```

* 最后，将转义后的数字通过集合判断是否重复出现

#### 实现（Golang）
```
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
```
#### 测试（Golang）
```
package no0036_valid_sudoku

import "testing"

// 定义单元测试数据类型
type testCaseModel struct {
	requestModel
	responseModel
}

// 定义被测试方法请求参数数据类型
type requestModel struct {
	board [][]byte
}

// 定义被测试方法返回参数数据类型
type responseModel struct {
	result bool
}

// 测试方法
func TestIsValidSudoku(t *testing.T) {
	testCaseArr := []testCaseModel {
		// 合法的九宫格
		{
			requestModel {
				[][]byte {
					[]byte("53..7...."),
					[]byte("6..195..."),
					[]byte(".98....6."),
					[]byte("8...6...3"),
					[]byte("4..8.3..1"),
					[]byte("7...2...6"),
					[]byte(".6....28."),
					[]byte("...419..5"),
					[]byte("....8..79"),
				}},
			responseModel {true},

		},
		// 行存在重复数据
		{
			requestModel {
				[][]byte {
					[]byte("53..5...."),
					[]byte("6..195..."),
					[]byte(".98....6."),
					[]byte("8...6...3"),
					[]byte("4..8.3..1"),
					[]byte("7...2...6"),
					[]byte(".6....28."),
					[]byte("...419..5"),
					[]byte("....8..79"),
				}},
			responseModel {false},

		},
		// 列存在重复数据
		{
			requestModel {
				[][]byte {
					[]byte("53..7...."),
					[]byte("6..195..."),
					[]byte(".98....6."),
					[]byte("8...6...3"),
					[]byte("4..8.3..1"),
					[]byte("5...2...6"),
					[]byte(".6....28."),
					[]byte("...419..5"),
					[]byte("....8..79"),
				}},
			responseModel {false},

		},
		// 小九宫格存在重复数据
		{
			requestModel {
				[][]byte {
					[]byte("53..7...."),
					[]byte("6..195..."),
					[]byte(".98....6."),
					[]byte("8...6...3"),
					[]byte("4..8.3..1"),
					[]byte("7...2...6"),
					[]byte("96....28."),
					[]byte("..9417..5"),
					[]byte("....8..79"),
				}},
			responseModel {false},

		},
	}

	// 执行测试方法
	for _, v := range testCaseArr {
		req, res := v.requestModel, v.responseModel
		if res.result == isValidSudoku(req.board) {
			t.Log("比对成功")
		} else {
			t.Error("比对失败")
		}
	}

}
```






















































