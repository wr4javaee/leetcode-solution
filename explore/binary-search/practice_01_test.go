package binary_search

import "testing"

type testCaseModel struct {
	requestModel
	responseModel
}

type requestModel struct {
	nums []int
	target int
}

type responseModel struct {
	result int
}

func TestSearch(test *testing.T) {
	testCaseArr := []testCaseModel{
		{
			requestModel {
				[]int {-1,0,3,5,9,12},
				9,
			},
			responseModel {4},
		},
	}

	for _, testCase := range testCaseArr {
		request := testCase.requestModel
		response := testCase.responseModel
		searchRes := search(request.nums, request.target)
		if response.result != searchRes {
			test.Error("error", response.result, searchRes)
		}
	}

}
