package no0034_find_first_and_last_position_of_element_in_sorted_array
//
//给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
//
//你的算法时间复杂度必须是 O(log n) 级别。
//
//如果数组中不存在目标值，返回 [-1, -1]。
//
//示例 1:
//
//输入: nums = [5,7,7,8,8,10], target = 8
//输出: [3,4]
//示例 2:
//
//输入: nums = [5,7,7,8,8,10], target = 6
//输出: [-1,-1]
func searchRange(nums []int, target int) []int {
	// 解题思路：
	// 采用三个二分法函数，分别找左边界和右边界
	// 1、第一个二分法函数，用于定位Target, 若数组中无法定位target，返回[-1, -1]
	// 2、若定位到target, 则进行另外两个二分法函数, 分别定位左边界与右边界

	// 特殊参数校验
	numsLen := len(nums)
	if numsLen == 0 {
		return []int{ -1, -1 }
	}

	lowInx := 0
	highInx := numsLen - 1

	// 第一个二分法, 先尝试定位到某一个target
	// 简单二分法
	for lowInx <= highInx {
		midInx := ( lowInx + highInx ) / 2
		midVal := nums[midInx]

		// 定位到一个target
		if midVal == target {
			// 变形二分法, 保证查找空间在每个步骤中至少有 3 个元素
			return []int { findLeftTarget(nums, target, lowInx, highInx), findRightTarget(nums, target, lowInx, highInx) }
		}

		if midVal > target {
			highInx = midInx - 1
		} else {
			lowInx = midInx + 1
		}
	}

	return []int { -1, -1}
}

// 变形二分法, 保证查找空间在每个步骤中至少有 3 个元素
// 若midInx值等于target, 判断左邻节点是否等于mid, 若不相等则左区间为midInx, 若相等则向左进行二分
// 若midInx值不等于target, 则进行正常变形二分, 直到找到左区间为止
func findLeftTarget(nums []int, target int, lowInx int, highInx int) (leftInx int) {
	for lowInx + 1 < highInx {
		midInx := ( lowInx + highInx ) / 2
		midVal := nums[midInx]
		midLeftInx := midInx - 1
		midLeftVal := nums[midLeftInx]

		// midInx值等于target且midLeftInx值不等于target, 定位到了左区间
		if midVal == target && midLeftVal != target {
			return midInx
		} else if midVal == target && midLeftVal == target {
			// 向左进行二分
			highInx = midInx
			continue
		} else if midVal > target {
			// 向左进行二分
			highInx = midInx
			continue
		} else {
			// 向右进行二分
			lowInx = midInx
		}
	}

	// 针对剩余少于3个元素区间进行判断
	lowVal := nums[lowInx]
	highVal := nums[highInx]
	if lowInx == highInx && lowVal == target {
		return lowInx
	} else if lowInx == highInx && lowVal != target {
		// 不可能出现
		return -1
	} else if lowInx != highInx && lowVal == target {
		return lowInx
	} else if lowInx != highInx && highVal == target {
		return highInx
	} else {
		return -1
	}
}

// 变形二分法, 保证查找空间在每个步骤中至少有 3 个元素
// 若midInx值等于target, 判断右邻节点是否等于mid, 若不相等则右区间为midInx, 若相等则向右进行二分
// 若midInx值不等于target, 则进行正常变形二分, 直到找到右区间为止
func findRightTarget(nums []int, target int, lowInx int, highInx int) (rightInx int) {
	for lowInx + 1 < highInx {
		midInx := ( lowInx + highInx ) / 2
		midVal := nums[midInx]
		midRightInx := midInx + 1
		midRightVal := nums[midRightInx]

		// midInx值等于target且midLeftInx值不等于target, 定位到了左区间
		if midVal == target && midRightVal != target {
			return midInx
		} else if midVal == target && midRightVal == target {
			// 向右进行二分
			lowInx = midInx
			continue
		} else if midVal > target {
			// 向左进行二分
			highInx = midInx
			continue
		} else {
			// 向右进行二分
			lowInx = midInx
		}
	}

	// 针对剩余少于3个元素区间进行判断
	lowVal := nums[lowInx]
	highVal := nums[highInx]
	if lowInx == highInx && lowVal == target {
		return lowInx
	} else if lowInx == highInx && lowVal != target {
		// 不可能出现
		return -1
	} else if lowInx != highInx && highVal == target {
		return highInx
	} else if lowInx != highInx && lowVal == target {
		return lowInx
	} else {
		return -1
	}
}
