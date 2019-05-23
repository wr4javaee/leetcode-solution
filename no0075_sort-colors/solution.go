package no0075_sort_colors
//
//给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
//
//此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
//
//注意:
//不能使用代码库中的排序函数来解决这道题。
//
//示例:
//
//输入: [2,0,2,1,1,0]
//输出: [0,0,1,1,2,2]
//进阶：
//
//一个直观的解决方案是使用计数排序的两趟扫描算法。
//首先，迭代计算出0、1 和 2 元素的个数，然后按照0、1、2的排序，重写当前数组。
//你能想出一个仅使用常数空间的一趟扫描算法吗？
func sortColors(nums []int)  {

	// 异常校验
	numsLen := len(nums)
	if numsLen <= 1 {
		return
	}

	// 快速排序
	low := 0
	high := numsLen - 1
	sort(nums, low, high)
}

// 采用三向切分的快速排序方法
func sort(nums []int, low int, high int) {
	// 递归基线判定
	if high <= low {
		return
	}

	// 维护lt索引指向第一个重复的v, 使得nums[low ... lt - 1]小于v
	lt := low
	// 维护gt索引, 使得nums[gt + 1 ... high]大于v
	gt := high
	// 切分点v, 可随机从nums中选定
	v := nums[low]
	// 指针i使得nums[lt ... i - 1]都等于v、
	// nums[i ... gt]的元素：
	// 1、若nums[i] == v, i ++
	// 2、若nums[i] < v, 交换nums[lt]与nums[i], 并lt ++, i ++
	// 3、若nums[i] > v, 交换nums[gt]与nums[i], 并gt --
	i := low + 1

	// 当gt缩减到小于等于i时跳出循环
	for gt >= i {
 		ltVal := nums[lt]
		gtVal := nums[gt]
		iVal := nums[i]

		if iVal == v {
			// 若nums[i] == v, i ++
			i ++
		} else if iVal < v {
			// 若nums[i] < v, 交换nums[lt]与nums[i], 并lt ++, i ++
			nums[i] = ltVal
			nums[lt] = iVal
			i ++
			lt ++
		} else {
			// 若nums[i] > v, 交换nums[gt]与nums[i], 并gt --
			nums[i] = gtVal
			nums[gt] = iVal
			gt --
		}
	}

	// 进行完一次快排, 递归
	sort(nums, low, lt - 1)
	sort(nums, gt + 1, high)
}
