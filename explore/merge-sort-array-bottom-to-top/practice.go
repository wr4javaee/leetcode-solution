package merge_sort_array_bottom_to_top

// 归并排序（自底向上）
// 数组按照从小到大的方式进行排序
func mergeSort(nums []int) []int {
	// 异常校验
	numsLen := len(nums)
	if numsLen == 0 {
		return nums
	}

	// 创建一个用于归并的临时数组
	tmpNums := make([]int, len(nums))

	// sz子数组大小
	for sz := 1; sz < numsLen; sz = sz * 2 {
		// 11归并 -> 22归并 -> 44归并 -> 88归并
		for low := 0; low < numsLen - sz; low = low + ( sz * 2 ) {
			// 进行归并
			mid := low + sz - 1
			high := low + ( sz * 2 ) - 1
			if high > numsLen - 1 {
				// 防止最后的子数组越界
				high = numsLen - 1
			}
			merge(nums, tmpNums, low, mid, high)
		}
	}

	return nums
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
