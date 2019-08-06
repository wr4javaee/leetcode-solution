package no0015_3sum

import "sort"

func threeSum(nums []int) [][]int {
	// validate
	numsLen := len(nums)

	// sorted
	sort.Ints(nums)

	// O(n^2)
	res := make([][]int, 0, numsLen)
	for i := 0; i < numsLen - 2; i++ {

		c := nums[i]
		l := i + 1
		r := numsLen - 1

		if c > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}

		for l < r {
			a := nums[l]
			b := nums[r]
			s := c + a + b
			if s == 0 {
				res = append(res, []int{ a, b, c })
				// 去重
				for l < r && nums[r] == nums[r - 1] {
					r --
				}
				for l < r && nums[l] == nums[l + 1] {
					l ++
				}

				l ++
			} else if s > 0 {
				for l < r && nums[r] == nums[r - 1] {
					r --
				}
				r --
			} else {
				for l < r && nums[l] == nums[l + 1] {
					l ++
				}
				l ++
			}

			continue
		}
	}

	return res
}
