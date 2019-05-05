package no0003_longest_substring_without_repeating_characters

// 暴力法
func lengthOfLongestSubstring(s string) int {
	strLen := len(s)
	ans := 0
	for i := 0; i < strLen; i++ {
		for j := i + 1; j <= strLen; j++ {
			subStr := s[i : j]
			isAllUnique, subStrLen := allUnique(subStr)
			if isAllUnique {
				// 取出
				ans = Max(ans, subStrLen)
			}
			continue
		}
	}
	return ans
}

// 所有字符串唯一返回true, 否则返回false
func allUnique(str string) (isAllUnique bool, strLen int) {
	strLen = len(str)
	valiMaps := make(map[int32]int32, strLen)
	for _, v := range str {
		if _, ok := valiMaps[v]; ok {
			isAllUnique = false
			return
		}
		valiMaps[v] = v
	}
	isAllUnique = true
	return
}