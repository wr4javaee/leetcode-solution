package merge_sort_array_top_to_bottom

// 归并排序（自顶向下）
// 数组按照从小到大的方式进行排序
func mergeSort(nums []int) []int {
	// 异常校验
	numsLen := len(nums)
	if numsLen == 0 {
		return nums
	}

	// 创建一个用于归并的临时数组
	tmpNums := make([]int, len(nums))

	// 进行归并排序
	low := 0
	high := numsLen - 1
	sort(nums, tmpNums, low, high)

	return nums
}

// 注意在golang中nums为数组切片，为引用的值拷贝（不是引用指向的值拷贝）
func sort(nums []int, tmpNums []int, low int, high int) {
	// 基准条件
	if high <= low {
		return
	}

	mid := ( low + high ) / 2

	// 将左半边进行归并
	sort(nums, tmpNums, low, mid)

	// 将右半边进行归并
	sort(nums, tmpNums, mid + 1, high)

	// 数据归并
	merge(nums, tmpNums, low, mid, high)
}

// 将nums[low] ~ nums[mid]与nums[mid + 1] ~ nums[high]进行原地归并
// 注意在golang中nums为数组切片，为引用的值拷贝（不是引用指向的值拷贝）
func merge(nums []int, tmpNums []int, low int, mid int, high int) {
	// 初始化临时数组, 用于归并
	for k := low; k <= high; k++ {
		tmpNums[k] = nums[k]
	}

	// 依次比对两个数组的值, 将较小的数据存入原始数组中
	i := low
	j := mid + 1
	for k := low; k <= high; k++ {
		// 越界情况
		if i > mid {
			// i越界, 只归并j的数组
			nums[k] = tmpNums[j]
			j ++
		} else if j > high {
			// j越界, 只归并i的数组
			nums[k] = tmpNums[i]
			i ++
		} else if tmpNums[i] < tmpNums[j] {
			// i较小, 归并i
			nums[k] = tmpNums[i]
			i ++
		} else {
			// j较小, 归并j
			nums[k] = tmpNums[j]
			j ++
		}
	}
}
