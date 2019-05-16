package selection_sort

// 选择排序
// 从小到大
// O(n^2)
func selectionSort(nums []int) []int {
	numsLen := len(nums)

	if numsLen <= 1 {
		return nums
	}

	// i从第0个索引开始遍历
	for i := 0; i < numsLen - 1; i ++  {
		minValInx := i
		// j从第i + 1个索引开始遍历
		for j := i + 1; j < numsLen; j ++ {
			// 找出较小的数据
			if nums[j] < nums[minValInx] {
				minValInx = j
			}
		}

		// 交换数据
		tmp := nums[i]
		nums[i] = nums[minValInx]
		nums[minValInx] = tmp
	}

	return nums
}

