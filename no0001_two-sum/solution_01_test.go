package no0001_two_sum

import (
	"testing"
)

func TestCase01(t *testing.T) {
	expectResult := []int{0, 1}
	nums := []int{2, 7, 11, 15}
	target := 9
	actualResult := twoSum(nums, target)
	if expectResult[0] == actualResult[0] && expectResult[1] == actualResult[1] {
		t.Log("测试通过")
	} else {
		t.Error("测试未通过")
	}
}