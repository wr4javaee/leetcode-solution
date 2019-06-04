package no0053_maximum_subarray

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

func maxSubArray(nums []int) int {
	// 边界校验
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}

	// DF动态规划
	// 初始化最大和
	maxSum := INT_MIN
	// 备忘录
	// key为数组索引, val为截止到该索引（包含）的最大子序列之和
	processMap := make(map[int]int, 100)
	for i := 0; i < numsLen; i++ {
		// i == 0
		if i == 0 {
			maxSum = nums[i]
			processMap[i] = maxSum
			continue
		}

		// i >= 1
		lastSum, _ := processMap[i - 1]
		numsVal := nums[i]
		currSum := max(lastSum + numsVal, numsVal)
		processMap[i] = currSum

		// 记录最大的结果
		if maxSum < currSum {
			maxSum = currSum
		}
	}

	return maxSum
}

// 返回a、b的较大者
func max(a int, b int) int {
	if a >= b {
		return a
	}

	return b
}


