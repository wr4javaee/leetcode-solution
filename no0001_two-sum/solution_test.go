package no0001_two_sum

import (
	"testing"
)

// 定义单元测试数据结构体
type testCaseModel struct {
	requestModel
	responseModel
}

// 定义被测试方法请求参数结构体
type requestModel struct {
	nums []int
	target int
}

// 定义被测试方法返回参数结构体
type responseModel struct {
	result []int
}

func TestTwoSum(t *testing.T) {
	// 构建测试数据
	testCaseArr := []testCaseModel {
		// 正常数据
		{
			requestModel {
				[]int{2, 7, 11, 15},
				9,
			},
			responseModel {
				[]int{0, 1},
			},
		},
		// 正常数据
		{
			requestModel {
				[]int{3, 2, 4},
				6,
			},
			responseModel {
				[]int{1, 2},
			},
		},
		// 正常数据
		{
			requestModel {
				[]int{3, 3},
				6,
			},
			responseModel {
				[]int{0, 1},
			},
		},
		// 找不到目标值数据
		{
			requestModel {
				[]int{2, 7, 11, 15},
				10,
			},
			responseModel {
				nil,
			},
		},
		// 入参为空的数据
		{
			requestModel {
				[]int{},
				9,
			},
			responseModel {
				nil,
			},
		},
	}

	// 执行测试方法
	// 测试方法一
	for i, v := range testCaseArr {
		actualResult := twoSum01(v.requestModel.nums, v.requestModel.target)
		expectResult := v.responseModel.result
		if resEquals(&actualResult, &expectResult) {
			t.Log("方法一测试通过, 测试用例为", i)
		} else {
			t.Error("方法一测试未通过, 测试用例为", i)
		}
	}

	// 测试方法二
	for i, v := range testCaseArr {
		actualResult := twoSum02(v.requestModel.nums, v.requestModel.target)
		expectResult := v.responseModel.result
		if resEquals(&actualResult, &expectResult) {
			t.Log("方法二测试通过, 测试用例为", i)
		} else {
			t.Error("方法二测试未通过, 测试用例、实际值、期望值分别为", i, actualResult, expectResult)
		}
	}

	// 测试方法三
	for i, v := range testCaseArr {
		actualResult := twoSum03(v.requestModel.nums, v.requestModel.target)
		expectResult := v.responseModel.result
		if resEquals(&actualResult, &expectResult) {
			t.Log("方法三测试通过, 测试用例为", i)
		} else {
			t.Error("方法三测试未通过, 测试用例、实际值、期望值分别为", i, actualResult, expectResult)
		}
	}
}

func resEquals(res1 *[]int, res2 *[]int) bool {
	if *res1 == nil || *res2 == nil {
		return true
	}

	if len(*res1) != len(*res2) {
		return false
	}

	if (*res1)[0] == (*res2)[0] && (*res1)[1] == (*res2)[1] {
		return true
	}

	if (*res1)[0] == (*res2)[1] && (*res1)[1] == (*res2)[0] {
		return true
	}

	return false
}