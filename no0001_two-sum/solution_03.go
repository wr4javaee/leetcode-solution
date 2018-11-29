package no0001_two_sum

// 一遍哈希表
func twoSum03(nums []int, target int) []int {
	// 初始化map
	numMap := make(map[int]int)

	// 遍历数组
	for i, v := range nums {
		// 计算期望值
		expectNum := target - v

		// 看期望值是否在map中, 若存在便找到解
		if mapVal, ok := numMap[expectNum]; ok {
			return []int {i, mapVal}
		}

		// 将数组中val作为map的key, 数组索引作为map的val
		numMap[v] = i
	}

	// 没有找到结果, 返回空
	return nil
}