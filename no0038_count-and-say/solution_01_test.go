package no0038_count_and_say

import "testing"

// 定义单元测试结构体
type testCaseModel struct {
	requestModel
	responseModel
}

// 定义被测试方法请求参数结构体
type requestModel struct {
	n int
}

// 定义被测试方法返回参数结构体
type responseModel struct {
	result string
}

// 测试方法
func TestCountAndSay(t *testing.T) {
	testCaseArr := []testCaseModel {
		{
			requestModel {
				4,
			},
			responseModel {"1211"},
		},
		{
			requestModel {
				1,
				},
			responseModel {"1"},
		},
	}

	// 执行测试方法
	for i, v := range testCaseArr {
		req, res := v.requestModel, v.responseModel
		if res.result == countAndSay(req.n) {
			t.Log("测试通过, 测试用例为", i)
		} else {
			t.Error("测试未通过, 测试用例为", i)
		}
	}

}