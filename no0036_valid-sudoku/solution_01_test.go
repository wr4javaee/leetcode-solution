package no0036_valid_sudoku

import "testing"

type testCaseModel struct {
	requestModel
	responseModel
}

type requestModel struct {
	board [][]byte
}

type responseModel struct {
	result bool
}

func TestIsValidSudoku(t *testing.T) {
	testCaseArr := []testCaseModel {
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
	}

	for _, v := range testCaseArr {
		req, res := v.requestModel, v.responseModel
		if res.result == isValidSudoku(req.board) {
			t.Log("比对成功")
		} else {
			t.Error("比对失败")
		}
	}

}