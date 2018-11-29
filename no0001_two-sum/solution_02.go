package no0001_two_sum

// 两遍哈希表
func twoSum02(nums []int, target int) []int {
	// 初始化map
	numMap := make(map[int]int)

	// 遍历数组, 将数组中val作为map的key, 数组索引作为map的val
	for i, v := range nums {
		numMap[v] = i
	}

	// 再次遍历数组
	for i, v := range nums {
		expectNum := target - v

		// 按照题目要求不可重复使用同一个索引, 因此这里需要保证数组的索引和map中取出的索引不能相等
		if mapVal, ok := numMap[expectNum]; ok && i != mapVal {
			return []int {i, mapVal}
		}
	}

	// 没有找到结果, 返回空
	return nil
}