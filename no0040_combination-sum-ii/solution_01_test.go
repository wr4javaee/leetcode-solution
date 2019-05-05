package no0040_combination_sum_ii

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
				[]int{1, 1},
				1},
			responseModel {[][]int{
				{1},
			}},
		},
		{
			requestModel {
				[]int{10,1,2,7,6,1,5},
			8},
			responseModel {[][]int{
				{1, 7},
				{1, 2, 5},
				{2, 6},
				{1, 1, 6},
			}},

		},
	}

	// 执行测试方法
	for _, v := range testCaseArr {
		req, _ := v.requestModel, v.responseModel
		combinationSum2(req.candidates, req.target)
	}
}