package no0033

//搜索旋转排序数组
//假设按照升序排序的数组在预先未知的某个点上进行了旋转。
//
//( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
//
//搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
//
//你可以假设数组中不存在重复的元素。
//
//你的算法时间复杂度必须是 O(log n) 级别。
//
//示例 1:
//
//输入: nums = [4,5,6,7,0,1,2], target = 0
//输出: 4
//示例 2:
//
//输入: nums = [4,5,6,7,0,1,2], target = 3
//输出: -1
func search(nums []int, target int) int {
	// 对于排序升序无重复元素, 旋转后的数组,
	// 首索引值必定大于尾索引值（数组长度大于1）
	// 异常校验
	numsLen := len(nums)
	if numsLen == 0 {
		return -1
	}

	startInx := 0
	endInx := numsLen - 1

	// 循环
	for startInx <= endInx {
		// 中间值索引
		midInx := (startInx + endInx) / 2

		// 因为是旋转, 先判定旋转数组区间
		midVal := nums[midInx]
		startVal := nums[startInx]
		endVal := nums[endInx]

		// 如果中值为目标值, 直接返回结果
		if midVal == target {
			return midInx
		}

		if startInx == endInx {
			break
		}

		// 数组前半部存在旋转, 后半部为升序数组
		if startVal > midVal {
			// 判断目标值是否在升序数组中
			if target > midVal && target <= endVal {
				// 目标值在升序数组中
				startInx = midInx + 1
			} else {
				// 目标值不在升序数组中
				endInx = midInx - 1
			}
			continue
		}

		// 数组前半部为升序数组, 后半部存在旋转
		if midVal > endVal {
			// 判断目标值是否在升序数组中
			if target < midVal && target >= startVal {
				// 目标值在升序数组中
				endInx = midInx - 1
			} else {
				// 目标值不在升序数组中
				startInx = midInx + 1
			}
			continue
		}

		// 循环部分已为升序数组, 二分查找
		if startVal < endVal {
			if midVal < target {
				startInx = midInx + 1
			} else {
				endInx = midInx - 1
			}
			continue
		}
	}

	// 未能找到结果
	return -1
}

