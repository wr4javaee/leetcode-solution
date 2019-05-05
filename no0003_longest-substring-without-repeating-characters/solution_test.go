package no0003_longest_substring_without_repeating_characters

import "testing"

// 定义单元测试结构体
type testCaseModel struct {
	requestModel
	responseModel
}

// 定义被测试方法请求参数结构体
type requestModel struct {
	s string
}

// 定义被测试方法返回参数结构体
type responseModel struct {
	result int
}

// 测试方法
func TestIsValidSudoku(t *testing.T) {
	testCaseArr := []testCaseModel {
		{
			requestModel {"au"},
			responseModel {2},

		},
		{
			requestModel {"tmmzuxt"},
			responseModel {5},

		},
		{
			requestModel {"abcabcbb"},
			responseModel {3},

		},
		{
			requestModel {"bbbbb"},
			responseModel {1},

		},
		{
			requestModel {"pwwkew"},
			responseModel {3},
		},
		{
			requestModel {"dvdf"},
			responseModel {3},
		},
	}

	// 执行测试方法一
	for i, v := range testCaseArr {
		if v.result == lengthOfLongestSubstring(v.s) {
			t.Log("测试通过, 测试用例为", i)
		} else {
			t.Error("测试未通过, 测试用例为", i)
		}
	}

	// 执行测试方法二
	for i, v := range testCaseArr {
		if v.result == lengthOfLongestSubstring02(v.s) {
			t.Log("测试通过, 测试用例为", i)
		} else {
			t.Error("测试未通过, 测试用例为", i)
		}
	}

}