package no0006_zigzag_conversion

import "testing"

// 定义单元测试结构体
type testCaseModel struct {
	requestModel
	responseModel
}

// 定义被测试方法请求参数结构体
type requestModel struct {
	s string
	numRows int
}

// 定义被测试方法返回参数结构体
type responseModel struct {
	result string
}

// 测试方法
func TestConvert(t *testing.T) {
	testCaseArr := []testCaseModel {
		{
			requestModel {"LEETCODEISHIRING", 3},
			responseModel {"LCIRETOESIIGEDHN"},

		},
		{
			requestModel {"AB", 1},
			responseModel {"AB"},

		},
	}

	// 执行测试方法一
	for i, v := range testCaseArr {
		if v.result == convert(v.s, v.numRows) {
			t.Log("测试通过, 测试用例为", i)
		} else {
			t.Error("测试未通过, 测试用例为", i)
		}
	}
}