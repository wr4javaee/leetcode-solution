package no0039_combination_sum

import "testing"

// 定义单元测试结构体
type testCaseModel struct {
	requestModel
	responseModel
}

// 定义被测试方法请求参数结构体
type requestModel struct {
	candidates []int
	target int
}

// 定义被测试方法返回参数结构体
type responseModel struct {
	result [][]int
}

// 测试方法
func TestIsValidSudoku(t *testing.T) {
	testCaseArr := []testCaseModel {
		{
			requestModel {
				[]int{2, 3, 6, 7},
			7},
			responseModel {[][]int{
				{7},
				{2, 2, 3},
			}},

		},
	}

	// 执行测试方法
	for _, v := range testCaseArr {
		req, _ := v.requestModel, v.responseModel
		combinationSum(req.candidates, req.target)
	}
}