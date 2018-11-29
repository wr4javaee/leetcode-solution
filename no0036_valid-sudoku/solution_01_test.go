package no0036_valid_sudoku

import "testing"

// 定义单元测试结构体
type testCaseModel struct {
	requestModel
	responseModel
}

// 定义被测试方法请求参数结构体
type requestModel struct {
	board [][]byte
}

// 定义被测试方法返回参数结构体
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
	for i, v := range testCaseArr {
		req, res := v.requestModel, v.responseModel
		if res.result == isValidSudoku(req.board) {
			t.Log("测试通过, 测试用例为", i)
		} else {
			t.Error("测试未通过, 测试用例为", i)
		}
	}

}