package no0018_4sum

import "sort"

func fourSum(nums []int, target int) [][]int {

	// validate
	numsLen := len(nums)

	// sorted
	sort.Ints(nums)

	// O(n^3)
	res := make([][]int, 0, numsLen)
	for i := 0; i < numsLen - 3; i ++ {
		a := nums[i]
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}

		for j := i + 1; j < numsLen - 2; j ++ {
			b := nums[j]
			if j > i + 1 && nums[j] == nums[j - 1] {
				continue
			}

			left := j + 1
			right := numsLen - 1

			for left < right {
				c := nums[left]
				d := nums[right]
				s := a + b + c + d
				if s == target {
					for left < right && nums[left] == nums[left + 1] {
						left ++
					}
					for left < right && nums[right] == nums[right - 1] {
						right --
					}

					left ++
					res = append(res, []int{ a, b, c, d })
				} else if s > target {
					for left < right && nums[right] == nums[right - 1] {
						right --
					}
					right --
				} else {
					for left < right && nums[left] == nums[left + 1] {
						left ++
					}
					left ++
				}
			}
		}
	}

	return res
}
