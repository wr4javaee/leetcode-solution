package no0003_longest_substring_without_repeating_characters

// 滑块法
func lengthOfLongestSubstring02(s string) int {
	strLen := len(s)
	// 返回值
	ans := 0
	// 滑块区间，小曲笺你好:)
	i := 0
	j := i
	// 用于判断字符串是否重复
	valiMap := make(map[uint8]uint8)

	// 滑块从左向右滑动
	for i < strLen && j < strLen {
		jStr := s[j]
		iStr := s[i]
		_, isRepeat := valiMap[jStr]
		if isRepeat && iStr != jStr {
			// 发现有重复字符, 且i的索引字符不等于j的索引字符, i向右滑动, 并删除map中保存的i字符
			i ++
			delete(valiMap, iStr)
		} else if isRepeat && iStr == jStr {
			// 发现有重复字符, 且i的索引字符等于j的索引字符, i, j均向右滑动
			i ++
			j ++
		} else {
			// 非重复字符, j向右边滑动, 并将结果缓存至map中
			valiMap[jStr] = jStr
			j ++
		}

		// 记录滑块长度
		if j == strLen {
			ans = Max(ans, strLen - i)
			continue
		}
		ans = Max(ans, j - i)
	}
	return ans
}