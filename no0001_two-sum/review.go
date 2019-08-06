package no0001_two_sum

func twoSum(nums []int, target int) []int {

	// request validate
	numsLen := len(nums)

	// O(n)
	numsSet := make(map[int]int, numsLen)
	for i := 0; i < numsLen; i++ {
		currNum := nums[i]
		expectNum := target - currNum
		expectInx, has := numsSet[expectNum]
		if !has {
			numsSet[currNum] = i
			continue
		}

		return []int{ i, expectInx }
	}

	return nil
}
