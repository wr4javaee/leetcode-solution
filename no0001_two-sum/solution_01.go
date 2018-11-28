package no0001_two_sum

// 方法一：暴力法
func twoSum(nums []int, target int) []int {

	// 参数校验
	if len(nums) == 0 {
		return nil
	}

	// 暴力法
	for i, v := range nums {
		for j:= i + 1; j < len(nums); j ++ {
			if v + nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}






