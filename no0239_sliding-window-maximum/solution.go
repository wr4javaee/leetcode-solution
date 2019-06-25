package no0239_sliding_window_maximum

import "container/list"

func maxSlidingWindow(nums []int, k int) []int {
	// 异常校验
	if nums == nil {
		return []int{}
	}

	// 使用双向队列
	deque := list.New()
	numsLen := len(nums)

	// 遍历数组
	res := make([]int, 0, numsLen - k + 1)
	for i := 0; i < numsLen; i++ {
		// 维持双向队列按照从大到小的顺序排列
		// 若队首大于当前元素，则出队
		for deque.Len() > 0 {
			backInx := deque.Back().Value.(int)
			if nums[backInx] < nums[i] {
				// 队尾小于当前元素
				deque.Remove(deque.Back())
				continue
			}
			// 队首大于等于当前元素
			break
		}

		deque.PushBack(i)
		frontInx := deque.Front().Value.(int)
		// 判断队首是否越界
		if frontInx < i - k + 1 {
			// 越界, 移出
			deque.Remove(deque.Front())
		}

		// 判断是否需要输出结果
		if i + 1 >= k {
			res = append(res, nums[deque.Front().Value.(int)])
		}

	}

	return res
}
