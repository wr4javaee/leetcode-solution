package insertion_sort_array

func insertionSort(nums []int) []int {
	numsLen := len(nums)

	if numsLen <= 1 {
		return nums
	}

	// i从第0个索引开始遍历
	for i := 1; i < numsLen; i++ {
		for j := i; j > 0; j-- {
			if nums[j] < nums[j - 1] {
				// 交换较小值
				tmp := nums[j]
				nums[j] = nums[j - 1]
				nums[j - 1] = tmp
			}
		}
	}

	return nums
}
