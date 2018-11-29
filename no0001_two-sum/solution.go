package no0001_two_sum

// 方法一：暴力法
func twoSum01(nums []int, target int) []int {
	for i, v := range nums {
		for j:= i + 1; j < len(nums); j ++ {
			if v + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}