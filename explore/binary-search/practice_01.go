package binary_search

//
// 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，
// 写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
//
// 示例1：
// 输入: nums = [-1,0,3,5,9,12], target = 9
// 输出: 4
// 解释: 9 出现在 nums 中并且下标为 4
//
// 示例2：
// 输入: nums = [-1,0,3,5,9,12], target = 2
// 输出: -1
// 解释: 2 不存在 nums 中因此返回 -1
//
// 提示：
// 你可以假设 nums 中的所有元素是不重复的。
// n 将在 [1, 10000]之间。
// nums 的每个元素都将在 [-9999, 9999]之间。
func search(nums []int, target int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return -1
	}
	lowIndex := 0
	highIndex := numsLen - 1
	for lowIndex <= highIndex {
		midIndex := (lowIndex + highIndex) / 2
		midValue := nums[midIndex]
		if midValue == target {
			return midIndex
		} else if midValue < target {
			lowIndex = midIndex + 1
		} else {
			// midValue > target
			highIndex = midIndex - 1
		}
	}

	return -1
}
